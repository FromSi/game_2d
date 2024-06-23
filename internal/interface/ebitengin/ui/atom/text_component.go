package atom

import (
	"bytes"
	assetfont "github.com/fromsi/game_2d/asset/font"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"image/color"
	"log"
)

type FontType int

const (
	Tiny5RegularFontType FontType = iota
)

var (
	tiny5RegularFontFaceSource *text.GoTextFaceSource
)

func init() {
	newTiny5RegularFontFaceSource, err := text.NewGoTextFaceSource(bytes.NewReader(assetfont.Tiny5Regular_ttf))

	if err != nil {
		log.Fatal(err)
	}

	tiny5RegularFontFaceSource = newTiny5RegularFontFaceSource
}

type TextComponent struct {
	Text           string
	Type           FontType
	StartGeometryX float64
	StartGeometryY float64
	Size           float64
	Color          color.Color
	OnClick        func()
}

func (component *TextComponent) getFontFaceSource() *text.GoTextFaceSource {
	switch component.Type {
	default:
		return tiny5RegularFontFaceSource
	}
}

func (component *TextComponent) GetEndGeometries() (float64, float64) {
	textFace := &text.GoTextFace{
		Source: component.getFontFaceSource(),
		Size:   component.Size,
	}

	endGeometryX, endGeometryY := text.Measure(component.Text, textFace, 0)

	endGeometryX += component.StartGeometryX
	endGeometryY += component.StartGeometryY

	return endGeometryX, endGeometryY
}

func (component *TextComponent) OnDraw(Screen *ebiten.Image) {
	drawOptions := &text.DrawOptions{}

	drawOptions.GeoM.Translate(component.StartGeometryX, component.StartGeometryY)
	drawOptions.ColorScale.ScaleWithColor(component.Color)

	textFace := &text.GoTextFace{
		Source: component.getFontFaceSource(),
		Size:   component.Size,
	}

	text.Draw(Screen, component.Text, textFace, drawOptions)
}

func (component *TextComponent) HandleClick(geometryX, geometryY float64) bool {
	if !component.IsWithin(geometryX, geometryY) {
		return false
	}

	if component.OnClick != nil {
		component.OnClick()

		return true
	}

	return false
}

func (component *TextComponent) IsWithin(geometryX, geometryY float64) bool {
	endGeometryX, endGeometryY := component.GetEndGeometries()

	isWithinX := (component.StartGeometryX <= geometryX) && (endGeometryX >= geometryX)
	isWithinY := (component.StartGeometryY <= geometryY) && (endGeometryY >= geometryY)

	return isWithinX && isWithinY
}

type BuilderTextComponent struct {
	component *TextComponent
}

func NewBuilderTextComponent() *BuilderTextComponent {
	return &BuilderTextComponent{
		component: &TextComponent{
			Text:           "",
			Type:           Tiny5RegularFontType,
			StartGeometryX: 0,
			StartGeometryY: 0,
			Size:           20,
			Color:          color.White,
		},
	}
}

func (builder *BuilderTextComponent) GetComponent() *TextComponent {
	return builder.component
}

func (builder *BuilderTextComponent) SetText(text string) *BuilderTextComponent {
	builder.component.Text = text

	return builder
}

func (builder *BuilderTextComponent) SetType(_type FontType) *BuilderTextComponent {
	builder.component.Type = _type

	return builder
}

func (builder *BuilderTextComponent) SetStartGeometryX(startGeometryX float64) *BuilderTextComponent {
	builder.component.StartGeometryX = startGeometryX

	return builder
}

func (builder *BuilderTextComponent) SetStartGeometryY(startGeometryY float64) *BuilderTextComponent {
	builder.component.StartGeometryY = startGeometryY

	return builder
}

func (builder *BuilderTextComponent) SetSize(size float64) *BuilderTextComponent {
	builder.component.Size = size

	return builder
}

func (builder *BuilderTextComponent) SetColor(color color.Color) *BuilderTextComponent {
	builder.component.Color = color

	return builder
}

func (builder *BuilderTextComponent) SetOnClick(onClick func()) *BuilderTextComponent {
	builder.component.OnClick = onClick

	return builder
}
