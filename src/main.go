package main

import (
	"agoravote-app-backend/src/routes"
	"log"
)

func main() {
	router := routes.SetupRoutes()
	log.Println("Starting server on :8081")
	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
