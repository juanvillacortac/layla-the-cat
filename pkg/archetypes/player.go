package archetypes

import (
	"layla/pkg/components"
	"layla/pkg/layers"
	"layla/pkg/tags"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func NewPlayer(ecs *ecs.ECS) *donburi.Entry {
	w := ecs.World

	entry := w.Entry(ecs.Create(
		layers.Default,
		tags.Player,
		components.Player,
		components.Camera,
		components.Object,
		components.Input,
	))

	return entry
}
