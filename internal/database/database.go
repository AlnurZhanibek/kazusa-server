package database

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
)

func New() *sql.DB {
	dbConfig := mysql.Config{
		User:      os.Getenv("DB_USER"),
		Passwd:    os.Getenv("DB_PASS"),
		Net:       "tcp",
		Addr:      os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		DBName:    os.Getenv("DB_NAME"),
		ParseTime: true,
	}

	db, err := sql.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
