package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB инициализирует подключение к PostgreSQL
func InitDB() {
	if App == nil {
		log.Fatal("❌ Конфигурация не загружена. Вызовите config.Load() перед InitDB()")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		App.DB.Host,
		App.DB.User,
		App.DB.Password,
		App.DB.DBName,
		App.DB.Port,
		App.DB.SSLMode,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // Silent = минимум логов
	})

	if err != nil {
		log.Fatalf("❌ Не удалось подключиться к PostgreSQL: %v", err)
	}

	// Настройка пула соединений
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	log.Println("✅ Подключение к PostgreSQL успешно!")
}
