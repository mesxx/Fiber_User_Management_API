package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"os"
)

func EncryptText(text string) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")
	if len(secretKey) < 32 {
		return "", errors.New("secretKey must be 32 bytes")
	} else if len(secretKey) > 32 {
		secretKey = secretKey[:32]
	}

	block, err1 := aes.NewCipher([]byte(secretKey))
	if err1 != nil {
		return "", err1
	}

	ciphertext := make([]byte, aes.BlockSize+len([]byte(text)))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(text))

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

func DecryptText(hash string) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")
	if len(secretKey) < 32 {
		return "", errors.New("secretKey must be 32 bytes")
	} else if len(secretKey) > 32 {
		secretKey = secretKey[:32]
	}

	key := []byte(secretKey)
	block, err1 := aes.NewCipher(key)
	if err1 != nil {
		return "", err1
	}

	ciphertextBytes, err2 := base64.URLEncoding.DecodeString(hash)
	if err2 != nil {
		return "", err2
	}

	if len(ciphertextBytes) < aes.BlockSize {
		return "", errors.New("hash too short")
	}

	iv := ciphertextBytes[:aes.BlockSize]
	ciphertextBytes = ciphertextBytes[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertextBytes, ciphertextBytes)

	return string(ciphertextBytes), nil
}
