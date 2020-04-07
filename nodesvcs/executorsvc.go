package nodesvcs

//PUBKEY -- the public key of key pairs for communication between scheduler and executor
//generate a key pair by following command in linux:
//ssh-keygen -t rsa -P "" -C "commNode@clustershell.ire.com" -f /tmp/clustershell
//cat /tmp/clustershell ### copy the output and paste to PUBKEY, and hide the private key /tmp/clustershell
const (
	PUBKEY = `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDc0QK2LqxlU6StKJ7/3qZ+A5M4w2vaaDuub5BO5v+PPMn65CChdztoK9/QSermpTmc+rjSv2UJJSHqztRxr01DUehBuoOgT7LSh0GMm2HyobMpgVG6v/b8k2S//AkuxmltxkFzgUApBGsWWcgbn88HQLc3QF4P/60It9EMSwy7GEppdsvhGvfS5UyKqDD2Xuz8zHKP5WU1/lt4f1m44K1ZgZTzQv8r53des37m/TmVrYX1vBEflHSC7xTWF37WUXpET+jBru69zUoiM2jzbw89BX2i+S9fbDZTb5OAHVbnevQTNBmkt1D0uwcgur0Xw4hS/Vr/fpt2i9F6Askcougr commNode@clustershell.ire.com`

	XCTUNIXSOCKET = `/run/ire_clshexecutor.sock`

	XCTUDPPORT = "0.0.0.0:33225"
)

//ExecutorSVC --
type ExecutorSVC struct {
	publicKey  string
	unixSocket string
	udpPort    string
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
func (s *ExecutorSVC) HandleListenOnUDP() error {
	return nil
}

//ListenOnUDP -- a func of commNode interface
//lauch a continuous listening on UDP port
func (s *ExecutorSVC) ListenOnUDP(port string) error {
	return nil
}

//ListenOnUnixSocket -- a func of commNode interface
//lauch a continuous listening on UNIX domain socket
func (s *ExecutorSVC) ListenOnUnixSocket(unixSocket string) error {
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
func (s *ExecutorSVC) Init() error {
	s.publicKey = PUBKEY
	s.unixSocket = XCTUNIXSOCKET
	s.udpPort = XCTUDPPORT

	return nil
}
