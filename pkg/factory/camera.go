package factory

import (
	"layla/pkg/archetypes"
	"layla/pkg/components"
	"layla/pkg/config"

	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreateCamera(ecs *ecs.ECS, width, height int) *donburi.Entry {
	camera := archetypes.NewCamera(ecs)

	w, h := float64(config.C.Width), float64(config.C.Height)
	obj := resolv.NewObject(0, 0, w, h)

	components.SetObject(camera, obj)
	components.Camera.SetValue(camera, components.CameraData{
		W: width,
		H: height,
	})
	obj.SetShape(resolv.NewRectangle(0, 0, w, h))

	return camera
}
