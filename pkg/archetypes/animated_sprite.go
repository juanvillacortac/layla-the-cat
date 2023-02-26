package archetypes

import (
	"layla/pkg/components"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func NewAnimatedSprite(ecs *ecs.ECS, cmps ...donburi.IComponentType) *donburi.Entry {
	w := ecs.World

	cmps = append(cmps, components.AnimatedSprite)

	entry := w.Entry(ecs.World.Create(
		cmps...,
	))

	return entry
}
