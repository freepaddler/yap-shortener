package shorter

import (
	"math/rand"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Short(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Int63()%int64(len(letters))]
	}
	return string(b)
}
