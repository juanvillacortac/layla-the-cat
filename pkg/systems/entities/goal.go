package esystems

import (
	"layla/pkg/assets"
	"layla/pkg/components"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func UpdateGoal(ecs *ecs.ECS, e *donburi.Entry) {
	o := components.GetObject(e)
	spr := components.AnimatedSprite.Get(e)
	if v, ok := o.Data.(bool); ok && v {
		spr.Anim.Resume()
	}
}

func DrawGoal(ecs *ecs.ECS, e *donburi.Entry, screen *ebiten.Image) {
	o := components.GetObject(e)
	opt := &ebiten.DrawImageOptions{}

	opt.GeoM.Translate(o.X-4, o.Y-1)
	if camera := components.GetCamera(ecs); camera != nil {
		opt.GeoM.Translate(-math.Round(camera.X), -math.Round(camera.Y))
	}
	screen.DrawImage(assets.GoalSprite, opt)
}
