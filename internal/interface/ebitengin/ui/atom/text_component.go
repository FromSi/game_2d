package atom

import (
	"bytes"
	"github.com/fromsi/game_2d/asset/font"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"image/color"
	"log"
)

const (
	DefaultSize  = 20
	DefaultColor = 0xffff
)

var (
	textFaceSource *text.GoTextFaceSource
)

func init() {
	newGoTextFaceSource, err := text.NewGoTextFaceSource(bytes.NewReader(font.Tiny5Regular_ttf))

	if err != nil {
		log.Fatal(err)
	}

	textFaceSource = newGoTextFaceSource
}

type TextComponent struct {
	Text           string
	StartGeometryX float64
	StartGeometryY float64
	Size           float64
	Color          color.Color
}

func (component *TextComponent) FillByDefault() {
	if component.Size == 0 {
		component.Size = DefaultSize
	}

	if component.Color == nil {
		component.Color = color.Gray16{Y: DefaultColor}
	}
}

func (component *TextComponent) Draw(Screen *ebiten.Image) {
	component.FillByDefault()

	drawOptions := &text.DrawOptions{}

	drawOptions.GeoM.Translate(component.StartGeometryX, component.StartGeometryY)
	drawOptions.ColorScale.ScaleWithColor(component.Color)

	textFace := &text.GoTextFace{
		Source: textFaceSource,
		Size:   component.Size,
	}

	text.Draw(Screen, component.Text, textFace, drawOptions)
}
