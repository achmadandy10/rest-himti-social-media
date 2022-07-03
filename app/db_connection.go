package app

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	AppTimeZone := GetEnv("APP_TIMEZONE", "Asia/Shanghai")

	DBHost := GetEnv("DB_HOST", "localhost")
	DBPort := GetEnv("DB_PORT", "3306")
	DBName := GetEnv("DB_DATABASE", "golang")
	DBUsername := GetEnv("DB_USERNAME", "root")
	DBPassword := GetEnv("DB_PASSWORD", "")
	DBDriver := GetEnv("DB_DRIVER", "mysql")

	if DBDriver == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUsername, DBPassword, DBHost, DBPort, DBHost)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("Failed connection to database name: " + DBName)
		}

		return db
	} else {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", DBHost, DBUsername, DBPassword, DBName, DBPort, AppTimeZone)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("Failed connection to database name: " + DBName)
		}

		return db
	}
}

func DisconectDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}

	dbSQL.Close()
}
