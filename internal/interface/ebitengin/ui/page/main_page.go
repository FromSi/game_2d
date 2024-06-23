package page

import (
	"fmt"
	"github.com/fromsi/game_2d/internal/interface/ebitengin/ui"
	"github.com/fromsi/game_2d/internal/interface/ebitengin/ui/atom"
	"github.com/fromsi/game_2d/internal/interface/ebitengin/ui/template"
	"github.com/hajimehoshi/ebiten/v2"
)

type MainPage struct {
	components []ui.Component
}

func (page *MainPage) OnDraw(Screen *ebiten.Image) {
	for i := len(page.components) - 1; i >= 0; i-- {
		component := page.components[i]

		component.OnDraw(Screen)
	}
}

func (page *MainPage) HandleClick(geometryX, geometryY float64) bool {
	for i := len(page.components) - 1; i >= 0; i-- {
		component := page.components[i]

		if component.HandleClick(geometryX, geometryY) {
			return true
		}
	}

	return false
}

type BuilderMainPage struct {
	component *MainPage
}

func NewBuilderMainPage() *BuilderMainPage {
	components := []ui.Component{
		template.
			NewBuilderMainTemplate().
			SetStartGeometryX(50).
			SetComponents([]ui.Component{
				atom.
					NewBuilderTextComponent().
					SetText("Hello World!").
					SetSize(30).
					SetOnClick(func() {
						fmt.Println("Hey")
					}).GetComponent(),
			}).GetComponent(),
	}

	return &BuilderMainPage{
		component: &MainPage{components: components},
	}
}

func (builder *BuilderMainPage) GetComponent() *MainPage {
	return builder.component
}
