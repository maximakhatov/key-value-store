package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/maximakhatov/key-value-store/internal/handlers"
	"github.com/maximakhatov/key-value-store/internal/resp"

	"github.com/spf13/viper"
)

type Config struct {
	Port int `mapstructure:"KV_PORT"`
}

// TODO replace fmt with logging
func main() {
	config := readConfig()

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(config.Port))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Listening on port", config.Port)

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Received new connection")

	defer conn.Close()

	for {
		r := resp.NewResp(conn)
		value, err := r.Read()
		if err != nil {
			fmt.Println("Error reading from client:", err.Error())
			return
		}

		if value.Type != resp.ARRAY {
			fmt.Println("Invalid request, expected array")
			continue
		}

		if len(value.Array) == 0 {
			fmt.Println("Invalid request, expected array length > 0")
			continue
		}

		command := strings.ToUpper(value.Array[0].Bulk)
		args := value.Array[1:]

		writer := resp.NewWriter(conn)
		handler, ok := handlers.Handlers[command]
		if !ok {
			fmt.Println("Invalid command: ", command)
			writer.Write(resp.Value{Type: resp.STRING, Str: ""})
			continue
		}

		result := handler(args)
		writer.Write(result)
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
