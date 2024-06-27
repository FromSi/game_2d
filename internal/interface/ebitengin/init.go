package ebitengin

import (
	"github.com/fromsi/game_2d/internal/interface/ebitengin/controller"
	"github.com/fromsi/game_2d/internal/interface/ebitengin/ui"
	uipage "github.com/fromsi/game_2d/internal/interface/ebitengin/ui/page"
	"github.com/hajimehoshi/ebiten/v2"
)

type Application struct {
	Page             ui.Page
	Title            string
	Width            int
	Height           int
	drawController   controller.DrawController
	updateController controller.UpdateController
}

func (application *Application) Draw(screen *ebiten.Image) {
	drawRequest := controller.DrawRequest{
		Page:   application.Page,
		Screen: screen,
	}

	application.drawController.Handle(&drawRequest)
}

func (application *Application) Update() error {
	updateRequest := controller.UpdateRequest{
		Page: application.Page,
	}

	updateResponse := application.updateController.Handle(&updateRequest)

	application.Page = updateResponse.Page

	return updateResponse.Err
}

func (application *Application) Layout(_ int, _ int) (int, int) {
	return application.Width, application.Height
}

func (application *Application) Run() error {
	ebiten.SetWindowSize(application.Width, application.Height)

	ebiten.SetWindowTitle(application.Title)

	application.Page = uipage.NewBuilderMainPage().GetComponent()

	application.drawController = controller.DrawController{}
	application.updateController = controller.UpdateController{}

	return ebiten.RunGame(application)
}
