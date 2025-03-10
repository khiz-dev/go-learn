package main

import (
	"fmt"
	"github.com/encrypter"
	"github.com/sym"
)

func main() {
	var ed encrypter.EncryptDecrypt = &sym.SymmetricEncrypter{}

	encryptionKey := "6368616e676520746869732070617373776f726420746f206120736563726574"
	plainText := "This is a sensitive information"

	ed.SetEncryptionKey(encryptionKey)
	ed.SetPlainText(plainText)
	fmt.Println("Encryption key: ", encryptionKey)
	fmt.Println("Plain text: ", plainText)

	encryptedText, err := ed.Encrypt()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Encrypted text: ", encryptedText)

	decryptedText, err := ed.Decrypt()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Decrypted text: ", decryptedText)
}
