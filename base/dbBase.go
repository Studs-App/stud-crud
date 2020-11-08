package base

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var (
	DB *gorm.DB
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load .env file: %s", err.Error())
	}
	dbHost := os.Getenv("DB_HOST")
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbConfig := DBConfig{
		Host:     dbHost,
		Port:     dbPort,
		User:     dbUser,
		Password: dbPassword,
		DBName:   dbName,
	}
	return &dbConfig
}

func DBUrl(dbConfig *DBConfig) string {
	// return fmt.Sprintf(
	// 	"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
	// 	dbConfig.Host,
	// 	dbConfig.User,
	// 	dbConfig.Password,
	// 	dbConfig.DBName,
	// 	dbConfig.Port,
	// )

	dsn := url.URL{
		User:     url.UserPassword(dbConfig.User, dbConfig.Password),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", dbConfig.Host, dbConfig.Port),
		Path:     dbConfig.DBName,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	return dsn.String()
}
