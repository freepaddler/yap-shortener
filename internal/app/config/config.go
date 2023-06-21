package config

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/caarlos0/env/v8"
	flag "github.com/spf13/pflag"
)

const (
	defaultServerAddress = "127.0.0.1:8080"
	defaultServerURL     = "http://127.0.0.1:8080"
)

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS"`
	ServerURL     string `env:"BASE_URL"`
}

func NewConfig() *Config {
	var c Config
	var output io.Writer = os.Stderr
	flag.CommandLine.SetOutput(output)
	// sorting is based on long args, doesn't look too good
	flag.CommandLine.SortFlags = false
	// avoid message "pflag: help requested"
	flag.ErrHelp = errors.New("")

	// cmd params
	flag.StringVarP(
		&c.ServerAddress,
		"serverAddress",
		"a",
		defaultServerAddress,
		"server address `HOST:PORT`",
	)
	flag.StringVarP(
		&c.ServerURL,
		"ServerURL",
		"b",
		defaultServerURL,
		"server `URL`",
	)

	flag.Parse()

	// env vars
	if err := env.Parse(&c); err != nil {
		fmt.Println("Error while parsing ENV", err)
	}

	return &c
}
