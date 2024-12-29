package resp

import (
	"bufio"
	"io"
)

const (
	ERROR   = '-'
	NULL    = 'n'
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

type Resp struct {
	reader *bufio.Reader
}

func NewResp(reader io.Reader) *Resp {
	return &Resp{reader: bufio.NewReader(reader)}
}