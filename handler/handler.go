package handler

import (
	"sync"

	"github.com/ritayyshh/RedisServer/resp"
	"github.com/ritayyshh/RedisServer/storage"
)

func ping(args []resp.Value) resp.Value {
	return resp.Value{Typ: "string", Str: "PONG"}
}

var store = storage.NewStorage()

var HSETs = map[string]map[string]string{}
var HSETsMu = sync.RWMutex{}

func hset(args []resp.Value) resp.Value {
	if len(args) != 3 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'hset' command"}
	}

	hash := args[0].Bulk
	key := args[1].Bulk
	value := args[2].Bulk

	HSETsMu.Lock()
	if _, ok := HSETs[hash]; !ok {
		HSETs[hash] = map[string]string{}
	}
	HSETs[hash][key] = value
	HSETsMu.Unlock()

	return resp.Value{Typ: "string", Str: "OK"}
}

func hget(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'hget' command"}
	}

	hash := args[0].Bulk
	key := args[1].Bulk

	HSETsMu.RLock()
	value, ok := HSETs[hash][key]
	HSETsMu.RUnlock()

	if !ok {
		return resp.Value{Typ: "null"}
	}

	return resp.Value{Typ: "bulk", Bulk: value}
}

func hgetall(args []resp.Value) resp.Value {
	if len(args) != 1 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'hgetall' command"}
	}

	hash := args[0].Bulk

	HSETsMu.RLock()
	hashMap, ok := HSETs[hash]
	HSETsMu.RUnlock()

	if !ok {
		return resp.Value{Typ: "array", Array: []resp.Value{}}
	}

	// Build array of key-value pairs
	values := []resp.Value{}
	for key, value := range hashMap {
		values = append(values, resp.Value{Typ: "bulk", Bulk: key})
		values = append(values, resp.Value{Typ: "bulk", Bulk: value})
	}

	return resp.Value{Typ: "array", Array: values}
}

var Handlers = map[string]func([]resp.Value) resp.Value{
	"PING":    ping,
	"SET":     store.Set,
	"GET":     store.Get,
	"HSET":    hset,
	"HGET":    hget,
	"HGETALL": hgetall,
}
