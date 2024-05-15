package db

import (
	"database/sql"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() error {
	pwd, _ := os.Getwd()
	dirPath := pwd + "/data"

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.Mkdir(dirPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	db, err := sql.Open("sqlite3", dirPath+"/pay.db")
	if err != nil {
		return err
	}
	DB = db
	return nil
}

func Migrate() (*migrate.Migrate, error) {
	driver, err := sqlite3.WithInstance(DB, &sqlite3.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "sqlite3", driver)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func Close() error {
	return DB.Close()
}
