package storage

import (
	"time"
)

func (storage *Storage) Set(
	key string,
	value interface{},
	timestamp time.Time,
) error {
	newStoreWrite := &StoreWrite{
		createdAt: timestamp,
		value:     value,
	}

	storage.mu.Lock()
	(*storage.store)[key] = *newStoreWrite
	storage.mu.Unlock()

	return nil
}
