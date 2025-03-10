package sym

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
)

type SymmetricEncrypter struct {
	encryptionKey []byte
	plainText     []byte
	encryptedText string
	cipherText    []byte
}

func (s *SymmetricEncrypter) SetEncryptionKey(encryptionKey string) {
	var err error
	s.encryptionKey, err = hex.DecodeString(encryptionKey)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func (s *SymmetricEncrypter) SetPlainText(plainText string) {
	s.plainText = []byte(plainText)
}

func (s SymmetricEncrypter) GetEncryptedText() string {
	return s.encryptedText
}

func (s *SymmetricEncrypter) Encrypt() (string, error) {
	if s.encryptionKey == nil {
		return "", errors.New("encryption key is not set")
	}

	if s.plainText == nil {
		return "", errors.New("plain text is not set")
	}

	block, err := aes.NewCipher(s.encryptionKey)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}
	cipherText := gcm.Seal(nil, nonce, s.plainText, nil)
	s.cipherText = append(nonce, cipherText...)

	s.encryptedText = hex.EncodeToString(s.cipherText)
	return s.encryptedText, nil
}

func (s *SymmetricEncrypter) Decrypt() (string, error) {
	if s.encryptionKey == nil {
		return "", errors.New("encryption key is not set")
	}

	if s.cipherText == nil {
		return "", errors.New("cipher text is not set")
	}

	if s.encryptedText == "" {
		return "", errors.New("encrypted text is not set")
	}

	block, _ := aes.NewCipher(s.encryptionKey)
	gcm, _ := cipher.NewGCM(block)
	nonceSize := gcm.NonceSize()
	if len(s.cipherText) < nonceSize {
		return "", errors.New("cipher Text too short")
	}
	nonce, cipherText := s.cipherText[:nonceSize], s.cipherText[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}
	return string(plainText), nil
}

func (s *SymmetricEncrypter) CheckValue(plainText string, encryptedText string) (bool, error) {
	return false, nil
}
