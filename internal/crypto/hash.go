package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
	"os"
)

// Encrypt plaintext to a safe encryption using the encryption key provided at runtime
func Encrypt(plainText string) string {
	block, _ := aes.NewCipher([]byte(os.Getenv("PASSWORD_ENCRYPTION_KEY")))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	ciphertextByte := gcm.Seal(nonce, nonce, []byte(plainText), nil)
	return base64.StdEncoding.EncodeToString(ciphertextByte)
}

// Decrypt cipherText back to plaintext using the encryption key provided at runtime
func Decrypt(cipherText string) string {
	// prepare cipher
	keyByte := []byte(os.Getenv("PASSWORD_ENCRYPTION_KEY"))
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	nonceSize := gcm.NonceSize()
	ciphertextByte, _ := base64.StdEncoding.DecodeString(cipherText)
	nonce, ciphertextByteClean := ciphertextByte[:nonceSize], ciphertextByte[nonceSize:]
	plaintextByte, err := gcm.Open(nil, nonce, ciphertextByteClean, nil)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	plainText := string(plaintextByte)
	return plainText
}
