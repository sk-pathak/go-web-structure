package config

import (
    "github.com/joho/godotenv"
    "log"
    "os"
)

type Config struct {
    Port         string
    DBDriver     string
    DBUser       string
    DBPassword   string
    DBHost       string
    DBPort       string
    DBName       string
    LogLevel     string
}

func LoadConfig() (*Config, error) {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, proceeding with environment variables")
    }

    return &Config{
        Port:       getEnv("PORT", "8080"),
        DBDriver:   getEnv("DB_DRIVER", "postgres"),
        DBUser:     getEnv("DB_USER", "sumit_pathak"),
        DBPassword: getEnv("DB_PASSWORD", "19209449"),
        DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:     getEnv("DB_PORT", "5432"),
        DBName:     getEnv("DB_NAME", "go-database"),
        LogLevel:   getEnv("LOG_LEVEL", "info"),
    }, nil
}

func getEnv(key, defaultValue string) string {
    value, exists := os.LookupEnv(key)
    if !exists {
        return defaultValue
    }
    return value
}
