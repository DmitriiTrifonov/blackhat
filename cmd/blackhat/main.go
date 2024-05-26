package main

import (
	"github.com/DmitriiTrifonov/blackhat/internal/app"
	"log"
)

func main() {
	a := app.New(app.DefaultConfig)

	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}
