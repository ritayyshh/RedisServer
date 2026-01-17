package storage

import (
	"sync"
)

type StoreWrite struct {
	value string
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
