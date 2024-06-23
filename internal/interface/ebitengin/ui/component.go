package ui

import "github.com/hajimehoshi/ebiten/v2"

type Page interface {
	OnDraw(*ebiten.Image)
	HandleClick(float64, float64) bool
}

type Component interface {
	OnDraw(*ebiten.Image)
	HandleClick(float64, float64) bool
	IsWithin(float64, float64) bool
	GetEndGeometryX() float64
	GetEndGeometryY() float64
	GetStartGeometryX() float64
	GetStartGeometryY() float64
	SetStartGeometryX(float64)
	SetStartGeometryY(float64)
}
