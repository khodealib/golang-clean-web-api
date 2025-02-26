package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Server struct {
	Port    string `yaml:"port"`
	Host    string `yaml:"host"`
	RunMode string `yaml:"runMode"`
}

type Logger struct {
	FilePath string `yaml:"filePath"`
	Level    string `yaml:"level"`
	Encoding string `yaml:"encoding"`
}

type Cors struct {
	AllowOrigins string `yaml:"allowOrigins"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type Redis struct {
	Host               string `yaml:"host"`
	Port               string `yaml:"port"`
	Password           string `yaml:"password"`
	DB                 int    `yaml:"db"`
	MinIdleConnections int    `yaml:"minIdleConnections"`
	PoolSize           int    `yaml:"poolSize"`
	PoolTimeout        int    `yaml:"poolTimeout"`
}

type Config struct {
	Server   Server   `yaml:"server"`
	Logger   Logger   `yaml:"logger"`
	Cors     Cors     `yaml:"cors"`
	Postgres Postgres `yaml:"postgres"`
	Redis    Redis    `yaml:"redis"`
}

func getConfigPath(env string) string {
	if env == "docker" {
		return "../config/config-docker.yml"
	} else if env == "production" {
		return "../config/config-production.yml"
	}
	return "../config/config-development.yml"
}

func LoadConfig(fileName string, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(fileName)
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, fmt.Errorf("config file not found: %w", err)
		}
		return nil, err
	}

	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		fmt.Printf("Error unmarshalling config: %v", err)
		return nil, err
	}
	return &config, nil
}

func GetConfig() *Config {
	env := os.Getenv("ENV")
	v, err := LoadConfig(getConfigPath(env), "yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	config, err := ParseConfig(v)
	if err != nil {
		log.Fatalf("Error parsing config: %v", err)
	}
	return config
}
