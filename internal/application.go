package internal

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Application struct {
	Title  string
	Width  int
	Height int
}

type Game struct {
	app Application
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		fmt.Println("+1")
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(_ int, _ int) (int, int) {
	return g.app.Width, g.app.Height
}

func (app *Application) Run() error {
	ebiten.SetWindowSize(app.Width, app.Height)

	ebiten.SetWindowTitle(app.Title)

	return ebiten.RunGame(&Game{app: *app})
}
