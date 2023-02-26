package factory

import (
	"image/color"
	"layla/pkg/archetypes"
	"layla/pkg/components"
	"layla/pkg/config"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/ganim8/v2"
)

var confettiImages []*ebiten.Image = make([]*ebiten.Image, 0)

var confettiColors []color.Color = []color.Color{
	color.White,
	color.RGBA{254, 174, 52, 255},
	color.RGBA{255, 0, 68, 255},
	color.RGBA{99, 199, 77, 255},
	color.RGBA{0, 149, 233, 255},
	color.RGBA{181, 80, 136, 255},
	color.RGBA{254, 231, 97, 255},
}

func init() {
	for _, c := range confettiColors {
		img := ebiten.NewImage(1, 1)
		img.Fill(c)
		confettiImages = append(confettiImages, img)
	}
}

func CreateConfetti(ecs *ecs.ECS) *donburi.Entry {
	c := archetypes.NewConfetti(ecs)

	g := ganim8.NewGrid(1, 1, 1, 1)

	components.AnimatedSprite.Set(c, &components.AnimatedSpriteData{
		X:    float64(rand.Intn(config.Width)),
		Y:    0,
		Anim: ganim8.New(confettiImages[rand.Intn(len(confettiImages))], g.Frames(1, 1), time.Second),
	})

	tween := gween.NewSequence(
		gween.New(-1, 1, 0.1, ease.OutSine),
	)
	tween.SetYoyo(true)
	tween.SetLoop(-1)
	components.TweenSeq.Set(c, tween)

	return c
}
