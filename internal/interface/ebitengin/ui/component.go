package ui

import "github.com/hajimehoshi/ebiten/v2"

type Component interface {
	OnDraw(Screen *ebiten.Image)
	HandleClick(geometryX, geometryY float64) bool
	IsWithin(geometryX, geometryY float64) bool
	GetEndGeometries() (float64, float64)
}
