package archetypes

import (
	"layla/pkg/components"
	"layla/pkg/layers"
	"layla/pkg/tags"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func NewBrokenWall(ecs *ecs.ECS) *donburi.Entry {
	w := ecs.World

	entry := w.Entry(ecs.Create(
		layers.Default,
		tags.BrokenWall,
		tags.Enemy,
		components.Entity,
		components.TimerSystem,
		components.Object,
		components.AnimatedSprite,
		components.TweenSeq,
	))

	return entry
}
