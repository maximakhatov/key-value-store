package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"

	"github.com/maximakhatov/key-value-store/internal/handlers"
	"github.com/maximakhatov/key-value-store/internal/resp"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	Port     int    `mapstructure:"KV_PORT"`
	LogLevel string `mapstructure:"LOG_LEVEL"`
}

var config Config

func init() {
	config = readConfig()

	var zerologLogLevel zerolog.Level
	switch strings.ToLower(config.LogLevel) {
	case "trace":
		zerologLogLevel = zerolog.TraceLevel
	case "debug":
		zerologLogLevel = zerolog.DebugLevel
	case "info":
		zerologLogLevel = zerolog.InfoLevel
	case "warn":
		zerologLogLevel = zerolog.WarnLevel
	case "error":
		zerologLogLevel = zerolog.ErrorLevel
	case "fatal":
		zerologLogLevel = zerolog.FatalLevel
	case "panic":
		zerologLogLevel = zerolog.PanicLevel
	default:
		zerologLogLevel = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(zerologLogLevel)
}

func main() {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(config.Port))
	if err != nil {
		log.Err(err)
		return
	}
	defer listener.Close()
	log.Info().Int("port", config.Port).Msg("Server started listening")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Err(err)
			return
		}
		log.Info().Msg("Received new connection")
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	protocol := resp.NewProtocol(conn, conn)

	for {
		value, err := protocol.Read()
		if err != nil {
			if err == io.EOF {
				log.Info().Msg("Client disconnected")
			} else {
				log.Err(err).Msg("Error reading from client")
			}
			return
		}

		if value.Type != resp.ARRAY {
			log.Warn().Msg("Invalid request, expected array")
			continue
		}

		if len(value.Array) == 0 {
			log.Warn().Msg("Invalid request, expected array length > 0")
			continue
		}

		command := strings.ToUpper(value.Array[0].Bulk)
		args := value.Array[1:]

		handler, ok := handlers.Handlers[command]
		if !ok {
			log.Warn().Str("command", command).Msg("Invalid command")
			protocol.Write(resp.Value{Type: resp.STRING, Str: ""})
			continue
		}

		result := handler(args)
		err = protocol.Write(result)
		if err != nil {
			log.Err(err).Msg("Error writing to client")
		}
	}
}

func readConfig() Config {
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
	return config
}
