package components

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/ganim8/v2"
)

type AnimatedSpriteData struct {
	X, Y             float64
	OffsetX, OffsetY float64
	Anim             *ganim8.Animation
	Paused           bool
	FlipX            bool
	FlipY            bool
	HookCamera       bool
	Mask             bool
	InvertColor      bool
}

func (spr *AnimatedSpriteData) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	opt := ganim8.DrawOpts(spr.X+spr.OffsetX, spr.Y+spr.OffsetY)
	if camera := GetCamera(ecs); camera != nil && spr.HookCamera {
		opt.SetPos(spr.X-math.Round(camera.X)+spr.OffsetX, spr.Y-math.Round(camera.Y)+spr.OffsetY)
	}
	if spr.Mask {
		opt.ColorM.Scale(0, 0, 0, 1)
	}
	if spr.InvertColor {
		opt.ColorM.Invert()
	}
	spr.Anim.Draw(screen, opt)
}

var AnimatedSprite = donburi.NewComponentType[AnimatedSpriteData]()
