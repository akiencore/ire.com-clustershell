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

//PassiveHandler -- the passive processing functions
type PassiveHandler interface {
	// handle returned message after invoking SendMsg
	HandleSendMsg() error
	//handle recerived bytes stream on local unix domain socket
	HandleListenOnUnixSocket() error
	//handle recerived bytes stream on UDP port
	HandleListenOnUDP(context.Context) error
}

//CommNode -- We use UDP to avoid maintaining a continuous connect which TCP must do.
type CommNode interface {
	Crypter
	PassiveHandler

	//lauch a continuous listening on UDP port
	ListenOnUDP() error
	//lauch a continuous listening on UNIX domain socket
	ListenOnUnixSocket() error

	SendFile(fileName string, destIPPort string, destPath string, fileMode []byte) error
	SendMsg(message []byte, destIPPort string) error

	Init(wg *sync.WaitGroup) error
}
