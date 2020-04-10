package nodesvcs

import (
	"bufio"
	"context"
	"net"
	"os/exec"
	"sync"

	"ire.com/clustershell/communicate"

	"ire.com/clustershell/logger"
)

//PUBKEY -- the public key of key pairs for communication between scheduler and executor
//generate a key pair by following command in linux:
//ssh-keygen -t rsa -P "" -C "commNode@clustershell.ire.com" -f /tmp/clustershell
//cat /tmp/clustershell ### copy the output and paste to PUBKEY, and hide the private key /tmp/clustershell
const (
	PUBKEY = `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDc0QK2LqxlU6StKJ7/3qZ+A5M4w2vaaDuub5BO5v+PPMn65CChdztoK9/QSermpTmc+rjSv2UJJSHqztRxr01DUehBuoOgT7LSh0GMm2HyobMpgVG6v/b8k2S//AkuxmltxkFzgUApBGsWWcgbn88HQLc3QF4P/60It9EMSwy7GEppdsvhGvfS5UyKqDD2Xuz8zHKP5WU1/lt4f1m44K1ZgZTzQv8r53des37m/TmVrYX1vBEflHSC7xTWF37WUXpET+jBru69zUoiM2jzbw89BX2i+S9fbDZTb5OAHVbnevQTNBmkt1D0uwcgur0Xw4hS/Vr/fpt2i9F6Askcougr commNode@clustershell.ire.com`

	XCTUNIXSOCKET = `/var/run/ire_clshexecutor.sock`

	XCTUDPPORT = ":33225"
)

//LineChan -- for transferring cmd stdout and stderr
type LineChan struct {
	TaskID    communicate.TaskIDType
	LineType  byte //1 -- stdout, 2 -- stderr
	LineBytes []byte
}

//ExecutorSVC --
type ExecutorSVC struct {
	publicKey  string
	unixSocket string
	udpPort    string

	LineChan chan LineChan

	//WG -- waitgroup for this service
	wg *sync.WaitGroup
}

//Encrypt -- a func of commNode interface
func (s *ExecutorSVC) Encrypt(unEncrypted []byte, comKey string) (
	encrypted []byte, err error) {
	return nil, nil
}

//Decrpyt -- a func of commNode interface
func (s *ExecutorSVC) Decrpyt(encrypted []byte, comKey string) (
	unEncrypted []byte, err error) {
	return nil, nil
}

//HandleSendMsg -- a func of commNode interface
//handle returned message after invoking SendMsg
func (s *ExecutorSVC) HandleSendMsg() error {
	return nil
}

//HandleListenOnUnixSocket -- a func of commNode interface
//handle recerived bytes stream on local unix domain socket
func (s *ExecutorSVC) HandleListenOnUnixSocket() error {
	return nil
}

//HandleListenOnUDP -- a func of commNode interface
//handle recerived bytes stream on UDP port
func (s *ExecutorSVC) HandleListenOnUDP(ctx context.Context) error {
	recvPort, err := net.ListenPacket("udp", s.udpPort)
	if err != nil {
		logger.Error(err)
	}

	//todo: send back cmd output to caller
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		i := 0
		for {
			select {
			case lc := <-s.LineChan:
				i++
				logger.Debug(i, "got LineChan:", lc)
				//todo send back to caller
			case <-ctx.Done():
				break
			}
		}
	}()

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		defer recvPort.Close()

		for {
			buf := make([]byte, 1024*40)
			//n, addr, err := conn.ReadFrom(buf)
			n, adr, err := recvPort.ReadFrom(buf)
			if err != nil {
				logger.Error(err)
				continue
			}

			if n > 30 { //ignor nonsense short packets
				p, err := communicate.UnMarshalPacket(buf[:n])
				if err != nil {
					continue //ignore all unacquainted msgs
				}

				m := new(communicate.MSGObj)
				err = m.UnMarshalMSG(p.Payload)
				if err != nil {
					logger.Error("UnMarshalMSG --", err)
					continue //ignore all unacquainted msgs
				}

				logger.Info("got a task:", *m, "from", adr)
				go s.DoTask(ctx, m)
				//logger.Debug(n, "bytes got. data=", string(p.SrcID[:]),
				//	string(p.Payload), p, err)
			}
		}
	}()

	return nil
}

//DoTask --
func (s *ExecutorSVC) DoTask(ctx context.Context, m *communicate.MSGObj) {
	Tid := m.TaskID
	defer logger.Debug(Tid, "exit DoTask...")

	logger.Debug(Tid, "enter DoTask...")
	if m.ObjType == communicate.ObjTypeShellCmd {

		s.DoShellCMD(ctx, m.Obj.(communicate.ShellCMD).Script, Tid)

	}
}

//DoShellCMD --
func (s *ExecutorSVC) DoShellCMD(ctx context.Context, cmdStr string,
	tid communicate.TaskIDType) {

	logger.Debug(tid, "enter DoShellCMD...")

	cmd := exec.CommandContext(ctx, "bash", cmdStr)

	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	err := cmd.Start()
	if err != nil {
		logger.Error("cmd.Start() failed:", err)
	}

	innerWG := new(sync.WaitGroup)
	innerWG.Add(1)
	go func() {
		defer innerWG.Done()

		logger.Debug(tid, "enter goroutine stdoutIn ...")

		scanner := bufio.NewScanner(stdoutIn)

		for scanner.Scan() {
			s.LineChan <- LineChan{
				TaskID:    tid,
				LineType:  1,
				LineBytes: scanner.Bytes(),
			}

			select {
			default:
				continue
			case <-ctx.Done():
				break
			}
		}
		logger.Debug(tid, "exit goroutine stdoutIn ...")
	}()

	innerWG.Add(1)
	go func() {
		defer innerWG.Done()
		logger.Debug(tid, "enter goroutine stderrIn ...")

		scanner := bufio.NewScanner(stderrIn)

		for scanner.Scan() {
			s.LineChan <- LineChan{
				TaskID:    tid,
				LineType:  1,
				LineBytes: scanner.Bytes(),
			}

			select {
			default:
				continue
			case <-ctx.Done():
				break
			}
		}
		logger.Debug(tid, "exit goroutine stderrIn ...")
	}()

	innerWG.Wait()

	err = cmd.Wait()
	if err != nil {
		logger.Error("cmd.Run() failed:", err)
	}

	logger.Debug(tid, "exit DoShellCMD ...")
}

//ListenOnUDP -- a func of commNode interface
//lauch a continuous listening on UDP port
func (s *ExecutorSVC) ListenOnUDP() error {
	return nil
}

//ListenOnUnixSocket -- a func of commNode interface
//lauch a continuous listening on UNIX domain socket
func (s *ExecutorSVC) ListenOnUnixSocket() error {
	return nil
}

//SendFile -- a func of commNode interface
func (s *ExecutorSVC) SendFile(fileName string, destIPPort string, destPath string, fileMode []byte) error {
	return nil
}

//SendMsg -- a func of commNode interface
func (s *ExecutorSVC) SendMsg(message []byte, destIPPort string) error {
	return nil
}

//Init -- this is a default func which will be invoked automatically at instance creating.
func (s *ExecutorSVC) Init(wg *sync.WaitGroup) error {
	s.publicKey = PUBKEY
	s.unixSocket = XCTUNIXSOCKET
	s.udpPort = XCTUDPPORT
	s.wg = wg

	return nil
}
