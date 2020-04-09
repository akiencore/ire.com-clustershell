package nodesvcs

import (
	"fmt"
	"net"
	"os"
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

	SCHUNIXSOCKET   = `/var/run/ire_clshsheduler.sock`
	SOCKETMSGMAXLEN = 1024 * 32

	SHCUDPPORT = "0.0.0.0:52233"
)

//SchedulerSVC --
type SchedulerSVC struct {
	privateKey string
	unixSocket string
	udpPort    string

	wg *sync.WaitGroup
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

//HandleSendMsg -- a func of commNode interface
//handle returned message after invoking SendMsg
func (s *SchedulerSVC) HandleSendMsg() error {
	return nil
}

//HandleListenOnUnixSocket -- a func of commNode interface
//handle recerived bytes stream on local unix domain socket
func (s *SchedulerSVC) HandleListenOnUnixSocket() error {
	return nil
}

//HandleListenOnUDP -- a func of commNode interface
//handle recerived bytes stream on UDP port
func (s *SchedulerSVC) HandleListenOnUDP() error {
	return nil
}

//ListenOnUDP -- a func of commNode interface
//lauch a continuous listening on UDP port
func (s *SchedulerSVC) ListenOnUDP() error {

	return nil
}

//ListenOnUnixSocket -- a func of commNode interface
//lauch a continuous listening on UNIX domain socket
func (s *SchedulerSVC) ListenOnUnixSocket() error {
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
					buf := make([]byte, SOCKETMSGMAXLEN+1)
					n, uaddr, err := conn.ReadFromUnix(buf)
					if err != nil {
						logger.Debug("unix socket ReadFromUnix:", err)
						break
					}
					logger.Debug("ListenOnUnixSocket: received", n, "bytes from", uaddr)
					logger.Debug("ListenOnUnixSocket:", string(buf))

					//if n > SOCKETMSGMAXLEN {
					//todo
					//}

					err = communicate.GetMSGListFromSocketBuf(buf[:n], conn)
					if err != nil {
						logger.Debug("GetMSGListFromSocketBuf:", err)
						break
					}
					logger.Debug("GetMSGListFromSocketBuf got map len", len(communicate.SocketTasksMap))

					for tid, task := range communicate.SocketTasksMap {
						logger.Debug("tid=", tid, "task=", task)
					}

				}
			}(conn)

		}

	}()

	logger.Info("listening on ", s.unixSocket)
	return nil
}

//SendFile -- a func of commNode interface
func (s *SchedulerSVC) SendFile(fileName string, destIPPort string, destPath string, fileMode []byte) error {
	return nil
}

//SendMsg -- a func of commNode interface
func (s *SchedulerSVC) SendMsg(message []byte, destIPPort string) error {
	sendPort, err := net.ListenPacket("udp", ":0")
	defer sendPort.Close()

	if err != nil {
		logger.Error(err)
	}

	dst, err := net.ResolveUDPAddr("udp", destIPPort)
	if err != nil {
		return err
	}

	var srcidbuf [20]byte
	copy(srcidbuf[:], []byte("SCHEDULER01"))
	p := communicate.Packet{
		SrcID:   srcidbuf,
		TaskID:  communicate.TASKID,
		Payload: message,
	}

	data, err := p.MarshalPacket()
	if err != nil {
		return err
	}

	_, err = sendPort.WriteTo(data, dst)
	if err != nil {
		return err
	}

	return nil
}

//Init -- initialising key and port.
func (s *SchedulerSVC) Init(wg *sync.WaitGroup) error {
	s.privateKey = PVTKEY
	s.unixSocket = SCHUNIXSOCKET
	s.udpPort = SHCUDPPORT
	s.wg = wg

	return nil
}
