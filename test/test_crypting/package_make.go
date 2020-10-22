package main

import (
	"fmt"

	"ire.com/clustershell/crypting"
)

func main() {
	fmt.Println("\n", "package_make", "\n")

	//GenerateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey)
	pvtkey, pubkey := crypting.GenerateKeyPair(512)

	fmt.Println("private_key", pvtkey)
	fmt.Println("public_key", pubkey, "\n")

	//PrivateKeyToBytes(priv *rsa.PrivateKey) []byte
	pvtbyte := crypting.PrivateKeyToBytes(pvtkey)
	fmt.Println("private_key_bytes", "\n", string(pvtbyte))
	//PublicKeyToBytes(pub *rsa.PublicKey) []byte
	pubbyte := crypting.PublicKeyToBytes(pubkey)
	fmt.Println("public_key_bytes", "\n", string(pubbyte), "\n")

	//BytesToPrivateKey(priv []byte) *rsa.PrivateKey
	pvtkey2 := crypting.BytesToPrivateKey(pvtbyte)
	fmt.Println("private_key2", pvtkey2)
	//BytesToPublicKey(pub []byte) *rsa.PublicKey
	pubkey2 := crypting.BytesToPublicKey(pubbyte)
	fmt.Println("public_key2", pubkey2, "\n")

	fmt.Println("private_key", pvtkey)
	fmt.Println("public_key", pubkey)

	//EncryptWithPublicKey(msg []byte, pub *rsa.PublicKey) []byte
	msg := []byte("Hello World")
	ciphertext := crypting.EncryptWithPublicKey(msg, pubkey)
	fmt.Println("ciphertext", string(ciphertext))
	//DecryptWithPrivateKey(ciphertext []byte, priv *rsa.PrivateKey) []byte
	plaintext := crypting.DecryptWithPrivateKey(ciphertext, pvtkey)
	fmt.Println("plaintext", string(plaintext))
}
