package components

import (
	"bytes"
	"image"
	_ "image/png"
	"layla/pkg/assets"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/ganim8/v2"
)

var (
	ParticlesSpriteSheet *ebiten.Image
	ParticlesGrid        *ganim8.Grid
)

type ParticlesType int
type ParticlesLayer int

const (
	ParticlesJump ParticlesType = iota
	ParticlesFall
)

const (
	ParticlesFrontLayer ParticlesLayer = iota
	ParticlesBackLayer
)

type ParticleAnimation struct {
	Frames   []*image.Rectangle
	Duration time.Duration
}

var ParticlesAnimations map[ParticlesType]ParticleAnimation

func init() {
	img, _, err := image.Decode(bytes.NewReader(assets.ParticlesPng))
	if err != nil {
		log.Fatal(err)
	}
	ParticlesSpriteSheet = ebiten.NewImageFromImage(img)
	ParticlesGrid = ganim8.NewGrid(16, 16, 16*4, 16*2)

	ParticlesAnimations = map[ParticlesType]ParticleAnimation{
		ParticlesJump: {
			Frames:   ParticlesGrid.Frames("1-4", 1),
			Duration: 80 * time.Millisecond,
		},
		ParticlesFall: {
			Frames:   ParticlesGrid.Frames("1-4", 2),
			Duration: 40 * time.Millisecond,
		},
	}
}

type ParticlesData struct {
	Type  ParticlesType
	FlipX bool
	X     float64
	Y     float64
	Anim  *ganim8.Animation
	Layer ParticlesLayer
}

var Particles = donburi.NewComponentType[ParticlesData]()
