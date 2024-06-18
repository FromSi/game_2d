package ebitengin

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Application struct {
	Title  string
	Width  int
	Height int
}

func (g *Application) Layout(_ int, _ int) (int, int) {
	return g.Width, g.Height
}

func (app *Application) Run() error {
	ebiten.SetWindowSize(app.Width, app.Height)

	ebiten.SetWindowTitle(app.Title)

	return ebiten.RunGame(app)
}
