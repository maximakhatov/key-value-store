package handlers

import (
	"sync"

	"github.com/maximakhatov/key-value-store/internal/resp"
)

var hstorage = map[string]map[string]string{}
var hstorageMutex = sync.RWMutex{}

func HSet(args []resp.Value) resp.Value {
	if len(args) != 3 {
		return resp.Value{Type: resp.ERROR, Str: "ERR wrong number of arguments for 'hset' command"}
	}

	hash := args[0].Bulk
	key := args[1].Bulk
	value := args[2].Bulk

	hstorageMutex.Lock()
	if _, ok := hstorage[hash]; !ok {
		hstorage[hash] = map[string]string{}
	}
	hstorage[hash][key] = value
	hstorageMutex.Unlock()

	return resp.Value{Type: resp.STRING, Str: "OK"}
}

func HGet(args []resp.Value) resp.Value {
	if len(args) != 2 {
		return resp.Value{Type: resp.ERROR, Str: "ERR wrong number of arguments for 'hget' command"}
	}

	hash := args[0].Bulk
	key := args[1].Bulk

	hstorageMutex.RLock()
	value, ok := hstorage[hash][key]
	hstorageMutex.RUnlock()

	if !ok {
		return resp.Value{Type: resp.NULL}
	}

	return resp.Value{Type: resp.BULK, Bulk: value}
}
