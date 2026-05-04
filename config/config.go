package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Port       uint   `json:"port"`
	GrpcPort   uint   `json:"grpc_port"`
	DbUrl      string `json:"db_url"`
	Production bool   `json:"production"`
}

const cfgFileName = "config.json"

func LoadConfig() *Config {
	cfg := &Config{
		Port:     4000,
		GrpcPort: 9000,
		DbUrl:    "",
	}
	_, err := os.Stat(cfgFileName)
	if err != nil {
		log.Printf("No config file found => %v\n", err)
	}
	file, err := os.Open(cfgFileName)
	if err != nil {
		log.Printf("Error while reading config file=> %v\n", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(cfg); err != nil {
		log.Printf("Error while decoding config file=> %v\n", err)
	}
	return cfg
}
