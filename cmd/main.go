package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

func main() {
	port := 6379     // TODO read from config
	buffSize := 1024 // TODO read from config

	listener, err := net.Listen("tcp", strconv.Itoa(port))
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	for {
		// TODO add logging
		buf := make([]byte, buffSize)

		_, err = conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("error reading from client:", err.Error())
			os.Exit(1)
		}

		conn.Write([]byte("+OK\r\n"))
	}
}
