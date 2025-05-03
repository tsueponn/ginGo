package db

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

func RunMigrationsUp() {
	viper.AutomaticEnv()

	user := viper.GetString("POSTGRES_USER")
	password := viper.GetString("POSTGRES_PASSWORD")
	dbname := viper.GetString("POSTGRES_DB")
	host := viper.GetString("POSTGRES_HOST")
	port := viper.GetString("POSTGRES_PORT")

	if port == "" {
		port = "3000"
	}
	if host == "" {
		host = "localhost"
	}

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbname,
	)

	m, err := migrate.New("file://migrations", dbURL)
	if err != nil {
		log.Fatalf("❌ Failed to initialize migration: %v", err)
	}

	version, dirty, err := m.Version()
	if err == nil && dirty {
		log.Printf("⚠️ Database is dirty at version %d — forcing clean...", version)
		if err := m.Force(int(version)); err != nil {
			log.Fatalf("❌ Failed to force migration: %v", err)
		}
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("❌ Migration up failed: %v", err)
	}

	fmt.Println("✅ Migrations applied successfully.")
}

func RunMigrationsDown() {
	viper.AutomaticEnv()

	user := viper.GetString("POSTGRES_USER")
	password := viper.GetString("POSTGRES_PASSWORD")
	dbname := viper.GetString("POSTGRES_DB")
	host := viper.GetString("POSTGRES_HOST")
	port := viper.GetString("POSTGRES_PORT")

	if port == "" {
		port = "3000"
	}
	if host == "" {
		host = "localhost"
	}

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbname,
	)

	m, err := migrate.New("file://migrations", dbURL)
	if err != nil {
		log.Fatalf("❌ Failed to initialize migration: %v", err)
	}

	version, dirty, err := m.Version()
	if err == nil && dirty {
		log.Printf("⚠️ Database is dirty at version %d — forcing clean...", version)
		if err := m.Force(int(version)); err != nil {
			log.Fatalf("❌ Failed to force migration: %v", err)
		}
	}

	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("❌ Migration down failed: %v", err)
	}

	fmt.Println("✅ Migrations rolled back successfully.")
}
