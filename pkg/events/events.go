package events

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/events"
)

type SceneEvent struct {
	Scene interface {
		Ecs() *ecs.ECS
		Update()
		Draw(screen *ebiten.Image)
	}
}

var SwitchSceneEvents = events.NewEventType[SceneEvent]()

type PauseLevelEvent struct{}

var PauseLevelEvents = events.NewEventType[PauseLevelEvent]()

type RestartLevelEvent struct {
	Deaths int
}

var RestartLevelEvents = events.NewEventType[RestartLevelEvent]()

var WinLevelEvents = events.NewEventType[struct{}]()

var LoadLevelEvents = events.NewEventType[int]()

var ExitLevelEvents = events.NewEventType[struct{}]()
