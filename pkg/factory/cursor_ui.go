package factory

import (
	"layla/pkg/tags"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreateCursorUi(ecs *ecs.ECS) *donburi.Entry {
	return ecs.World.Entry(ecs.World.Create(tags.CursorUi))
}
