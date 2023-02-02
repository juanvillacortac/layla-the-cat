package scenes

import (
	"layla/pkg/components"
	"layla/pkg/config"
	"layla/pkg/events"
	"layla/pkg/factory"
	"layla/pkg/systems"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	e "github.com/yohamta/donburi/features/events"
	"github.com/yohamta/donburi/filter"
)

type PauseScene struct {
	ecs      *ecs.ECS
	levelEcs *ecs.ECS
	main     *ecs.ECS
}

func NewPauseScene(main *ecs.ECS, levelEcs *ecs.ECS) *PauseScene {
	p := &PauseScene{main: main, levelEcs: levelEcs, ecs: ecs.NewECS(donburi.NewWorld())}

	systems.AddSystems(p.ecs)

	events.PauseLevelEvents.Subscribe(p.ecs.World, func(w donburi.World, event events.PauseLevelEvent) {
		events.PauseLevelEvents.Publish(p.levelEcs.World, event)
	})

	events.RestartLevelEvents.Subscribe(p.ecs.World, func(w donburi.World, event events.RestartLevelEvent) {
		events.RestartLevelEvents.Publish(p.levelEcs.World, event)
	})

	events.SwitchSceneEvents.Subscribe(p.ecs.World, func(w donburi.World, event events.SceneEvent) {
		if e, ok := components.Level.First(p.levelEcs.World); ok {
			components.Level.Get(e).Renderer.Clear()
			query := donburi.NewQuery(filter.Contains())
			query.Each(p.levelEcs.World, func(e *donburi.Entry) {
				e.Remove()
			})
		}
		events.SwitchSceneEvents.Publish(p.main.World, event)
	})

	events.ExitLevelEvents.Subscribe(p.ecs.World, func(w donburi.World, event struct{}) {
		config.C.Touch = true
		if e, ok := components.Level.First(p.levelEcs.World); ok {
			components.Level.Get(e).Renderer.Clear()
			query := donburi.NewQuery(filter.Contains())
			query.Each(p.levelEcs.World, func(e *donburi.Entry) {
				e.Remove()
			})
		}
		events.SwitchSceneEvents.Publish(p.main.World, events.SceneEvent{
			Scene: NewTitleScreenScene(p.main),
		})
	})

	factory.CreatePauseScreen(p.ecs)

	return p
}

func (p *PauseScene) Update() {
	p.ecs.Update()

	e.ProcessAllEvents(p.ecs.World)
}

func (p *PauseScene) Draw(screen *ebiten.Image) {
	p.ecs.Draw(screen)
}

func (p *PauseScene) Ecs() *ecs.ECS {
	return p.ecs
}
