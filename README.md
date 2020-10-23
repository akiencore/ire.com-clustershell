## A rsa crypting package for confidential information

This is the extension、outsourcing of another project: https://github.com/ChengWu-NJ/clustershell

I wrote the package in directory crypting/ to improve it's confidentiality.

</br>
* work log: https://www.yuque.com/akiencore/qgdgg5/sq9rbl

****
# About the package

* Purpose: apply rsa public key/private key and symmetric key to ensure those messages of communication between schedulers and executors are well encrypted, only receivers can decrypt them

* Programming Language: golang

****
# Idea
* symmetric key: generate a random symmetric key, and use the same key to encrypt and decrypt. In this module, we use package chacha20poly1305(https://godoc.org/golang.org/x/crypto/chacha20poly1305)  and "crypto/rand" to finish this step. The symmetric key here is to encrypt/decrypt those message that nodes attempt to send or receive.  
* rsa keypairs: generate a pair of rsa keys as public key and private key to encrypt and decrypt. In this module, we first generate a random rsa keypair strings in keys.go, and use functions in crypting.go to handle those strings. The rsa keypairs here are to encrypt/decrypt the same symmetric key above to keep that key from being cracked by unknown hackers. 
* The sender and receiver are corresponding to "scheduler" or "executor" in the project, so that there are 2 pairs of rsa keys for both characters.

****
# Content

</br>
keys.go: after input "make" in clustershell folder, there are 2 pairs of rsa keys (saved as string).
-   For Scheduler: PubkeyScheduler, PvtkeyScheduler
-   For Executor: PubkeyExecutor, PvtkeyExecutor

crypting.go: provide methods to do crypting functions
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
