package main

import (
	"github.com/jnaraujo/pay/internal/db"
)

func main() {
	err := db.InitDB()
	if err != nil {
		panic(err)
	}

	err = db.Migrate()
	if err != nil {
		panic(err)
	}
}
