package controller

import (
	"github.com/fromsi/game_2d/internal/interface/ebitengin/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type UpdateController struct {
	Page ui.Component
}

type UpdateRequest struct {
}

type UpdateResponse struct {
	Err error
}

func (controller *UpdateController) Handle(request *UpdateRequest) *UpdateResponse {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		controller.Page.HandleClick(getCursorPosition())
	}

	return &UpdateResponse{Err: nil}
}

func getCursorPosition() (float64, float64) {
	cursorPositionX, cursorPositionY := ebiten.CursorPosition()

	return float64(cursorPositionX), float64(cursorPositionY)
}
