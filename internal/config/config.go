package config

import (
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerCfg ServerConfig
	DbCfg     DatabaseConfig
}

type ServerConfig struct {
	Port   int
	AppUrl string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SSLMode  string
	TimeZone string
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		config, err := load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
		instance = config
	})
	return instance
}

func load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file %v", err)
	}

	serverCfg := ServerConfig{
		Port:   parseInt(getEnv("SERVER_PORT", "8000")),
		AppUrl: getEnv("APP_URL", "http://localhost:8000"),
	}

	dbConfig := DatabaseConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "123456"),
		DbName:   getEnv("DB_NAME", "todolist"),
		SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		TimeZone: getEnv("DB_TIMEZONE", "UTC"),
	}

	return &Config{
		ServerCfg: serverCfg,
		DbCfg:     dbConfig,
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func parseInt(value string) int {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return intValue
}

func parseDuration(value string) time.Duration {
	duration, err := time.ParseDuration(value)
	if err != nil {
		return 1 * time.Hour
	}
	return duration
}
