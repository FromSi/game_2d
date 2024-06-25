package page

import (
	"fmt"
	"github.com/fromsi/game_2d/internal/interface/ebitengin/ui"
	"github.com/fromsi/game_2d/internal/interface/ebitengin/ui/atom"
	"github.com/fromsi/game_2d/internal/interface/ebitengin/ui/template"
	"github.com/hajimehoshi/ebiten/v2"
)

type MainPage struct {
	componentData *ui.ComponentData
}

func (page *MainPage) OnDraw(Screen *ebiten.Image) {
	componentDataIterator := page.componentData.GetNormalComponentDataIterator()

	for componentDataIterator.HasNext() {
		component := componentDataIterator.GetNext()

		(*component).OnDraw(Screen)
	}
}

func (page *MainPage) HandleClick(geometryX, geometryY float64) bool {
	componentDataIterator := page.componentData.GetNormalComponentDataIterator()

	for componentDataIterator.HasNext() {
		component := componentDataIterator.GetNext()

		if (*component).HandleClick(geometryX, geometryY) {
			return true
		}
	}

	return false
}

type BuilderMainPage struct {
	component *MainPage
}

func NewBuilderMainPage() *BuilderMainPage {
	textHelloWorld := atom.
		NewBuilderTextComponent().
		SetText("Hello World!").
		SetSize(30).
		SetOnClick(func() { fmt.Println("Hey") }).
		GetComponent()

	componentDataForData := ui.
		NewBuilderComponentData().
		AddComponent(textHelloWorld).
		GetComponentData()

	templateMain := template.
		NewBuilderMainTemplate().
		SetStartGeometryX(30).
		SetComponentData(componentDataForData).
		GetComponent()

	componentDataForMainPage := ui.
		NewBuilderComponentData().
		AddComponent(templateMain).
		GetComponentData()

	return &BuilderMainPage{
		component: &MainPage{componentData: componentDataForMainPage},
	}
}

func (builder *BuilderMainPage) GetComponent() *MainPage {
	return builder.component
}
