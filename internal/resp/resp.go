package resp

import (
	"bufio"
	"io"
)

const (
	ERROR   = '-'
	NULL    = '0'
	STRING  = '+'
	INTEGER = ':'
	BULK    = '$'
	ARRAY   = '*'
)

type Value struct {
	Type  rune
	Str   string
	Num   int
	Bulk  string
	Array []Value
}

type RESP struct {
	reader *bufio.Reader
}

func NewResp(reader io.Reader) *RESP {
	return &RESP{reader: bufio.NewReader(reader)}
}
