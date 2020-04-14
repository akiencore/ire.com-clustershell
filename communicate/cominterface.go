package communicate

import (
	"context"
	"sync"
)

//Crypter --
/*Scheduler and Executor will commucate with each other in encrypted message by rsa key pairs.
* Schedule encrypts message by private key before sending to Executor, and Executor decrypts
* received messages by public key.
* Conversly, Executor encrypts message by public key before sending to Scheduler, and Scheduler
* decrypts received messages by private key.
 */
type Crypter interface {
	Encrypt(unEncrypted []byte, comKey string) (encrypted []byte, err error)
	Decrpyt(encrypted []byte, comKey string) (unEncrypted []byte, err error)
}

//CommNode -- We use UDP to avoid maintaining a continuous connect which TCP must do.
type CommNode interface {
	Crypter

	//handle recerived bytes stream on local unix domain socket
	HandleUnixSocket() error
	//handle udp sending port
	HandleSendPort() error
	//handle udp receiving port
	HandleRecvPort() error

	//funcs to wrap kinds of message to packet
	//SendFile(fileName string, destIPPort string, destPath string, fileMode []byte) ([]byte, error)
	//SendMsg(taskID TaskIDType, message []byte, destIPPort string) ([]byte, error)

	Init(ctx context.Context, wg *sync.WaitGroup) error
}
