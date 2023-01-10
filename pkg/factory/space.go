package factory

import (
	"layla/pkg/archetypes"
	"layla/pkg/components"

	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreateSpace(ecs *ecs.ECS, width, height int) *donburi.Entry {
	space := archetypes.NewSpace(ecs)

	spaceData := resolv.NewSpace(width, height, 16, 16)
	components.Space.Set(space, spaceData)

	return space
}
