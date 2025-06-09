package config

import (
	"church_pesca_teams/domain"
	"fmt"
	"log"
	"os" // Used for environment variables

	"gorm.io/driver/postgres" // PostgreSQL driver for GORM
	"gorm.io/gorm"            // GORM ORM library
)

func InitDb() *gorm.DB {

	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	DB_SSLMODE := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s  sslmode=%s ",
		DB_HOST,
		DB_USER,
		DB_PASSWORD,
		DB_NAME,
		DB_PORT,
		DB_SSLMODE,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	db.AutoMigrate(&domain.PescaTeam{})

	return db
}
