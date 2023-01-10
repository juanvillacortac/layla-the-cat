package components

import (
	"layla/pkg/tags"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type CameraData struct {
	X float64
	Y float64
	W int
	H int
}

var Camera = donburi.NewComponentType[CameraData]()

func GetCamera(ecs *ecs.ECS) *CameraData {
	entry, ok := tags.Player.First(ecs.World)
	if !ok {
		return nil
	}
	return Camera.Get(entry)
}
