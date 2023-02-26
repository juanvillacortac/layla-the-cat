package archetypes

import (
	"layla/pkg/components"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/examples/platformer/layers"
)

func NewWorldScreen(ecs *ecs.ECS) *donburi.Entry {
	w := ecs.World

	entry := w.Entry(ecs.Create(
		layers.Default,
		components.WorldScreen,
		components.MapRenderer,
		components.Camera,
		components.Object,
		components.Level,
		components.AnimatedSprite,
	))

	return entry
}
