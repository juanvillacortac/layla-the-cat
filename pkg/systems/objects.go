package systems

import (
	"layla/pkg/components"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func UpdateObjects(ecs *ecs.ECS) {
	components.Object.Each(ecs.World, func(e *donburi.Entry) {
		obj := components.GetObject(e)
		obj.Update()
	})
}
