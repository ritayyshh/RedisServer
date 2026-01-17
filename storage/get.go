package storage

import (
	"github.com/ritayyshh/RedisServer/resp"
)

func (storage *Storage) Get(args []resp.Value) resp.Value {
	if len(args) != 1 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'get' command"}
	}

	key := args[0].Bulk

	storage.mu.Lock()
	storeWrite, ok := (*storage.store)[key] //store[key]
	storage.mu.Unlock()

	if !ok {
		return resp.Value{Typ: "null"}
	}

	return resp.Value{Typ: "bulk", Bulk: storeWrite.value}
}
