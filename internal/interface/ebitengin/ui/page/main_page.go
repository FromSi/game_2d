package page

import (
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
	mainPage := &MainPage{}

	textHelloWorld := atom.
		NewBuilderTextComponent().
		SetText("Hello World!").
		SetSize(30).
		GetComponent()

	tileComponent := atom.
		NewBuilderTileComponent().
		SetStartGeometryY(50).
		SetOnClick(func() { textHelloWorld.Text = "Clicked" }).
		SetScale(5).
		GetComponent()

	componentDataForData := ui.
		NewBuilderComponentData().
		AddComponent(textHelloWorld).
		AddComponent(tileComponent).
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

	mainPage.componentData = componentDataForMainPage

	return &BuilderMainPage{
		component: mainPage,
	}
}

func (builder *BuilderMainPage) GetComponent() *MainPage {
	return builder.component
}
