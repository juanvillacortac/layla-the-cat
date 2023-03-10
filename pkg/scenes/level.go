package scenes

import (
	"layla/pkg/components"
	"layla/pkg/config"
	"layla/pkg/events"
	"layla/pkg/factory"
	"layla/pkg/layers"
	"layla/pkg/systems"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	e "github.com/yohamta/donburi/features/events"
	"github.com/yohamta/donburi/filter"
)

type LevelScene struct {
	number     int
	ecs        *ecs.ECS
	main       *ecs.ECS
	pauseScene Scene
	paused     bool
}

func NewLevelScene(main *ecs.ECS, levelNumber int, deaths int) *LevelScene {
	level := &LevelScene{main: main, ecs: ecs.NewECS(donburi.NewWorld()), number: levelNumber}

	events.PauseLevelEvents.Subscribe(level.ecs.World, func(w donburi.World, event events.PauseLevelEvent) {
		level.paused = !level.paused
		config.C.Touch = !level.paused
		if level.paused && level.pauseScene == nil {
			level.pauseScene = NewPauseScene(level.main, level.ecs)
		} else {
			level.pauseScene = nil
		}
	})

	events.RestartLevelEvents.Subscribe(level.ecs.World, func(w donburi.World, event events.RestartLevelEvent) {
		config.C.Touch = true
		if e, ok := components.Level.First(level.ecs.World); ok {
			components.MapRenderer.Get(e).Renderer.Clear()
			query := donburi.NewQuery(filter.Contains())
			query.Each(level.ecs.World, func(e *donburi.Entry) {
				e.Remove()
			})
		}
		events.SwitchSceneEvents.Publish(level.main.World, events.SceneEvent{
			Scene: NewLevelScene(level.main, level.number, event.Deaths),
		})
	})

	events.ExitLevelEvents.Subscribe(level.ecs.World, func(w donburi.World, event struct{}) {
		config.C.Touch = true
		if e, ok := components.Level.First(level.ecs.World); ok {
			components.MapRenderer.Get(e).Renderer.Clear()
			query := donburi.NewQuery(filter.Contains())
			query.Each(level.ecs.World, func(e *donburi.Entry) {
				e.Remove()
			})
		}
		events.SwitchSceneEvents.Publish(level.main.World, events.SceneEvent{
			Scene: NewWorldScreenScene(level.main),
		})
	})

	events.WinLevelEvents.Subscribe(level.ecs.World, func(w donburi.World, event struct{}) {
		config.C.Touch = true
		if e, ok := components.Level.First(level.ecs.World); ok {
			components.MapRenderer.Get(e).Renderer.Clear()
			query := donburi.NewQuery(filter.Contains())
			query.Each(level.ecs.World, func(e *donburi.Entry) {
				e.Remove()
			})
		}
		events.SwitchSceneEvents.Publish(level.main.World, events.SceneEvent{
			Scene: NewLevelScene(level.main, level.number+1, 0),
		})
	})

	systems.AddSystems(level.ecs)

	factory.CreateLevel(level.ecs, level.number, deaths)

	level.ecs.AddRenderer(layers.Transition, systems.DrawTransitions)

	return level
}

func (level *LevelScene) Ecs() *ecs.ECS {
	return level.ecs
}

func (level *LevelScene) Update() {
	if !level.paused {
		level.ecs.Update()
	}

	if level.pauseScene != nil {
		level.pauseScene.Update()
	}

	e.ProcessAllEvents(level.ecs.World)
}

func (level *LevelScene) Draw(screen *ebiten.Image) {
	level.ecs.Draw(screen)

	if level.pauseScene != nil {
		level.pauseScene.Draw(screen)
	}
}
