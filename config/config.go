package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	HttpEndpoint  string   `yaml:"http_endpoint"`
	WsEndpoint    string   `yaml:"ws_endpoint"`
	BotAddress    string   `yaml:"bot_address"`
	PrivateKey    string   `yaml:"private_key"`
	OneInchApiKey string   `yaml:"one_inch_api_key"`
	Markets       []string `yaml:"markets"`
}

func LoanConfig(path string) *Config {
	buf, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var cfg Config
	if err = yaml.Unmarshal(buf, &cfg); err != nil {
		panic(err)
	}
	return &cfg
}
