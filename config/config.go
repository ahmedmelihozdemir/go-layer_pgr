package config

import (
	"fmt"
	"os"

	"context"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	APIVersion string
	Prefix     string
	Port       string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DSN        string
}

var DB *gorm.DB
var RDB *redis.Client
var Ctx = context.Background()

func LoadConfig() *Config {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
	}

	RDB = redis.NewClient(&redis.Options{
		Addr: "localhost:" +
			os.Getenv("REDIS_PORT"),
	})

	return &Config{
		APIVersion: os.Getenv("API_VERSION"),
		Prefix:     "/api/" + os.Getenv("API_VERSION"),
		Port:       os.Getenv("PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DSN:        dsn,
	}
}

func GetSecretKey() string {
	return os.Getenv("SECRET_KEY")
}
