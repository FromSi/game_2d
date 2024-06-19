package controller

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type DrawController struct {
}

type DrawRequest struct {
	Screen *ebiten.Image
}

type DrawResponse struct {
	Err error
}

func (controller *DrawController) Handle(request *DrawRequest) *DrawResponse {
	return &DrawResponse{Err: nil}
}
