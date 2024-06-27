package controller

import (
	"github.com/fromsi/game_2d/internal/interface/ebitengin/ui"
	"github.com/fromsi/game_2d/internal/interface/ebitengin/ui/page"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type UpdateController struct{}

type UpdateRequest struct {
	Page ui.Page
}

type UpdateResponse struct {
	Page ui.Page
	Err  error
}

func (controller *UpdateController) Handle(request *UpdateRequest) *UpdateResponse {
	response := &UpdateResponse{
		Page: request.Page,
		Err:  nil,
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		request.Page.HandleClick(getCursorPosition())
	} else if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		response.Page = page.NewBuilderMainPage().GetComponent()
	}

	return response
}

func getCursorPosition() (float64, float64) {
	cursorPositionX, cursorPositionY := ebiten.CursorPosition()

	return float64(cursorPositionX), float64(cursorPositionY)
}
