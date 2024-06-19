package ebitengin

import (
	"github.com/fromsi/game_2d/internal/interface/ebitengin/controller"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	drawController   = controller.DrawController{}
	updateController = controller.UpdateController{}
)

type Application struct {
	Title  string
	Width  int
	Height int
}

func (application *Application) Draw(screen *ebiten.Image) {
	drawRequest := controller.DrawRequest{
		Screen: screen,
	}

	drawController.Handle(&drawRequest)
}

func (application *Application) Update() error {
	updateRequest := controller.UpdateRequest{}

	return updateController.Handle(&updateRequest).Err
}

func (application *Application) Layout(_ int, _ int) (int, int) {
	return application.Width, application.Height
}

func (application *Application) Run() error {
	ebiten.SetWindowSize(application.Width, application.Height)

	ebiten.SetWindowTitle(application.Title)

	return ebiten.RunGame(application)
}
