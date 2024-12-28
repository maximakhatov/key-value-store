package handlers

import "github.com/maximakhatov/key-value-store/internal/resp"

func ping(args []resp.Value) resp.Value {
	if len(args) == 0 {
		return resp.Value{Type: resp.STRING, Str: "PONG"}
	}

	return resp.Value{Type: resp.STRING, Str: args[0].Bulk}
}
