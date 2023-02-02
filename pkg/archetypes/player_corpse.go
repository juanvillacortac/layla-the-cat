package archetypes

import (
	"layla/pkg/components"
	"layla/pkg/layers"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func NewPlayerCorpse(ecs *ecs.ECS) *donburi.Entry {
	w := ecs.World

	entry := w.Entry(ecs.Create(
		layers.Default,
		components.Entity,
		components.PlayerCorpse,
		components.TimerSystem,
	))

	return entry
}
