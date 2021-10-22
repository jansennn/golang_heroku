package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/jansennn/golang_heroku/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

//SetupDatabaseConnection is creating a new connection to our database
func SetupDatabaseConnection() *gorm.DB {
	// errEnv := godotenv.Load()
	// if errEnv != nil {
	// 	panic("Failed to load env file. Make sure .env file is exists!")
	// }

	//for development
	//dbUser := goDotEnvVariable("DB_USER")
	//dbPass := goDotEnvVariable("DB_PASSWORD")
	//dbHost := goDotEnvVariable("DB_HOST")
	//dbName := goDotEnvVariable("DB_NAME")
	//dbPort := goDotEnvVariable("DB_PORT")

	//for deployment
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	//for development
	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbName, dbPort)

	//for deployment
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	db.AutoMigrate(&entity.User{}, &entity.Product{}, &entity.Description{}, &entity.Project{}, &entity.Career{})
	println("Database connected!")
	return db
}

//CloseDatabaseConnection method is closing a connection between your app and your db
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
