package handlers

import (
	"sync"

	"github.com/maximakhatov/key-value-store/internal/resp"
)

var storage = map[string]string{}
var storageMutex = sync.RWMutex{}

func get(args []resp.Value) resp.Value {
	if len(args) != 1 {
		return resp.Value{Type: resp.ERROR, Str: "ERR wrong number of arguments for 'get' command"}
	}

	key := args[0].Bulk

	storageMutex.RLock()
	value, ok := storage[key]
	storageMutex.RUnlock()

	if !ok {
		return resp.Value{Type: resp.NULL}
	}

	return resp.Value{Type: resp.BULK, Bulk: value}
}

func set(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Type: resp.ERROR, Str: "ERR wrong number of arguments for 'set' command"}
	}

	key := args[0].Bulk
	value := args[1].Bulk

	storageMutex.Lock()
	storage[key] = value
	storageMutex.Unlock()

	return resp.Value{Type: resp.STRING, Str: "OK"}
}
