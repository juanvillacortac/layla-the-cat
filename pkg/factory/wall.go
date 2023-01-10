package factory

import (
	"layla/pkg/archetypes"
	"layla/pkg/components"

	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreateWall(ecs *ecs.ECS, obj *resolv.Object) *donburi.Entry {
	wall := archetypes.NewWall(ecs)
	components.SetObject(wall, obj)
	return wall
}
