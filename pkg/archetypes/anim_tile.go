package archetypes

import (
	"layla/pkg/components"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func NewAnimTile(ecs *ecs.ECS) *donburi.Entry {
	w := ecs.World

	entry := w.Entry(ecs.World.Create(
		components.AnimatedTile,
	))

	return entry
}

func NewAnimTileGroup(ecs *ecs.ECS) *donburi.Entry {
	w := ecs.World

	entry := w.Entry(ecs.World.Create(
		components.AnimatedTilesGroup,
	))

	return entry
}
