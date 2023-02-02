package systems

import (
	"layla/pkg/components"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func UpdateTimers(ecs *ecs.ECS) {
	components.TimerSystem.Each(ecs.World, func(e *donburi.Entry) {
		components.TimerSystem.Get(e).Update()
	})
}
