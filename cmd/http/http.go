package main

import (
	"log"
	"os"

	"github.com/Camil0Guerrero/http/internal/app"
)

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal("Error running the application: ", err)
	}
}
