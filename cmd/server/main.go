package main

import (
	"log"
	"rynds-api/internal/app"
)

func main() {
	app := app.New()
	log.Fatal(app.Listen(":3002"))
}
