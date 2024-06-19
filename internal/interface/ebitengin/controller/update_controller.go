package controller

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type UpdateController struct {
}

type UpdateRequest struct {
}

type UpdateResponse struct {
	Err error
}

func (controller *UpdateController) Handle(request *UpdateRequest) *UpdateResponse {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		fmt.Println("+1")
	}

	return &UpdateResponse{Err: nil}
}
