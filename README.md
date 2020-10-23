## crypting in rsa key and symmetric key

keys.go: 
* after input "make" in clustershell folder, there are 2 pairs of rsa keys (saved as string).
-   For Scheduler: PubkeyScheduler, PvtkeyScheduler
-   For Executor: PubkeyExecutor, PvtkeyExecutor

****

crypting.go:
* provide methods to do crypting
functions: 
-   GenerateKeyPair(int) (*rsa.PrivateKey, *rsa.PublicKey): return rsa privatekey and publickey according to provided length
-   PrivateKeyToBytes(*rsa.PrivateKey) []byte: convert rsa.PrivateKey to keybytes
-   PublicKeyToBytes(*rsa.PublicKey) []byte: convert rsa.PublicKey to keybytes
-   BytesToPrivateKey([]byte) *rsa.PrivateKey: convert keybytes to rsa.PrivateKey
-   BytesToPublicKey([]byte) *rsa.PublicKey: convert keybytes to rsa.PublicKey

-   EncryptWithPublicKey([]byte, *rsa.PublicKey) []byte: encrypt message with rsa public key, return ciphertext
-   DecryptWithPrivateKey([]byte, *rsa.PrivateKey) []byte: decrypt message with rsa private key, return plaintext

-   EncryptWithPublicKeyStr([]byte, string) ([]byte, error): encrypt message with rsa public key string, return ciphertext and error
-   DecryptWithPrivateKeyStr([]byte, string) ([]byte, error): encrypt message with rsa private key string, return plaintext and error

-   SymmKeyGen(int) []byte: return symmetric key according to provided length
-   EncryptMsg([]byte, string) ([]byte, []byte): encrypt message with randomly generated symmetric key, return encrypted message bytes and encrypted symmetric key that encrypted by provided rsa public key
-   DecryptMsg([]byte, []byte, string) ([]byte, error): decrypt symmetric key with provided rsa private key, and use this symmetric key to decrypt encrypted message, return plaintext and error
