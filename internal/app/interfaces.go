package app

type Storage interface {
	Put(string) string
	Get(string) (string, bool)
}
