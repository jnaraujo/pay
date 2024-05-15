package db

import (
	"database/sql"
	"os"

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

func Migrate() error {
	_, err := DB.Exec(`create table if not exists users (
		id text primary key,
		name text not null,
		created_at date null
	);`)
	return err
}

func Close() error {
	return DB.Close()
}
