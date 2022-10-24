package database

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	host := os.Getenv("DBHOST")
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASS")
	dbName := os.Getenv("DBNAME")
	ssl := os.Getenv("DBSSL")

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", host, user, password, dbName, ssl)

	gormDB, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed connection to database")
	}

	db, err := gormDB.DB()
	if err != nil {
		return nil, errors.New("failed conection to database")
	}
	db.SetConnMaxIdleTime(100)
	db.SetMaxOpenConns(10)

	return gormDB, nil
}
