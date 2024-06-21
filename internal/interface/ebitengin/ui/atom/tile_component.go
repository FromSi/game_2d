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
	ID             int
	Type           TileType
	StartGeometryX float64
	StartGeometryY float64
	Scale          float64
}

func (component *TileComponent) getStartGeometryTile() (int, int) {
	tileWidth := component.getEbitenImage().Bounds().Dx()
	tileSize := component.getTileSize()
	tileXCount := tileWidth / tileSize

	x := (component.ID % tileXCount) * tileSize
	y := (component.ID / tileXCount) * tileSize

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

func (component *TileComponent) fillByDefault() {
	if component.Scale == 0 {
		component.Scale = 5
	}
}

func (component *TileComponent) Draw(Screen *ebiten.Image) {
	component.fillByDefault()

	drawOptions := &ebiten.DrawImageOptions{}

	drawOptions.GeoM.Scale(component.Scale, component.Scale)
	drawOptions.GeoM.Translate(component.StartGeometryX, component.StartGeometryY)

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
