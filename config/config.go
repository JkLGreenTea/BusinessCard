package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"strings"
)

const (
	PathConfigFile = "/config/config.toml"
)

type Config struct {
	IP		string	`toml:"ip"`
	Port	string	`toml:"port"`
}

func NewConfig() *Config {
	cfg := &Config{}

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	_, err = toml.DecodeFile(strings.Replace(path, "\\", "/", -1) + PathConfigFile, cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}