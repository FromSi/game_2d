package ebitengin

import (
	"github.com/fromsi/game_2d/internal/interface/ebitengin/controller"
	uipage "github.com/fromsi/game_2d/internal/interface/ebitengin/ui/page"
	"github.com/hajimehoshi/ebiten/v2"
)

type Application struct {
	Title            string
	Width            int
	Height           int
	drawController   controller.DrawController
	updateController controller.UpdateController
}

func (application *Application) Draw(screen *ebiten.Image) {
	drawRequest := controller.DrawRequest{
		Screen: screen,
	}

	application.drawController.Handle(&drawRequest)
}

func (application *Application) Update() error {
	updateRequest := controller.UpdateRequest{}

	return application.updateController.Handle(&updateRequest).Err
}

func (application *Application) Layout(_ int, _ int) (int, int) {
	return application.Width, application.Height
}

func (application *Application) Run() error {
	ebiten.SetWindowSize(application.Width, application.Height)

	ebiten.SetWindowTitle(application.Title)

	mainPage := uipage.NewBuilderMainPage().GetComponent()

	application.drawController = controller.DrawController{Page: mainPage}
	application.updateController = controller.UpdateController{Page: mainPage}

	return ebiten.RunGame(application)
}
