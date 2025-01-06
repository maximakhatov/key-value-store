package resp

import (
	"fmt"
	"strconv"
)

func (v Value) marshal() []byte {
	switch v.Type {
	case ARRAY:
		return v.marshalArray()
	case BULK:
		return v.marshalBulk()
	case STRING:
		return v.marshalString()
	case NULL:
		return v.marshallNull()
	case ERROR:
		return v.marshallError()
	default:
		fmt.Println("Marshalling fail, unexpected type", v.Type)
		return []byte{}
	}
}

func (v Value) marshalString() []byte {
	var bytes []byte = make([]byte, 0, len(v.Str)+3)
	bytes = append(bytes, STRING)
	bytes = append(bytes, v.Str...)
	bytes = append(bytes, '\r', '\n')

	return bytes
}

func (v Value) marshalBulk() []byte {
	var bytes []byte
	bytes = append(bytes, BULK)
	bytes = append(bytes, strconv.Itoa(len(v.Bulk))...)
	bytes = append(bytes, '\r', '\n')
	bytes = append(bytes, v.Bulk...)
	bytes = append(bytes, '\r', '\n')

	return bytes
}

func (v Value) marshalArray() []byte {
	len := len(v.Array)
	var bytes []byte
	bytes = append(bytes, ARRAY)
	bytes = append(bytes, strconv.Itoa(len)...)
	bytes = append(bytes, '\r', '\n')

	for i := 0; i < len; i++ {
		bytes = append(bytes, v.Array[i].marshal()...)
	}

	return bytes
}

func (v Value) marshallError() []byte {
	var bytes []byte = make([]byte, 0, len(v.Str)+3)
	bytes = append(bytes, ERROR)
	bytes = append(bytes, v.Str...)
	bytes = append(bytes, '\r', '\n')

	return bytes
}

func (v Value) marshallNull() []byte {
	return []byte("$-1\r\n")
}

func (w *Protocol) Write(v Value) error {
	var bytes = v.marshal()

	_, err := w.writer.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}
