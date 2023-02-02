package archetypes

import (
	"layla/pkg/components"
	"layla/pkg/layers"
	"layla/pkg/tags"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func NewLevel(ecs *ecs.ECS) *donburi.Entry {
	w := ecs.World

	entry := w.Entry(ecs.Create(
		layers.Default,
		tags.Level,
		components.Level,
		components.Space,
	))

	return entry
}
