package components

import (
	"github.com/yohamta/donburi"
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

var AnimatedSprite = donburi.NewComponentType[AnimatedSpriteData]()
