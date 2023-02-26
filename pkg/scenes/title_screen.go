package scenes

import (
	"layla/pkg/audio"
	"layla/pkg/events"
	"layla/pkg/factory"
	"layla/pkg/systems"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	e "github.com/yohamta/donburi/features/events"
)

type TitleScreenScene struct {
	ecs  *ecs.ECS
	main *ecs.ECS
}

func NewTitleScreenScene(main *ecs.ECS) *TitleScreenScene {
	menu := &TitleScreenScene{main: main, ecs: ecs.NewECS(donburi.NewWorld())}

	systems.AddSystems(menu.ecs)

	audio.PlayBGM("title.mp3")
	factory.CreateTitleScreen(menu.ecs)
	factory.CreateGridBg(menu.ecs)
	factory.CreateTransition(menu.ecs, false, func() {
	})

	events.LoadLevelEvents.Subscribe(menu.ecs.World, func(w donburi.World, event int) {
		events.SwitchSceneEvents.Publish(menu.main.World, events.SceneEvent{
			Scene: NewWorldScreenScene(menu.main),
		})
	})

	return menu
}

func (m *TitleScreenScene) Update() {
	m.ecs.Update()

	e.ProcessAllEvents(m.ecs.World)
}

func (m *TitleScreenScene) Draw(screen *ebiten.Image) {
	m.ecs.Draw(screen)
}

func (m *TitleScreenScene) Ecs() *ecs.ECS {
	return m.ecs
}
