// Package config owns the shared MySQL connection used by the models layer.
package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// db is the package-level GORM handle. Models call GetDB() after Connect().
var db *gorm.DB

// Connect opens a MySQL connection using the DB_DSN environment variable.
// It fatals on missing DSN or connection failure because the API cannot serve without a database.
func Connect() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN is not set")
	}

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	db = conn
	fmt.Println("connected to MySQL")
}

// GetDB returns the shared GORM connection (must call Connect first).
func GetDB() *gorm.DB {
	return db
}
