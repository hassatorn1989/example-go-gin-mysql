package main

import (
	"go-gin-crud/config"
	"go-gin-crud/routes"
	"log"
)

func main() {
	config.InitDB()
	r := routes.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
