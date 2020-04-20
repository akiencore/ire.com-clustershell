package main

import (
	"fmt"

	"ire.com/clustershell/crypting"
	"ire.com/clustershell/crypting/keypairs"

	cryptorand "crypto/rand"

	chacha "golang.org/x/crypto/chacha20poly1305"
)

// Test collection of key strings
type Test struct {
	pvtkeystrScheduler string
	pubkeystrScheduler string
	pvtkeystrExecutor  string
	pubkeystrExecutor  string
}

const (
	// KeySize is the size of the key used by this AEAD, in bytes.
	KeySize = 32

	// NonceSize is the size of the nonce used with the standard variant of this
	// AEAD, in bytes.
	NonceSize = 12

	// NonceSizeX is the size of the nonce used with the XChaCha20-Poly1305
	// variant of this AEAD, in bytes.
	NonceSizeX = 24
)

func main() {
	fmt.Println("\n", "kengen_symm", "\n")

	test := Test{keypairs.PvtkeyScheduler, keypairs.PubkeyScheduler, keypairs.PvtkeyExecutor, keypairs.PubkeyExecutor}

	var enkey []byte

	// Encryption.
	var encryptedMsg []byte
	{
		key := make([]byte, KeySize)
		if _, err := cryptorand.Read(key); err != nil {
			panic(err)
		}

		aead, err := chacha.NewX(key) //both sender and receiver do this step
		if err != nil {
			panic(err)
		}

		msg := []byte("Gophers, gophers, gophers everywhere!")

		// Select a random nonce, and leave capacity for the ciphertext.
		nonce := make([]byte, aead.NonceSize(), aead.NonceSize()+len(msg)+aead.Overhead())
		if _, err := cryptorand.Read(nonce); err != nil {
			panic(err)
		}
		// Encrypt the message and append the ciphertext to the nonce.
		encryptedMsg = aead.Seal(nonce, nonce, msg, nil)

		//asymmetric encrypt
		enkey, err = crypting.EncryptWithPublicKeyStr(key, test.pubkeystrScheduler)
		if err != nil {
			panic(err)
		}
		fmt.Println(enkey)
	}
	
	// Decryption.
	{
		//asymmetric decrypt
		key, err := crypting.DecryptWithPrivateKeyStr(enkey, test.pvtkeystrScheduler)
		if err != nil {
			panic(err)
		}

		aead, err := chacha.NewX(key) //both sender and receiver do this step

		if err != nil {
			panic(err)
		}

		if len(encryptedMsg) < aead.NonceSize() {
			panic("ciphertext too short")
		}

		// Split nonce and ciphertext.
		nonce, ciphertext := encryptedMsg[:aead.NonceSize()], encryptedMsg[aead.NonceSize():]

		// Decrypt the message and check it wasn't tampered with.
		plaintext, err := aead.Open(nil, nonce, ciphertext, nil)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s\n", plaintext)
	}
	
	bs := make([]byte, 10)
	cryptorand.Read(bs);
	fmt.Println(bs)

	cs := make([]byte, 10)
	cryptorand.Read(cs);
	fmt.Println(cs)
}
