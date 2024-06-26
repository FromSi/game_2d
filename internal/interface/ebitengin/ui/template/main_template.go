package template

import (
	"github.com/fromsi/game_2d/internal/interface/ebitengin/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type MainTemplate struct {
	componentData  *ui.ComponentData
	startGeometryX float64
	startGeometryY float64
}

func (template *MainTemplate) GetStartGeometryX() float64 {
	return template.startGeometryX
}

func (template *MainTemplate) GetStartGeometryY() float64 {
	return template.startGeometryY
}

func (template *MainTemplate) GetComponentData() *ui.ComponentData {
	return template.componentData
}

func (template *MainTemplate) SetComponentData(componentData *ui.ComponentData) {
	template.componentData = componentData
}

func (template *MainTemplate) SetStartGeometryX(startGeometryX float64) {
	template.startGeometryX = startGeometryX
	componentDataIterator := template.componentData.GetNormalComponentDataIterator()

	for componentDataIterator.HasNext() {
		component := *componentDataIterator.GetNext()

		component.SetStartGeometryX(startGeometryX + component.GetStartGeometryX())
		template.componentData.UpdateById(component, componentDataIterator.GetId())
	}
}

func (template *MainTemplate) SetStartGeometryY(startGeometryY float64) {
	template.startGeometryY = startGeometryY
	componentDataIterator := template.componentData.GetNormalComponentDataIterator()

	for componentDataIterator.HasNext() {
		component := *componentDataIterator.GetNext()

		component.SetStartGeometryY(startGeometryY + component.GetStartGeometryY())
		template.componentData.UpdateById(component, componentDataIterator.GetId())
	}
}

func (template *MainTemplate) GetEndGeometryX() float64 {
	return math.MaxFloat64
}

func (template *MainTemplate) GetEndGeometryY() float64 {
	return math.MaxFloat64
}

func (template *MainTemplate) OnDraw(Screen *ebiten.Image) {
	componentDataIterator := template.componentData.GetReverseComponentDataIterator()

	for componentDataIterator.HasNext() {
		component := componentDataIterator.GetNext()

		(*component).OnDraw(Screen)
	}
}

func (template *MainTemplate) HandleClick(geometryX, geometryY float64) bool {
	if !template.IsWithin(geometryX, geometryY) {
		return false
	}

	componentDataIterator := template.componentData.GetNormalComponentDataIterator()

	for componentDataIterator.HasNext() {
		component := componentDataIterator.GetNext()

		if (*component).HandleClick(geometryX, geometryY) {
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
			componentData:  ui.NewBuilderComponentData().GetComponentData(),
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

func (builder *BuilderMainTemplate) SetComponentData(componentData *ui.ComponentData) *BuilderMainTemplate {
	builder.component.componentData = componentData

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
