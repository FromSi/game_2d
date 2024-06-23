package template

import (
	"github.com/fromsi/game_2d/internal/interface/ebitengin/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type MainTemplate struct {
	components     []ui.Component
	startGeometryX float64
	startGeometryY float64
}

func (template *MainTemplate) GetStartGeometryX() float64 {
	return template.startGeometryX
}

func (template *MainTemplate) GetStartGeometryY() float64 {
	return template.startGeometryY
}

func (template *MainTemplate) AddComponent(component ui.Component) {
	component.SetStartGeometryX(template.startGeometryX + component.GetStartGeometryX())
	component.SetStartGeometryY(template.startGeometryY + component.GetStartGeometryY())

	template.components = append(template.components, component)
}

func (template *MainTemplate) SetStartGeometryX(startGeometryX float64) {
	template.startGeometryX = startGeometryX

	for index, component := range template.components {
		template.components[index].SetStartGeometryX(startGeometryX + component.GetStartGeometryX())
	}
}

func (template *MainTemplate) SetStartGeometryY(startGeometryY float64) {
	template.startGeometryY = startGeometryY

	for index, component := range template.components {
		template.components[index].SetStartGeometryY(startGeometryY + component.GetStartGeometryY())
	}
}

func (template *MainTemplate) GetEndGeometryX() float64 {
	return math.MaxFloat64
}

func (template *MainTemplate) GetEndGeometryY() float64 {
	return math.MaxFloat64
}

func (template *MainTemplate) OnDraw(Screen *ebiten.Image) {
	for i := len(template.components) - 1; i >= 0; i-- {
		component := template.components[i]

		component.OnDraw(Screen)
	}
}

func (template *MainTemplate) HandleClick(geometryX, geometryY float64) bool {
	if !template.IsWithin(geometryX, geometryY) {
		return false
	}

	for i := len(template.components) - 1; i >= 0; i-- {
		component := template.components[i]

		if component.HandleClick(geometryX, geometryY) {
			return true
		}
	}

	return false
}

func (template *MainTemplate) IsWithin(geometryX, geometryY float64) bool {
	isWithinX := (template.startGeometryX <= geometryX) && (template.GetEndGeometryX() >= geometryX)
	isWithinY := (template.startGeometryY <= geometryY) && (template.GetEndGeometryY() >= geometryY)

	return isWithinX && isWithinY
}

type BuilderMainTemplate struct {
	component *MainTemplate
}

func NewBuilderMainTemplate() *BuilderMainTemplate {
	return &BuilderMainTemplate{
		component: &MainTemplate{
			startGeometryX: 0,
			startGeometryY: 0,
		},
	}
}

func (builder *BuilderMainTemplate) GetComponent() *MainTemplate {
	builder.component.SetStartGeometryX(builder.component.startGeometryX)
	builder.component.SetStartGeometryY(builder.component.startGeometryY)

	return builder.component
}

func (builder *BuilderMainTemplate) SetComponents(components []ui.Component) *BuilderMainTemplate {
	builder.component.components = components

	return builder
}

func (builder *BuilderMainTemplate) SetStartGeometryX(startGeometryX float64) *BuilderMainTemplate {
	builder.component.startGeometryX = startGeometryX

	return builder
}

func (builder *BuilderMainTemplate) SetStartGeometryY(startGeometryY float64) *BuilderMainTemplate {
	builder.component.startGeometryY = startGeometryY

	return builder
}
