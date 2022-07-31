package main

import (
	"NGL/router"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Server is starting")

	// config, err := config.LoadConfig(".")
	// if err != nil {
	// 	log.Fatal("config tidak di temukan", err)
	// }

	e := router.Start()

	port := os.Getenv("PORT")
	err := e.Start(":" + port)
	if err != nil {
		log.Fatal("aplikasi tidak jalan ", err)
	}
}
