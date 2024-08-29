package msgcrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// Encryptor implements Encryptor interface using AES algorithm
type Encryptor struct {
	key []byte
}

// NewAesEncryptor creates a new AesEncryptor with a given key
func NewAesEncryptor(key []byte) *Encryptor {
	return &Encryptor{
		key: key,
	}
}

// Encrypt encrypts plain text using AES algorithm
func (e *Encryptor) Encrypt(text string) (string, error) {
	block, err := aes.NewCipher(e.key)
	if err != nil {
		return "", err
	}

	plainTextBytes := []byte(text)
	cipherText := make([]byte, aes.BlockSize+len(plainTextBytes))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainTextBytes)

	encryptedText := base64.StdEncoding.EncodeToString(cipherText)

	return encryptedText, nil
}

// Decrypt decrypts the base64 encoded encrypted text using AES algorithm
func (e *Encryptor) Decrypt(text string) (string, error) {
	cipherText, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(e.key)
	if err != nil {
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherTextBytes := cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherTextBytes, cipherTextBytes)

	return string(cipherTextBytes), nil
}

// GenerateKey generates a random AES key of the specified size (16, 24, or 32 bytes)
func (e *Encryptor) GenerateKey(size int) ([]byte, error) {
	if size != 16 && size != 24 && size != 32 {
		return nil, errors.New("invalid key size")
	}

	key := make([]byte, size)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}

	e.key = key

	return key, nil
}
