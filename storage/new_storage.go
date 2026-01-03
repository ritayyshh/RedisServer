package storage

import (
	"sync"
	"time"
)

type StoreWrite struct {
	createdAt time.Time
	value     interface{}
}

type Storage struct {
	mu    sync.Mutex
	store *map[string]StoreWrite
}

func NewStorage() *Storage {
	newStore := make(map[string]StoreWrite)

	newStorage := &Storage{
		store: &newStore,
	}

	return newStorage
}
