package inMemmoryCashe

import "sync"

type SessionStorage struct {
	mu      sync.RWMutex
	session map[int64]string
}

func New() {
}
