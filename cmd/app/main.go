package main

import (
	"github.com/fromsi/game_2d/internal/interfaces/ebitengin"
	"log"
)

func main() {
	app := ebitengin.Application{
		Title:  "Game",
		Width:  500,
		Height: 800,
	}

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
