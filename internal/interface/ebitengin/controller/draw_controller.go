package controller

import (
	"github.com/fromsi/game_2d/internal/interface/ebitengin/ui"
	"github.com/hajimehoshi/ebiten/v2"
)

type DrawController struct{}

type DrawRequest struct {
	Page   ui.Page
	Screen *ebiten.Image
}

type DrawResponse struct {
	Err error
}

func (controller *DrawController) Handle(request *DrawRequest) *DrawResponse {
	request.Page.OnDraw(request.Screen)

	return &DrawResponse{Err: nil}
}
