package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jnaraujo/pay/internal/db"
)

func main() {
	err := db.InitDB()
	if err != nil {
		panic(err)
	}

	m, err := db.Migrate()
	if err != nil {
		panic(err)
	}

	cmd := os.Args[len(os.Args)-1]

	if cmd == "up" {
		err = m.Up()
	} else if cmd == "down" {
		err = m.Down()
	} else {
		fmt.Println("Invalid command")
		return
	}
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(err)
	}
}
