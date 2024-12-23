package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	Port       int `mapstructure:"KV_PORT"`
	BufferSize int `mapstructure:"KV_BUFFER_SIZE"`
}

func main() {
	config := readConfig()

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(config.Port))
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
		buf := make([]byte, config.BufferSize)

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

func readConfig() *Config {
	var config Config
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error reading config file: %w", err))
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("fatal error unmarshalling config: %w", err))
	}
	return &config
}
