package encrypter

type EncryptDecrypt interface {
	SetEncryptionKey(encryptionKey string)
	SetPlainText(plainText string)
	GetEncryptedText() string
	Encrypt() (string, error)
	Decrypt() (string, error)
	CheckValue(plainText string, encryptedText string) (bool, error)
}
