package page

import (
	"github.com/fromsi/game_2d/internal/interface/ebitengin/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type MainPage struct {
	Component      ui.Component
	StartGeometryX float64
	StartGeometryY float64
}

func (page *MainPage) GetEndGeometries() (float64, float64) {
	return math.MaxFloat64, math.MaxFloat64
}

func (page *MainPage) OnDraw(Screen *ebiten.Image) {
	page.Component.OnDraw(Screen)
}

func (page *MainPage) HandleClick(geometryX, geometryY float64) bool {
	if !page.IsWithin(geometryX, geometryY) {
		return false
	}

	return page.Component.HandleClick(geometryX, geometryY)
}

func (page *MainPage) IsWithin(geometryX, geometryY float64) bool {
	endGeometryX, endGeometryY := page.GetEndGeometries()

	isWithinX := (page.StartGeometryX <= geometryX) && (endGeometryX >= geometryX)
	isWithinY := (page.StartGeometryY <= geometryY) && (endGeometryY >= geometryY)

	return isWithinX && isWithinY
}

type BuilderMainPage struct {
	component *MainPage
}

func NewBuilderMainPage() *BuilderMainPage {
	return &BuilderMainPage{
		component: &MainPage{
			StartGeometryX: 0,
			StartGeometryY: 0,
		},
	}
}

func (builder *BuilderMainPage) GetComponent() *MainPage {
	return builder.component
}

func (builder *BuilderMainPage) SetComponent(component ui.Component) *BuilderMainPage {
	builder.component.Component = component

	return builder
}

func (builder *BuilderMainPage) SetStartGeometryX(startGeometryX float64) *BuilderMainPage {
	builder.component.StartGeometryX = startGeometryX

	return builder
}

func (builder *BuilderMainPage) SetStartGeometryY(startGeometryY float64) *BuilderMainPage {
	builder.component.StartGeometryY = startGeometryY

	return builder
}
