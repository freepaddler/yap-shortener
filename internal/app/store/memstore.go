package store

import (
	"fmt"

	"github.com/freepaddler/yap-shortener/internal/app/shortener"
)

type MemStore struct {
	repo map[string]string
}

func NewMemStore() *MemStore {
	return &MemStore{
		repo: make(map[string]string),
	}
}

func (s *MemStore) Put(url string) []byte {
	fmt.Println("MemStore PUT:", url)
	h := shorter.Short(8)
	s.repo[h] = url
	fmt.Println("MemStore PUT:", h)
	return []byte(h)
}

func (s *MemStore) Get(short string) (string, bool) {
	fmt.Println("MemStore GET:", short)
	if v, ok := s.repo[short]; ok {
		fmt.Println("MemStore GET:", v)
		return v, true
	}
	return "", false
}
