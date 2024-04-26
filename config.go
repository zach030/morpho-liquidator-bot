package main

type Config struct {
	HttpEndpoint string   `yaml:"http_endpoint"`
	WsEndpoint   string   `yaml:"ws_endpoint"`
	Markets      []string `yaml:"markets"`
	PrivateKey   string   `yaml:"private_key"`
}
