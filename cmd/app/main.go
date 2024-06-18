package main

import (
	"github.com/fromsi/game_2d/internal"
	"log"
)

func main() {
	app := internal.Application{
		Title:  "Game",
		Width:  500,
		Height: 800,
	}

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
