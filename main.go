package main

import (
	"NGL/config"
	"NGL/router"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Server is starting")

	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("config tidak di temukan", err)
	}

	e := router.New()

	e.Start(config.App.Host)
}
