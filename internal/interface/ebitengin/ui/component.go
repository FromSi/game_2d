package ui

import "github.com/hajimehoshi/ebiten/v2"

type Component interface {
	Draw(Screen *ebiten.Image)
}
