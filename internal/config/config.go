package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

func LoadDatabaseConfig() DatabaseConfig{
	_ = godotenv.Load()

	return DatabaseConfig{
		User:     getEnv("DB_USER", "crypin"),
		Password: getEnv("DB_PASSWORD", "crypin123"),
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "3306"),
		Name:     getEnv("DB_NAME", "CrypinDb"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func (c DatabaseConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		c.User, c.Password, c.Host, c.Port, c.Name)
}