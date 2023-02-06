package esystems

import (
	"layla/pkg/components"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/ganim8/v2"
)

func UpdateAnimatedSprites(ecs *ecs.ECS) {
	components.AnimatedSprite.Each(ecs.World, func(e *donburi.Entry) {
		spr := components.AnimatedSprite.Get(e)
		if !spr.Paused {
			spr.Anim.Sprite().SetFlipV(spr.FlipY)
			spr.Anim.Sprite().SetFlipH(spr.FlipX)
			spr.Anim.Update()
		}
	})
}

func DrawAnimatedSprites[T any](ecs *ecs.ECS, componentsType *donburi.ComponentType[T], screen *ebiten.Image) {
	componentsType.Each(ecs.World, func(e *donburi.Entry) {
		if !e.HasComponent(components.AnimatedSprite) {
			return
		}
		spr := components.AnimatedSprite.Get(e)
		if spr == nil {
			return
		}

		opt := ganim8.DrawOpts(spr.X+spr.OffsetX, spr.Y+spr.OffsetY)
		if camera := components.GetCamera(ecs); camera != nil && spr.HookCamera {
			opt.SetPos(spr.X-math.Round(camera.X)+spr.OffsetX, spr.Y-math.Round(camera.Y)+spr.OffsetY)
		}
		if spr.Mask {
			opt.ColorM.Scale(0, 0, 0, 1)
		}
		if spr.InvertColor {
			opt.ColorM.Invert()
		}
		spr.Anim.Draw(screen, opt)
	})
}
