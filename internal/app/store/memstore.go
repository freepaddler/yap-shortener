package store

import (
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
	h := shorter.Short(8)
	s.repo[h] = url
	return []byte(h)
}

func (s *MemStore) Get(short string) (string, bool) {
	if v, ok := s.repo[short]; ok {
		return v, true
	}
	return "", false
}
