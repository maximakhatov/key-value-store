package handlers

import "github.com/maximakhatov/key-value-store/internal/resp"

var Handlers = map[string]func([]resp.Value) resp.Value{
	"PING": ping,
	"GET":  get,
	"SET":  set,
	"HGET": hget,
	"HSET": hset,
}
