package factory

import (
	"layla/pkg/archetypes"
	"layla/pkg/components"

	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreatePlayer(ecs *ecs.ECS, x, y float64, viewportW, viewportH int) *donburi.Entry {
	player := archetypes.NewPlayer(ecs)

	obj := resolv.NewObject(x, y, 16, 16)
	obj.SetShape(resolv.NewRectangle(0, 0, 16, 16))

	components.SetObject(player, obj)
	components.Input.Set(player, &components.InputData{})
	components.Player.SetValue(player, components.PlayerData{
		FacingRight: true,
	})
	components.Camera.SetValue(player, components.CameraData{
		W: viewportW,
		H: viewportH,
	})

	return player
}
