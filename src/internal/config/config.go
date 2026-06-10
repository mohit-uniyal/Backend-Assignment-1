package config

import (
	"os"
	"strconv"

	"go.yaml.in/yaml/v3"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	DB     DBConfig     `yaml:"database"`
	Redis  RedisConfig  `yaml:"redis"`
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

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

func defaultConfig() Config {
	return Config{
		Server: ServerConfig{
			Port: 8080,
		},
		DB: DBConfig{
			Host:     "localhost",
			Port:     5432,
			User:     "postgres",
			Password: "postgres",
			Name:     "postgres",
		},
		Redis: RedisConfig{
			Host:     "localhost",
			Port:     6379,
			Password: "",
			Db:       0,
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
	cfg.DB.Host = getEnv(
		"DB_HOST",
		cfg.DB.Host,
	)
	cfg.DB.Port = getEnvInt(
		"DB_PORT",
		cfg.DB.Port,
	)
	cfg.DB.Name = getEnv(
		"DB_NAME",
		cfg.DB.Name,
	)
	cfg.DB.User = getEnv(
		"DB_USER",
		cfg.DB.User,
	)
	cfg.DB.Password = getEnv(
		"DB_PASSWORD",
		cfg.DB.Password,
	)
	cfg.Redis.Host = getEnv(
		"REDIS_HOST",
		cfg.Redis.Host,
	)
	cfg.Redis.Port = getEnvInt(
		"REDIS_PORT",
		cfg.Redis.Port,
	)
	cfg.Redis.Password = getEnv(
		"REDIS_PASSWORD",
		cfg.Redis.Password,
	)
	cfg.Redis.Db = getEnvInt(
		"REDIS_DB",
		cfg.Redis.Db,
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
