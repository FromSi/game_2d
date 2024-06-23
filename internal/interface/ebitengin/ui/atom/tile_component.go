package atom

import (
	"bytes"
	assetimage "github.com/fromsi/game_2d/asset/image"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"log"
)

type TileType int

const (
	ResourceTileType TileType = iota
	ToolTileType
)

var (
	resourceTileImage *ebiten.Image
	toolTileImage     *ebiten.Image
)

func init() {
	newResourceTileImage, _, err := image.Decode(bytes.NewReader(assetimage.ResourceTiles_png))

	if err != nil {
		log.Fatal(err)
	}

	newToolTileImage, _, err := image.Decode(bytes.NewReader(assetimage.ToolTiles_png))

	if err != nil {
		log.Fatal(err)
	}

	resourceTileImage = ebiten.NewImageFromImage(newResourceTileImage)
	toolTileImage = ebiten.NewImageFromImage(newToolTileImage)
}

type TileComponent struct {
	Id             int
	Type           TileType
	Scale          float64
	OnClick        func()
	startGeometryX float64
	startGeometryY float64
}

func (component *TileComponent) GetStartGeometryX() float64 {
	return component.startGeometryX
}

func (component *TileComponent) GetStartGeometryY() float64 {
	return component.startGeometryY
}

func (component *TileComponent) SetStartGeometryX(startGeometryX float64) {
	component.startGeometryX = startGeometryX
}

func (component *TileComponent) SetStartGeometryY(startGeometryY float64) {
	component.startGeometryY = startGeometryY
}

func (component *TileComponent) GetEndGeometryX() float64 {
	result := float64(component.getTileSize()) * component.Scale

	return result + component.startGeometryX
}

func (component *TileComponent) GetEndGeometryY() float64 {
	result := float64(component.getTileSize()) * component.Scale

	return result + component.startGeometryY
}

func (component *TileComponent) OnDraw(Screen *ebiten.Image) {
	drawOptions := &ebiten.DrawImageOptions{}

	drawOptions.GeoM.Scale(component.Scale, component.Scale)
	drawOptions.GeoM.Translate(component.startGeometryX, component.startGeometryY)

	startGeometryTileX, startGeometryTileY := component.getStartGeometryTile()

	tileSize := component.getTileSize()

	imageRect := image.Rect(
		startGeometryTileX,
		startGeometryTileY,
		startGeometryTileX+tileSize,
		startGeometryTileY+tileSize,
	)

	Screen.DrawImage(component.getEbitenImage().SubImage(imageRect).(*ebiten.Image), drawOptions)
}

func (component *TileComponent) HandleClick(geometryX, geometryY float64) bool {
	if !component.IsWithin(geometryX, geometryY) {
		return false
	}

	if component.OnClick != nil {
		component.OnClick()

		return true
	}

	return false
}

func (component *TileComponent) IsWithin(geometryX, geometryY float64) bool {
	isWithinX := (component.startGeometryX <= geometryX) && (component.GetEndGeometryX() >= geometryX)
	isWithinY := (component.startGeometryY <= geometryY) && (component.GetEndGeometryY() >= geometryY)

	return isWithinX && isWithinY
}

func (component *TileComponent) getStartGeometryTile() (int, int) {
	tileWidth := component.getEbitenImage().Bounds().Dx()
	tileSize := component.getTileSize()
	tileXCount := tileWidth / tileSize

	x := (component.Id % tileXCount) * tileSize
	y := (component.Id / tileXCount) * tileSize

	return x, y
}

func (component *TileComponent) getTileSize() int {
	switch component.Type {
	case ToolTileType:
		return assetimage.ToolTileSize
	default:
		return assetimage.ResourceTileSize
	}
}

func (component *TileComponent) getEbitenImage() *ebiten.Image {
	switch component.Type {
	case ToolTileType:
		return toolTileImage
	default:
		return resourceTileImage
	}
}

type BuilderTileComponent struct {
	component *TileComponent
}

func NewBuilderTileComponent() *BuilderTileComponent {
	return &BuilderTileComponent{
		component: &TileComponent{
			Id:             0,
			Type:           ResourceTileType,
			startGeometryX: 0,
			startGeometryY: 0,
			Scale:          1,
		},
	}
}

func (builder *BuilderTileComponent) GetComponent() *TileComponent {
	return builder.component
}

func (builder *BuilderTileComponent) SetId(id int) *BuilderTileComponent {
	builder.component.Id = id

	return builder
}

func (builder *BuilderTileComponent) SetType(_type TileType) *BuilderTileComponent {
	builder.component.Type = _type

	return builder
}

func (builder *BuilderTileComponent) SetStartGeometryX(startGeometryX float64) *BuilderTileComponent {
	builder.component.startGeometryX = startGeometryX

	return builder
}

func (builder *BuilderTileComponent) SetStartGeometryY(startGeometryY float64) *BuilderTileComponent {
	builder.component.startGeometryY = startGeometryY

	return builder
}

func (builder *BuilderTileComponent) SetScale(scale float64) *BuilderTileComponent {
	builder.component.Scale = scale

	return builder
}

func (builder *BuilderTileComponent) SetOnClick(onClick func()) *BuilderTileComponent {
	builder.component.OnClick = onClick

	return builder
}
