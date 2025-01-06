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

type Protocol struct {
	reader *bufio.Reader
	writer io.Writer
}

func NewProtocol(reader io.Reader, writer io.Writer) *Protocol {
	return &Protocol{reader: bufio.NewReader(reader), writer: writer}
}
