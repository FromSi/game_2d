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
	Size           float64
	Color          color.Color
	OnClick        func()
	startGeometryX float64
	startGeometryY float64
}

func (component *TextComponent) GetStartGeometryX() float64 {
	return component.startGeometryX
}

func (component *TextComponent) GetStartGeometryY() float64 {
	return component.startGeometryY
}

func (component *TextComponent) SetStartGeometryX(startGeometryX float64) {
	component.startGeometryX = startGeometryX
}

func (component *TextComponent) SetStartGeometryY(startGeometryY float64) {
	component.startGeometryY = startGeometryY
}

func (component *TextComponent) GetEndGeometryX() float64 {
	textFace := &text.GoTextFace{
		Source: component.getFontFaceSource(),
		Size:   component.Size,
	}

	endGeometryX, _ := text.Measure(component.Text, textFace, 0)

	return endGeometryX + component.startGeometryX
}

func (component *TextComponent) GetEndGeometryY() float64 {
	textFace := &text.GoTextFace{
		Source: component.getFontFaceSource(),
		Size:   component.Size,
	}

	_, endGeometryY := text.Measure(component.Text, textFace, 0)

	return endGeometryY + component.startGeometryY
}

func (component *TextComponent) OnDraw(Screen *ebiten.Image) {
	drawOptions := &text.DrawOptions{}

	drawOptions.GeoM.Translate(component.startGeometryX, component.startGeometryY)
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
	isWithinX := (component.startGeometryX <= geometryX) && (component.GetEndGeometryX() >= geometryX)
	isWithinY := (component.startGeometryY <= geometryY) && (component.GetEndGeometryY() >= geometryY)

	return isWithinX && isWithinY
}

func (component *TextComponent) getFontFaceSource() *text.GoTextFaceSource {
	switch component.Type {
	default:
		return tiny5RegularFontFaceSource
	}
}

type BuilderTextComponent struct {
	component *TextComponent
}

func NewBuilderTextComponent() *BuilderTextComponent {
	return &BuilderTextComponent{
		component: &TextComponent{
			Text:           "",
			Type:           Tiny5RegularFontType,
			startGeometryX: 0,
			startGeometryY: 0,
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
	builder.component.startGeometryX = startGeometryX

	return builder
}

func (builder *BuilderTextComponent) SetStartGeometryY(startGeometryY float64) *BuilderTextComponent {
	builder.component.startGeometryY = startGeometryY

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
