package main

import (
	"fmt"

	"github.com/jnaraujo/pay/internal/config"
	"github.com/jnaraujo/pay/internal/http"
)

func main() {
	err := config.InitEnv()
	if err != nil {
		panic(err)
	}

	fmt.Println("Hello, world!")

	err = http.NewServer()
	if err != nil {
		panic(err)
	}
}
