package client

import (
	"fmt"
	"net"

	"github.com/maximakhatov/key-value-store/internal/resp"
)

type Client struct {
	conn     net.Conn
	protocol *resp.Protocol
}

func New(addr string) (*Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &Client{conn: conn, protocol: resp.NewProtocol(conn, conn)}, nil
}

func (client *Client) Close() {
	client.conn.Close()
}

func (client *Client) Set(key, value string) error {
	command := resp.Value{
		Type: resp.ARRAY,
		Array: []resp.Value{
			{Type: resp.BULK, Bulk: "SET"},
			{Type: resp.BULK, Bulk: key},
			{Type: resp.BULK, Bulk: value},
		},
	}
	err := client.protocol.Write(command)
	if err != nil {
		return err
	}

	response, err := client.protocol.Read()
	if err != nil {
		return err
	}
	if response.Type != resp.STRING || response.Str != "OK" {
		return fmt.Errorf("server returned non-OK response: %v", response)
	}
	return nil
}

func (client *Client) Get(key string) (result string, null bool, e error) {
	command := resp.Value{
		Type: resp.ARRAY,
		Array: []resp.Value{
			{Type: resp.BULK, Bulk: "GET"},
			{Type: resp.BULK, Bulk: key},
		},
	}
	err := client.protocol.Write(command)
	if err != nil {
		return "", false, err
	}

	response, err := client.protocol.Read()
	if err != nil {
		return "", false, err
	}
	if response.Type != resp.BULK && response.Type != resp.NULL {
		return "", false, fmt.Errorf("server returned unexpected response: %v", response)
	}
	return response.Bulk, response.Type == resp.NULL, nil
}
