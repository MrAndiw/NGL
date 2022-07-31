package main

import (
	"NGL/router"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Server is starting")

	e := router.New()

	port := os.Getenv("PORT")
	err := e.Start(":" + port)
	if err != nil {
		log.Fatal("aplikasi tidak jalan ", err)
	}
}
