package esystems

import (
	"layla/pkg/assets"
	"layla/pkg/audio"
	"layla/pkg/components"
	"layla/pkg/factory"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func UpdateCollectable(ecs *ecs.ECS, e *donburi.Entry) {
	c := components.Collectable.Get(e)
	t := components.TweenSeq.Get(e)
	v, _, done := t.Update(1.0 / 60)
	c.OffsetY = float64(v)
	if done {
		t.Reset()
	}

	o := components.GetObject(e)

	if check := o.Check(0, 0, "player"); check != nil {
		components.ShakeCamera(ecs, 1, time.Millisecond*200)
		audio.PlaySE("sfx_3.wav")
		audio.PlaySE("sfx_2.wav")
		factory.CreateFlash(ecs, time.Millisecond*100)
		e.Remove()
	}
}

func DrawCollectable(ecs *ecs.ECS, e *donburi.Entry, screen *ebiten.Image) {
	o := components.GetObject(e)
	c := components.Collectable.Get(e)
	opt := &ebiten.DrawImageOptions{}

	opt.GeoM.Translate(o.X, o.Y+c.OffsetY)
	if camera := components.GetCamera(ecs); camera != nil {
		opt.GeoM.Translate(-math.Round(camera.X), -math.Round(camera.Y))
	}
	screen.DrawImage(assets.CollectableSprite, opt)
}
