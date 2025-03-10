package sym

import (
    "testing"
    "github.com/stretchr/testify/assert"
	"github.com/encrypter"
	"fmt"
)

func TestEncrpytion(t *testing.T) {
    var ed encrypter.EncryptDecrypt = &SymmetricEncrypter{}

	encryptionKey := "6368616e676520746869732070617373776f726420746f206120736563726574"
	plainText := "This is a sensitive information"

	ed.SetEncryptionKey(encryptionKey)
	ed.SetPlainText(plainText)

	encryptedText, err := ed.Encrypt()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	decryptedText, err := ed.Decrypt()
	if err != nil {
		fmt.Println("Error: ", err)
	}

    // This will fail the test if got != want but continue execution.
	assert.Equal(t, encryptedText, ed.GetEncryptedText(), "they should be equal")
    assert.Equal(t, plainText, decryptedText, "they should be equal")
}