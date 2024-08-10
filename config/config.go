package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Environment string `yaml:"env" end-default:"develompent"`
	Database    `yaml:"database"`
	HttpServer  `yaml:"http_server"`
}

type HttpServer struct {
	Port        string        `yaml:"port" env-default:"80"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeOut time.Duration `yaml:"idle_timeout" env-default:"4s"`
	Host        string        `yaml:"host" env-default:"localhost"`
}

type Database struct {
	Port             string `yaml:"port" env:"PORT" env-default:"5432"`
	Host             string `yaml:"host" env:"HOST" env-default:"localhost"`
	Name             string `yaml:"name" env:"NAME" env-default:"postgres"`
	User             string `yaml:"user" env:"USER" env-default:"admin"`
	Password         string `yaml:"password" env:"admin"`
	ConnectionString string `yaml:"connection_string" env:"CONNECTION_STRING"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatal("CONFIG_PATH is empty or not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file in path %s does not exists", configPath)
	}

	var config Config

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("cannot read config %s", err)
	}

	return &config
}
