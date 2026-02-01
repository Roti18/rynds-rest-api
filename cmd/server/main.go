package main

import (
	"log"
	"os"
	"rynds-api/internal/app"
)

func main() {
	app := app.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3002"
	}
	log.Fatal(app.Listen(":" + port))
}
