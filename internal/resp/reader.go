package resp

import (
	"strconv"

	"github.com/rs/zerolog/log"
)

func (r *Protocol) Read() (Value, error) {
	_type, err := r.reader.ReadByte()

	if err != nil {
		return Value{}, err
	}

	switch _type {
	case BULK:
		return r.readBulk()
	case ARRAY:
		return r.readArray()
	case STRING:
		return r.readString()
	default:
		log.Error().Str("type", string(_type)).Msg("Unknown type")
		return Value{}, nil
	}
}

func (r *Protocol) readLine() (line []byte, n int, err error) {
	for {
		b, err := r.reader.ReadByte()
		if err != nil {
			return nil, 0, err
		}
		n += 1
		line = append(line, b)
		if len(line) >= 2 && line[len(line)-2] == '\r' {
			break
		}
	}
	return line[:len(line)-2], n, nil
}

func (r *Protocol) readInteger() (x int, n int, err error) {
	line, n, err := r.readLine()
	if err != nil {
		return 0, 0, err
	}
	num, err := strconv.ParseInt(string(line), 10, 64)
	if err != nil {
		return 0, n, err
	}
	return int(num), n, nil
}

func (r *Protocol) readBulk() (Value, error) {
	v := Value{}

	len, _, err := r.readInteger()
	if err != nil {
		return v, err
	}
	if len == -1 {
		v.Type = NULL
		return v, nil
	}
	v.Type = BULK

	bulk := make([]byte, len)
	r.reader.Read(bulk)
	v.Bulk = string(bulk)
	r.readLine()

	return v, nil
}

func (r *Protocol) readArray() (Value, error) {
	v := Value{}
	v.Type = ARRAY

	length, _, err := r.readInteger()
	if err != nil {
		return v, err
	}

	v.Array = make([]Value, length)
	for i := 0; i < length; i++ {
		val, err := r.Read()
		if err != nil {
			return v, err
		}
		v.Array[i] = val
	}

	return v, nil
}

func (r *Protocol) readString() (Value, error) {
	v := Value{Type: STRING}
	line, _, err := r.readLine()
	if err != nil {
		return v, err
	}
	v.Str = string(line)
	return v, nil
}
