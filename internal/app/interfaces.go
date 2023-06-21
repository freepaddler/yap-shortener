package app

type Storage interface {
	Put(string) []byte
	Get(string) (string, bool)
}
