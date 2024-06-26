package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

func New() *sql.DB {
	dbConfig := mysql.Config{
		User:      os.Getenv("DB_USER"),
		Passwd:    os.Getenv("DB_PASS"),
		Net:       "tcp",
		Addr:      "127.0.0.1:3306",
		DBName:    os.Getenv("DB_NAME"),
		ParseTime: true,
	}

	db, err := sql.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		log.Fatalf("database connection error: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("database ping error: %v", err)
	}

	return db
}
