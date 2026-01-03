package storage

func (storage *Storage) Get(key string) interface{} {
	storage.mu.Lock()
	storeWrite, ok := (*storage.store)[key] //store[key]
	storage.mu.Unlock()

	if !ok {
		return nil
	}

	return storeWrite.value
}
