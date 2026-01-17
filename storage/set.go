package storage

import (
	"github.com/ritayyshh/RedisServer/resp"
)

func (storage *Storage) Set(
	args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'set' command"}
	}

	key := args[0].Bulk
	value := args[1].Bulk

	newStoreWrite := &StoreWrite{
		value: value,
	}

	storage.mu.Lock()
	(*storage.store)[key] = *newStoreWrite
	storage.mu.Unlock()

	return resp.Value{Typ: "string", Str: "OK"}
}
