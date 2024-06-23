package template

import (
	"github.com/fromsi/game_2d/internal/interface/ebitengin/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type MainTemplate struct {
	Components     []ui.Component
	StartGeometryX float64
	StartGeometryY float64
}

func (template *MainTemplate) GetEndGeometries() (float64, float64) {
	return math.MaxFloat64, math.MaxFloat64
}

func (template *MainTemplate) OnDraw(Screen *ebiten.Image) {
	for i := len(template.Components) - 1; i >= 0; i-- {
		component := template.Components[i]

		component.OnDraw(Screen)
	}
}

func (template *MainTemplate) HandleClick(geometryX, geometryY float64) bool {
	if !template.IsWithin(geometryX, geometryY) {
		return false
	}

	for i := len(template.Components) - 1; i >= 0; i-- {
		component := template.Components[i]

		if component.HandleClick(geometryX, geometryY) {
			return true
		}
	}

	return false
}

func (template *MainTemplate) IsWithin(geometryX, geometryY float64) bool {
	endGeometryX, endGeometryY := template.GetEndGeometries()

	isWithinX := (template.StartGeometryX <= geometryX) && (endGeometryX >= geometryX)
	isWithinY := (template.StartGeometryY <= geometryY) && (endGeometryY >= geometryY)

	return isWithinX && isWithinY
}

type BuilderMainTemplate struct {
	component *MainTemplate
}

func NewBuilderMainTemplate() *BuilderMainTemplate {
	return &BuilderMainTemplate{
		component: &MainTemplate{
			StartGeometryX: 0,
			StartGeometryY: 0,
		},
	}
}

func (builder *BuilderMainTemplate) GetComponent() *MainTemplate {
	return builder.component
}

func (builder *BuilderMainTemplate) SetComponents(components []ui.Component) *BuilderMainTemplate {
	builder.component.Components = components

	return builder
}

func (builder *BuilderMainTemplate) SetStartGeometryX(startGeometryX float64) *BuilderMainTemplate {
	builder.component.StartGeometryX = startGeometryX

	return builder
}

func (builder *BuilderMainTemplate) SetStartGeometryY(startGeometryY float64) *BuilderMainTemplate {
	builder.component.StartGeometryY = startGeometryY

	return builder
}
