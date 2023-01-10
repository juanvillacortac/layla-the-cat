package archetypes

import (
	"layla/pkg/components"
	"layla/pkg/layers"
	"layla/pkg/tags"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func NewCamera(ecs *ecs.ECS) *donburi.Entry {
	w := ecs.World

	entry := w.Entry(ecs.Create(
		layers.Default,
		tags.Camera,
		components.Camera,
		components.Object,
	))

	return entry
}
