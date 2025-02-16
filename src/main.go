package main

import (
    "agoravote-app-backend/src/routes"
    "agoravote-app-backend/src/database"
    "log"
    "net/http"
    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    database.ConnectDB()
    router := routes.SetupRoutes()
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatal(err)
    }
}