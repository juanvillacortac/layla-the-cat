package scenes

import (
	"layla/pkg/audio"
	"layla/pkg/components"
	"layla/pkg/events"
	"layla/pkg/factory"
	"layla/pkg/systems"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	e "github.com/yohamta/donburi/features/events"
	"github.com/yohamta/donburi/filter"
)

type WorldScreenScene struct {
	ecs  *ecs.ECS
	main *ecs.ECS
}

func NewWorldScreenScene(main *ecs.ECS) *WorldScreenScene {
	menu := &WorldScreenScene{main: main, ecs: ecs.NewECS(donburi.NewWorld())}

	systems.AddSystems(menu.ecs)

	audio.PlayBGM("world.mp3")
	audio.SetBGMVolume(2)
	factory.CreateWorldScreen(menu.ecs)
	factory.CreateCursorUi(menu.ecs)
	factory.CreateTransition(menu.ecs, false, func() {})

	events.LoadLevelEvents.Subscribe(menu.ecs.World, func(w donburi.World, event int) {
		if e, ok := components.WorldScreen.First(menu.ecs.World); ok {
			components.MapRenderer.Get(e).Renderer.Clear()
			query := donburi.NewQuery(filter.Contains())
			query.Each(menu.ecs.World, func(e *donburi.Entry) {
				e.Remove()
			})
		}
		if event == -1 {
			events.SwitchSceneEvents.Publish(menu.main.World, events.SceneEvent{
				Scene: NewTitleScreenScene(menu.main),
			})
			return
		}
		events.SwitchSceneEvents.Publish(menu.main.World, events.SceneEvent{
			Scene: NewLevelScene(menu.main, event, 0),
		})
	})

	return menu
}

func (m *WorldScreenScene) Update() {
	m.ecs.Update()

	e.ProcessAllEvents(m.ecs.World)
}

func (m *WorldScreenScene) Draw(screen *ebiten.Image) {
	m.ecs.Draw(screen)
}

func (m *WorldScreenScene) Ecs() *ecs.ECS {
	return m.ecs
}
