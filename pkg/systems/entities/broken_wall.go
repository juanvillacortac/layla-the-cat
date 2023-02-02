package esystems

import (
	"layla/pkg/components"
	"time"

	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func UpdateBrokenWall(ecs *ecs.ECS, e *donburi.Entry) {
	t := components.TweenSeq.Get(e)
	ts := components.TimerSystem.Get(e)
	o := components.GetObject(e)
	if t == nil {
		if check := o.Check(0, -2, "player"); check != nil {
			tween := gween.NewSequence(
				gween.New(-1, 1, 0.1, ease.OutSine),
				gween.New(1, -1, 0.1, ease.OutSine),
			)
			components.TweenSeq.Set(e, tween)
			ts.After(time.Second, func() {
				spr := components.AnimatedSprite.Get(e)
				spr.InvertColor = true
				spr.Mask = true
				spr.OffsetX = 0
				components.TweenSeq.Set(e, nil)
				o.Space.Remove(o)
				ts.After(time.Second/5, func() {
					e.Remove()
				})
			})
		}
	} else {
		spr := components.AnimatedSprite.Get(e)
		v, _, done := t.Update(1.0 / 60)
		spr.OffsetX = float64(v)
		if done {
			t.Reset()
		}
	}
}
