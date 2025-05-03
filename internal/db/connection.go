package db

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	// Viper reads environment variables automatically
	viper.AutomaticEnv()

	user := viper.GetString("POSTGRES_USER")
	password := viper.GetString("POSTGRES_PASSWORD")
	dbname := viper.GetString("POSTGRES_DB")
	dbhost := viper.GetString("POSTGRES_HOST")
	dbport := viper.GetString("POSTGRES_PORT")

	if dbport == "" {
		dbport = "3000" // fallback port if not set
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbhost, user, password, dbname, dbport)

	println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to DB: %w", err)
	}

	return db, nil
}
