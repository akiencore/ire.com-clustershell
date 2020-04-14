package nodesvcs

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"

	"ire.com/clustershell/communicate"
	"ire.com/clustershell/logger"
)

//PVTKEY -- THIS IS A VERY CONFIDENTIAL INFOMATION.
//ONLY ASSIGN THIS VARIABLE BEFORE YOU BUILD BINARY EXECUTIVE,
//AND REMOVE IT TO EMPTY AFTER YOUR BUIDING, AND MAKE SURE
//ALL OTHERS CANNOT GET IT IN SOURCE CODE.
const (
	PVTKEY = `-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAABFwAAAAdzc2gtcn
NhAAAAAwEAAQAAAQEA3NECti6sZVOkrSie/96mfgOTOMNr2mg7rm+QTub/jzzJ+uQgoXc7
aCvf0Enq5qU5nPq40r9lCSUh6s7Uca9NQ1HoQbqDoE+y0odBjJth8qGzKYFRur/2/JNkv/
wJLsZpbcZBc4FAKQRrFlnIG5/PB0C3N0BeD/+tCLfRDEsMuxhKaXbL4Rr30uVMiqgw9l7s
/Mxyj+VlNf5beH9ZuOCtWYGU80L/K+d3XrN+5v05la2F9bwRH5R0gu8U1hd+1lF6RE/owa
7uvc1KIjNo828PPQV9ovkvX2w2U2+TgB1W53r0EzQZpLdQ9LsHILq9F8OIUv1a/36bdovR
egLJHKLoKwAAA9hF42D6ReNg+gAAAAdzc2gtcnNhAAABAQDc0QK2LqxlU6StKJ7/3qZ+A5
M4w2vaaDuub5BO5v+PPMn65CChdztoK9/QSermpTmc+rjSv2UJJSHqztRxr01DUehBuoOg
T7LSh0GMm2HyobMpgVG6v/b8k2S//AkuxmltxkFzgUApBGsWWcgbn88HQLc3QF4P/60It9
EMSwy7GEppdsvhGvfS5UyKqDD2Xuz8zHKP5WU1/lt4f1m44K1ZgZTzQv8r53des37m/TmV
rYX1vBEflHSC7xTWF37WUXpET+jBru69zUoiM2jzbw89BX2i+S9fbDZTb5OAHVbnevQTNB
mkt1D0uwcgur0Xw4hS/Vr/fpt2i9F6AskcougrAAAAAwEAAQAAAQAHbRAGSHOLxNBY8nDG
chpvSnd9BTVvVuoK572WqbYWfxjb0yg1xei9jBwuyQ2ZTB0v5k6o577iG9mUJ/iFCjwy82
I4M6mRjpyY7+uIawFUQ5Pe5XZ2LVaFD//nZPZ6GDobcvtogwRBXOCtE7/jDfDMcaS3uvZu
205aaLQjEbMP56U50OSDGyOP2ZCbG8mS97I0kgmPP6gMsw/g8nd8zig+JRcJZ1tJPUJFii
MP2Vq2VbRY87rJ0F13W29Kk/tN+Sdsgg3rnRiPNm5LnRs1st5bvbXMGppdsf6KyrWuFCV9
fV4s9WsvYfEa9bkOWioyU2z3oA8FtzoSBuQ3WZVQUrZhAAAAgQDB3xzW3JcSE89n6PI/XI
HnW5GrT3vkxsudvHD74ARhkdcvZvZ7Bhhfpn79ErIBN6kpdjJGaPdw1Q5T61zMbxQ6wGqW
ArFhdTtf/tXwPcukkVPu6ybXChbMWOi5dPkBrF++K1GejmjK+J/9SYpkglj3aQixpE8IFl
Dungtpp+orxQAAAIEA88ZSiL3Ly4gCnw3WyHQ7hr6JXe9FHjQqXWgv5PP3laGzOr8uUC8Y
bVyDOM+yVN7wnwAPNeONqdj4+zrvK95gTPJ7uElSmLdF/bwYzECnj3ZoYGZFIXvk4tb5ZN
Y8YqaE+XZlfg9/unkMVVwWVJjYDvL5rrnW6b7OzRFNqWrwy6cAAACBAOfj8OJ/zqF93JW2
2ToxJg0B1aZdlL5e2m4Imqwkbdls2V2hc2JB+OkpKxfdKwCgY2KzsRFqHwj0MboglcDanl
Bv21dNuxipQlxsNqAW5MNORZbBEhVJs2+8brGXi/sIaT218W/RBVpxmNscFFC+Ygge2Yzv
59XPyU0XqAmi0z/dAAAAHWNvbW1Ob2RlQGNsdXN0ZXJzaGVsbC5pcmUuY29tAQIDBAU=
-----END OPENSSH PRIVATE KEY-----`

	SCHUNIXSOCKET = `/var/run/ire_clshsheduler.sock`
	//SCHDLPRGNAME -- name of this program
	SCHDLPRGNAME = "CSSch"
	SCHDLPATH    = "/opt/" + SCHDLPRGNAME + "/"
)

//SendMsgChan --
type SendMsgChan struct {
	TaskID  string
	Message []byte
	DestIP  string
}

//SchedulerSVC --
type SchedulerSVC struct {
	privateKey  string
	unixSocket  string
	RecvPC      net.PacketConn //coresponding to recvPort
	SendPC      net.PacketConn //coresponding to sendPort
	SendMsgChan chan SendMsgChan

	wg  *sync.WaitGroup
	ctx context.Context

	StdoutFileMap map[string]*os.File //taskid is key
	StderrFileMap map[string]*os.File
}

//Encrypt -- a func of commNode interface
func (s *SchedulerSVC) Encrypt(unEncrypted []byte, comKey string) (
	encrypted []byte, err error) {
	return nil, nil
}

//Decrpyt -- a func of commNode interface
func (s *SchedulerSVC) Decrpyt(encrypted []byte, comKey string) (
	unEncrypted []byte, err error) {
	return nil, nil
}

//HandleUnixSocket -- a func of commNode interface
//lauch a continuous listening on UNIX domain socket
func (s *SchedulerSVC) HandleUnixSocket() error {
	os.Remove(SCHUNIXSOCKET)
	addr, err := net.ResolveUnixAddr("unix", s.unixSocket)
	if err != nil {
		return fmt.Errorf("failed to resolve: %v", err)
	}

	list, err := net.ListenUnix("unix", addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		defer list.Close()

		for {
			conn, err := list.AcceptUnix()
			if err != nil {
				logger.Error("failed to create connection on unix socket")
				continue
			}

			s.wg.Add(1)
			go func(conn *net.UnixConn) {
				defer s.wg.Done()
				defer conn.Close()

				for {
					buf := make([]byte, communicate.MSGMAXLEN+1)
					n, _, err := conn.ReadFromUnix(buf)
					if err != nil {
						s := err.Error()
						s = s[len(s)-3:]
						if s != "EOF" {
							logger.Debug("unix socket ReadFromUnix:", err)
						}
						break //connection is over
					}
					logger.Debug("received", n, "bytes:", string(buf))

					//if n > SOCKETMSGMAXLEN {
					//todo
					//}

					pTasksOfConn, err := communicate.GetMSGListFromSocketBuf(buf[:n], conn)
					if err != nil {
						logger.Debug("GetMSGListFromSocketBuf:", err)
						break
					}
					logger.Debug("GetMSGListFromSocketBuf got map len", len(*pTasksOfConn))

					for tid, task := range *pTasksOfConn {
						logger.Debug("tid=", tid, "task=", task)
						err = s.LaunchTask(&task)
						if err != nil {
							logger.Error("LaunchTask --", err)
						}

						logger.Info("launched a task of", task.Msgobj.ObjType, "to",
							task.Msgobj.DestIP, ", taskid =", task.TaskID)
					}
				}
			}(conn)
		}
	}()

	logger.Info("listening on ", s.unixSocket)
	return nil
}

//LaunchTask --
func (s *SchedulerSVC) LaunchTask(t *(communicate.SocketTask)) error {

	if t.Msgobj.ObjType == communicate.ObjTypeShellCmd {
		var srcidbuf [20]byte
		copy(srcidbuf[:], t.Msgobj.SrcID)

		msgBytes, err := t.Msgobj.MarshalMSG()
		if err != nil {
			return err
		}

		p := communicate.Packet{
			PacketNum:  1,
			Index:      0,
			PayloadLen: 0,
			Checksum:   0,
			Payload:    msgBytes,
		}

		packetBytes, err := p.MarshalPacket()
		if err != nil {
			return err
		}

		s.SendMsgChan <- SendMsgChan{
			TaskID:  t.TaskID,
			Message: packetBytes,
			DestIP:  t.Msgobj.DestIP,
		}

	}

	return nil
}

//SendFile -- a func of commNode interface
func (s *SchedulerSVC) SendFile(fileName string, destIPPort string, destPath string, fileMode []byte) error {
	return nil
}

//HandleSendPort -- handle udp sending port
func (s *SchedulerSVC) HandleSendPort() error {
	var err error

	s.SendPC, err = net.ListenPacket("udp", communicate.SCHSENDPORT)
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
			case sm := <-s.SendMsgChan:
				i++
				logger.Debug(i, "send task", sm.TaskID)

				//send back to caller
				dst, err := net.ResolveUDPAddr("udp", sm.DestIP+communicate.XCTRECVPPORT)
				if err != nil {
					logger.Debug("HandleSendPort-ResolveUDPAddr-", err)
					continue
				}

				_, err = s.SendPC.WriteTo(sm.Message, dst)
				if err != nil {
					logger.Debug("HandleSendPort-WriteTo-", err)
					continue
				}

			case <-s.ctx.Done():
				break
			}
		}
	}()

	return nil
}

func getOutFileNames(taskID string) (string, string) {
	stdOutPath := fmt.Sprintf("%sstdout_%s.txt", SCHDLPATH, taskID)
	stdErrPath := fmt.Sprintf("%sstderr_%s.txt", SCHDLPATH, taskID)

	return stdOutPath, stdErrPath
}

func (s *SchedulerSVC) getOutFileHandler(ltype byte, ml *communicate.MSGLine) (*os.File, error) {
	var err error
	var f *os.File

	tidStr := ml.TaskID
	stdOutPath, stdErrPath := getOutFileNames(tidStr)

	logger.Info("got out files:", stdOutPath, stdErrPath)

	if ltype == STDERR {
		f = s.StderrFileMap[tidStr]
		if f == nil {

			f, err = os.OpenFile(stdErrPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return nil, err
			}
			s.StderrFileMap[tidStr] = f
		}
	} else if ltype == STDOUT {
		f = s.StdoutFileMap[tidStr]
		if f == nil {

			f, err = os.OpenFile(stdOutPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return nil, err
			}
			s.StdoutFileMap[tidStr] = f
		}
	}

	return f, nil
}

//WriteOutFile --
func (s *SchedulerSVC) WriteOutFile(ltype byte, ml *communicate.MSGLine) error {

	f, err := s.getOutFileHandler(ltype, ml)
	if err != nil {
		return err
	}

	if _, err := f.WriteString(ml.Line + "\n"); err != nil {
		return err
	}

	if strings.Contains(ml.Line, communicate.EOF) {
		if err := f.Close(); err != nil {
			return err
		}
		if ltype == STDERR {
			delete(s.StderrFileMap, string(ml.TaskID))
		} else if ltype == STDOUT {
			delete(s.StdoutFileMap, string(ml.TaskID))
		}
	}

	return nil
}

//HandleRecvPort -- handle udp receiving port
func (s *SchedulerSVC) HandleRecvPort() error {
	var err error

	s.RecvPC, err = net.ListenPacket("udp", communicate.SCHRECVPPORT)
	if err != nil {
		return err
	}

	//var theTO time.Duration
	//theTO = time.Duration(communicate.DEFAULTTIMEOUT) * time.Second  //todo, need register tasks and trace them

	go func() {
		//ctx, cancel := context.WithTimeout(s.ctx, theTO)
		//defer cancel()

		defer s.RecvPC.Close()

		for {
			buf := make([]byte, communicate.MAXUDPPACKET)
			//n, addr, err := conn.ReadFrom(buf)
			n, _, err := s.RecvPC.ReadFrom(buf)
			if err != nil {
				logger.Error(err)
				continue
			}

			if n > 0 {
				rTL, err := communicate.UnMarshalTL(buf[:n])
				if err != nil {
					logger.Error("UnMarshalTL-", err)
				}

				ml := communicate.MSGLine{}
				if err = ml.UnMarshalMSG(rTL.LineBytes); err != nil {
					logger.Error("UnMarshalMSG-", err)
					continue
				}

				//todo -- write to task return file
				logger.Debug("taskid:", ml.TaskID, "type:", rTL.LineType, "line:", ml.Line)
				if err := s.WriteOutFile(rTL.LineType, &ml); err != nil {
					logger.Error("WriteOutFile-", err)
				}
			}

			select {
			default:
				continue
			//case <-time.After(theTO):
			//	logger.Error("taskid:", taskID, "timeout")
			//	break
			case <-s.ctx.Done():
				break
			}
		}

	}()

	return nil
}

//Init -- initialising key and port.
func (s *SchedulerSVC) Init(ctx context.Context, wg *sync.WaitGroup) error {
	s.privateKey = PVTKEY
	s.unixSocket = SCHUNIXSOCKET
	s.wg = wg
	s.ctx = ctx

	s.StdoutFileMap = make(map[string]*os.File)
	s.StderrFileMap = make(map[string]*os.File)

	err := os.MkdirAll(SCHDLPATH, os.ModePerm)
	if err != nil {
		return nil
	}

	err = s.HandleSendPort()
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
