package config

import (
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string     `yaml:"env"`
	StoragePath string     `yaml:"storagepath"`
	Grpc        GrpcConfig `yaml:"grpc"`
}

type GrpcConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad(configPath string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("Failed to read config: file is not found")
	}
	cfg := new(Config)
	err := cleanenv.ReadConfig(configPath, cfg)
	if err != nil {
		panic("Failed to read config: " + err.Error())
	}
	return cfg
}
