package resp

import (
	"bytes"
	"strings"
	"testing"

	"github.com/go-test/deep"
)

func TestFailUnknownType(t *testing.T) {
	reader := strings.NewReader("? Unknown type")
	writer := bytes.NewBuffer([]byte{})
	protocol := NewProtocol(reader, writer)
	_, err := protocol.Read()
	if err == nil {
		t.Error("Expected error")
	}
	if err.Error() != "unknown type" {
		t.Errorf("Unexpected error %v", err)
	}
}

func TestReadString(t *testing.T) {
	reader := strings.NewReader("+String\r\n")
	writer := bytes.NewBuffer([]byte{})
	protocol := NewProtocol(reader, writer)
	result, err := protocol.Read()
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	expected := Value{Type: STRING, Str: "String"}
	if diff := deep.Equal(result, expected); diff != nil {
		t.Errorf("Unexpected result, diff: %v", diff)
	}
}
