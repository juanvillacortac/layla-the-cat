package archetypes

import (
	"layla/pkg/components"
	"layla/pkg/layers"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func NewLevelItem(ecs *ecs.ECS) *donburi.Entry {
	w := ecs.World

	entry := w.Entry(ecs.Create(
		layers.Default,
		components.Entity,
		components.Object,
		components.AnimatedSprite,
		components.LevelItem,
	))

	return entry
}
