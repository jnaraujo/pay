package main

import (
	"github.com/jnaraujo/pay/internal/config"
	"github.com/jnaraujo/pay/internal/db"
	"github.com/jnaraujo/pay/internal/http"
)

func main() {
	err := config.InitEnv()
	if err != nil {
		panic(err)
	}

	err = db.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = http.NewServer()
	if err != nil {
		panic(err)
	}
}
