package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App AppConfig
	DB  DBConfig
}

type AppConfig struct {
	Env  string
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

var App *Config

// Load загружает конфигурацию из .env файла
func Load() error {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env файл не найден, используем системные переменные")
	}

	App = &Config{
		App: AppConfig{
			Env:  getEnv("APP_ENV", "development"),
			Port: getEnv("APP_PORT", "8080"),
		},
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			DBName:   getEnv("DB_NAME", "myapp_db"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
	}

	// Проверка обязательных полей
	if App.DB.Password == "" {
		return fmt.Errorf("DB_PASSWORD обязателен")
	}

	log.Printf("✅ Конфигурация загружена (Env: %s, DB: %s)", App.App.Env, App.DB.DBName)
	return nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
