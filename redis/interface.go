package redis

type Redis interface {
	Set(key string, value interface{}) error
	GetBytes(key string) ([]byte, error)
	GetString(key string) (string, error)
}
