package nodesvcs

import (
	"bufio"
	"context"
	"io"
	"net"
	"os/exec"
	"strings"
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

	STDOUT = 1
	STDERR = 2

	//XCTPRGNAME -- name of this program
	XCTPRGNAME = "CSXct"
)

//ReturnLineChan -- for transferring cmd stdout and stderr
type ReturnLineChan struct {
	ReturnLine communicate.ReturnTaskLine
	ReturnAdr  net.Addr
}

//ExecutorSVC --
type ExecutorSVC struct {
	publicKey  string
	unixSocket string

	RecvPC net.PacketConn //coresponding to recvPort
	SendPC net.PacketConn //coresponding to sendPort

	ReturnLineChan chan ReturnLineChan

	//WG -- waitgroup for this service
	wg  *sync.WaitGroup
	ctx context.Context
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

//HandleRecvPort -- a func of commNode interface
//handle recerived bytes stream on UDP port
func (s *ExecutorSVC) HandleRecvPort() error {
	var err error
	s.RecvPC, err = net.ListenPacket("udp", communicate.XCTRECVPPORT)
	if err != nil {
		logger.Error("HandleRecvPort --", err)
		return err
	}

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		defer s.RecvPC.Close()

		for {
			buf := make([]byte, communicate.MAXUDPPACKET)
			//n, addr, err := conn.ReadFrom(buf)
			n, adr, err := s.RecvPC.ReadFrom(buf)
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
				go s.DoTask(s.ctx, m, adr)
				//logger.Debug(n, "bytes got. data=", string(p.SrcID[:]),
				//	string(p.Payload), p, err)
			}
		}
	}()

	return nil
}

//DoTask --
func (s *ExecutorSVC) DoTask(ctx context.Context, m *communicate.MSGObj, adr net.Addr) {
	tid := m.TaskID
	defer logger.Debug(tid, "exit DoTask...")

	logger.Debug(tid, "enter DoTask...")
	if m.ObjType == communicate.ObjTypeShellCmd {

		s.DoShellCMD(ctx, m.Obj.(communicate.ShellCMD).Script, tid, adr)

	}
}

func wrapStdoutLineChan(tid string, ltype byte,
	lbytes []byte, adr net.Addr) (*ReturnLineChan, error) {

	var r ReturnLineChan

	ml := communicate.MSGLine{
		TaskID: tid,
		Line:   string(lbytes),
	}

	lineBytes, err := ml.MarshalMSG()
	if err != nil {
		return nil, err
	}

	r = ReturnLineChan{
		ReturnLine: communicate.ReturnTaskLine{
			LineType:  ltype,
			LineBytes: lineBytes,
		},
		ReturnAdr: adr,
	}

	return &r, nil
}

//DoShellCMD --
func (s *ExecutorSVC) DoShellCMD(ctx context.Context, cmdStr string,
	tid string, adr net.Addr) {

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
	go func(tid string, stdoutIn io.ReadCloser, adr net.Addr) {
		defer innerWG.Done()

		logger.Debug(tid, "enter goroutine stdoutIn ...")

		scanner := bufio.NewScanner(stdoutIn)

		for scanner.Scan() {

			rLC, err := wrapStdoutLineChan(tid, STDOUT, scanner.Bytes(), adr)
			if err != nil {
				logger.Error("wrapStdoutLineChan1-", err)
				continue
			}

			s.ReturnLineChan <- *rLC

			select {
			default:
				continue
			case <-ctx.Done():
				break
			}
		}

		rLC, err := wrapStdoutLineChan(tid, STDOUT, []byte(communicate.STDOUTEOF), adr)
		if err != nil {
			logger.Error("wrapStdoutLineChan2-", err)
		} else {
			s.ReturnLineChan <- *rLC
		}

		logger.Debug(tid, "exit goroutine stdoutIn ...")
	}(tid, stdoutIn, adr)

	innerWG.Add(1)
	go func(tid string, stderrIn io.ReadCloser, adr net.Addr) {
		defer innerWG.Done()
		logger.Debug(tid, "enter goroutine stderrIn ...")

		scanner := bufio.NewScanner(stderrIn)

		for scanner.Scan() {
			rLC, err := wrapStdoutLineChan(tid, STDERR, scanner.Bytes(), adr)
			if err != nil {
				logger.Error("wrapStdoutLineChan3-", err)
				continue
			}

			s.ReturnLineChan <- *rLC

			select {
			default:
				continue
			case <-ctx.Done():
				break
			}
		}

		rLC, err := wrapStdoutLineChan(tid, STDERR, []byte(communicate.STDERREOF), adr)
		if err != nil {
			logger.Error("wrapStdoutLineChan2-", err)
		} else {
			s.ReturnLineChan <- *rLC
		}

		logger.Debug(tid, "exit goroutine stderrIn ...")
	}(tid, stderrIn, adr)

	innerWG.Wait()

	err = cmd.Wait()
	if err != nil {
		logger.Error("cmd.Run() failed:", err)
	}

	logger.Debug(tid, "exit DoShellCMD ...")
}

//HandleUnixSocket -- handle recerived bytes stream on local unix domain socket
func (s *ExecutorSVC) HandleUnixSocket() error {
	return nil
}

//HandleSendPort -- handle udp sending port
func (s *ExecutorSVC) HandleSendPort() error {
	var err error
	s.SendPC, err = net.ListenPacket("udp", communicate.XCTSENDPORT)
	if err != nil {
		return err
	}

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		defer s.SendPC.Close()

		i := 0
		for {
			select {
			case lc := <-s.ReturnLineChan:
				i++
				logger.Debug(i, "got LineChan, bytes number:", len(lc.ReturnLine.LineBytes))

				//send back to caller
				str := lc.ReturnAdr.String()
				ipright := strings.LastIndex(str, ":")
				dststr := str[:ipright] + communicate.SCHRECVPPORT
				dst, err := net.ResolveUDPAddr("udp", dststr)
				if err != nil {
					logger.Error("ResolveUDPAddr-", err)
					continue
				}

				returnLine, err := lc.ReturnLine.MarshalTL()
				if err != nil {
					logger.Error("MarshalTL-", err)
					continue
				}
				logger.Debug(i, "after marshal, returnLine bytes number:", len(returnLine))

				_, err = s.SendPC.WriteTo(returnLine, dst)
				if err != nil {
					logger.Error("WriteTo-", err)
					continue
				}

			case <-s.ctx.Done():
				break
			}
		}
	}()

	return nil
}

//Init -- this is a default func which will be invoked automatically at instance creating.
func (s *ExecutorSVC) Init(ctx context.Context, wg *sync.WaitGroup) error {
	s.publicKey = PUBKEY
	s.unixSocket = XCTUNIXSOCKET
	s.wg = wg
	s.ctx = ctx

	err := s.HandleSendPort()
	if err != nil {
		return nil
	}
	err = s.HandleRecvPort()
	if err != nil {
		return nil
	}
	err = s.HandleUnixSocket()
	if err != nil {
		return nil
	}

	return nil
}
