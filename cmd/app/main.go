package main

import (
	"github.com/fromsi/game_2d/internal/interface/ebitengin"
	"log"
)

func main() {
	ebitenginApplication := ebitengin.Application{
		Title:  "Game",
		Width:  500,
		Height: 800,
	}

	if err := ebitenginApplication.Run(); err != nil {
		log.Fatal(err)
	}
}
