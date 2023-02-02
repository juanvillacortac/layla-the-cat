package factory

import (
	"layla/pkg/archetypes"
	"layla/pkg/components"
	"time"

	"github.com/solarlune/ebitick"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreateFlash(ecs *ecs.ECS, duration time.Duration) *donburi.Entry {
	flash := archetypes.NewFlash(ecs)

	ts := ebitick.NewTimerSystem()
	components.TimerSystem.Set(flash, ts)

	ts.After(duration, func() {
		flash.Remove()
	})

	return flash
}
