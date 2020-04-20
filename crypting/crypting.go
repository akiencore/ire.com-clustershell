package crypting

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"

	"ire.com/clustershell/logger"

	cryptorand "crypto/rand"

	chacha "golang.org/x/crypto/chacha20poly1305"
)

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

// GenerateKeyPair generates a new key pair
func GenerateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey) {
	privkey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		logger.Error(err)
	}
	return privkey, &privkey.PublicKey
}

// PrivateKeyToBytes private key to bytes
func PrivateKeyToBytes(priv *rsa.PrivateKey) []byte {
	privBytes := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(priv),
		},
	)

	return privBytes
}

// PublicKeyToBytes public key to bytes
func PublicKeyToBytes(pub *rsa.PublicKey) []byte {
	pubASN1, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		logger.Error(err)
	}

	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})

	return pubBytes
}

// BytesToPrivateKey bytes to private key
func BytesToPrivateKey(priv []byte) *rsa.PrivateKey {
	block, _ := pem.Decode(priv)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		logger.Info("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			logger.Error(err)
		}
	}
	key, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		logger.Error(err)
	}
	return key
}

// BytesToPublicKey bytes to public key
func BytesToPublicKey(pub []byte) *rsa.PublicKey {
	block, _ := pem.Decode(pub)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		logger.Info("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			logger.Error(err)
		}
	}
	ifc, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		logger.Error(err)
	}
	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		logger.Error("not ok")
	}
	return key
}

// EncryptWithPublicKey encrypts data with public key
func EncryptWithPublicKey(msg []byte, pub *rsa.PublicKey) []byte {
	hash := sha1.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pub, msg, nil)
	if err != nil {
		logger.Error(err)
	}
	return ciphertext
}

// DecryptWithPrivateKey decrypts data with private key
func DecryptWithPrivateKey(ciphertext []byte, priv *rsa.PrivateKey) []byte {
	hash := sha1.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, priv, ciphertext, nil)
	if err != nil {
		logger.Error(err)
	}
	return plaintext
}

// Encrypt with public key string instead of rsa.PublicKey, finally return ciphertext and error
func EncryptWithPublicKeyStr(unencrypted []byte, pubkeystr string) ([]byte, error) {
	pubkey := BytesToPublicKey([]byte(pubkeystr))

	hash := sha1.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pubkey, unencrypted, nil)
	if err != nil {
		logger.Error(err)
	}

	return ciphertext, err
}

// Decrypt with private key string instead of rsa.PrivateKey, finally return plaintext and error
func DecryptWithPrivateKeyStr(encrypted []byte, pvtkeystr string) ([]byte, error) {
	pvtkey := BytesToPrivateKey([]byte(pvtkeystr))

	hash := sha1.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, pvtkey, encrypted, nil)
	if err != nil {
		logger.Error(err)
	}

	return plaintext, err
}

// SymmKeyGen generates new symmetric key for sender
func SymmKeyGen(length int) []byte {
	key := make([]byte, KeySize)
	if _, err := cryptorand.Read(key); err != nil {
		panic(err)
	}

	return key
}

// EncryptMsg - encryption with symmetric key, and use rsa public key string to hide that symmetric key
func EncryptMsg(msg []byte, pubkeystr string) ([]byte, []byte) {

	keySymm := SymmKeyGen(KeySize) //generate a random symmetric key as sender

	aead, err := chacha.NewX([]byte(keySymm)) //both sender and receiver do this step
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, aead.NonceSize(), aead.NonceSize()+len(msg)+aead.Overhead())
	if _, err := cryptorand.Read(nonce); err != nil {
		panic(err)
	}

	encryptedMsg := aead.Seal(nonce, nonce, msg, nil) //encrypt message with symmetric key

	encryptedKey, err := EncryptWithPublicKeyStr(keySymm, pubkeystr) //encrypt symmetric key with provided public key
	if err != nil {
		panic(err)
	}

	return encryptedMsg, encryptedKey
}

// DecryptMsg - symmetric decryption with symmetric key, while symmetric key is decrypted wih rsa private key string
func DecryptMsg(encryptedMsg []byte, encryptedKey []byte, pvtkeystr string) ([]byte, error) {

	keySymm, err := DecryptWithPrivateKeyStr(encryptedKey, pvtkeystr) //decrypt symmetric key with provided private key
	if err != nil {
		logger.Error(err)
	}

	aead, err := chacha.NewX([]byte(keySymm)) //both sender and receiver do this step
	if err != nil {
		panic(err)
	}

	if len(encryptedMsg) < aead.NonceSize() {
		panic("ciphertext too short")
	}

	nonce, ciphertext := encryptedMsg[:aead.NonceSize()], encryptedMsg[aead.NonceSize():]

	plaintext, err := aead.Open(nil, nonce, ciphertext, nil) //decrypt message with symmetric key
	if err != nil {
		panic(err)
	}

	return plaintext, err
}
