package factory

import (
	"layla/pkg/archetypes"
	"layla/pkg/components"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreatePushable(ecs *ecs.ECS, x, y float64, tile *ebiten.Image) *donburi.Entry {
	pushable := archetypes.NewPushable(ecs)

	obj := resolv.NewObject(x, y, 16, 16, "pushable")
	obj.SetShape(resolv.NewRectangle(0, 0, 16, 16))

	components.SetObject(pushable, obj)
	components.Pushable.SetValue(pushable, components.PushableData{
		Tile: tile,
	})

	return pushable
}
