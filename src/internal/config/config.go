package config

import (
	"os"
	"strconv"

	"go.yaml.in/yaml/v3"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	DB     DBConfig     `yaml:"database"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

func defaultConfig() Config {
	return Config{
		Server: ServerConfig{
			Port: 8080,
		},
	}
}

func loadFile(path string, cfg *Config) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, cfg)
}

func getEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}

func getEnvInt(key string, fallback int) int {
	if value, exists := os.LookupEnv(key); exists {
		n, err := strconv.Atoi(value)
		if err == nil {
			return n
		}
	}
	return fallback
}

func overrideWithEnv(cfg *Config) {
	cfg.Server.Port = getEnvInt(
		"APP_PORT",
		cfg.Server.Port,
	)
}

func LoadConfig(path string) (*Config, error) {
	//1. Load default configs
	cfg := defaultConfig()

	//2. Load From yaml
	err := loadFile(path, &cfg)
	if err != nil {
		return nil, err
	}

	//3. override with environment variables
	overrideWithEnv(&cfg)

	return &cfg, nil
}
