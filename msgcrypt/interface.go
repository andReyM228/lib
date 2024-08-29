package msgcrypt

type MsgCrypt interface {
	Encrypt(text string) (string, error)
	Decrypt(text string) (string, error)
	GenerateKey(size int) ([]byte, error)
}
