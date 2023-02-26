package systems

import (
	"layla/pkg/components"
	"layla/pkg/config"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func UpdateConfetti(ecs *ecs.ECS, e *donburi.Entry) {
	t := components.TweenSeq.Get(e)
	spr := components.AnimatedSprite.Get(e)
	v, _, _ := t.Update(1.0 / 60)
	spr.OffsetX = float64(v)
	spr.Y += 1

	if spr.Y > float64(config.Height) {
		e.Remove()
	}
}
