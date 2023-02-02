package systems

import (
	"layla/pkg/components"
	"layla/pkg/tags"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/ganim8/v2"
)

func UpdateParticles(ecs *ecs.ECS) {
	tags.Particles.Each(ecs.World, func(e *donburi.Entry) {
		particles := components.Particles.Get(e)
		particles.Anim.Update()
	})
}

func drawParticles(screen *ebiten.Image, ecs *ecs.ECS, particle *components.ParticlesData) {
	opt := ganim8.DrawOpts(particle.X, particle.Y)
	if camera := components.GetCamera(ecs); camera != nil {
		x, y := particle.X-math.Round(camera.X), particle.Y-math.Round(camera.Y)
		opt.SetPos(x, y)
	}
	particle.Anim.Draw(screen, opt)
}

func DrawParticles(layer components.ParticlesLayer) func(ecs *ecs.ECS, screen *ebiten.Image) {
	return func(ecs *ecs.ECS, screen *ebiten.Image) {
		tags.Particles.Each(ecs.World, func(e *donburi.Entry) {
			particle := components.Particles.Get(e)
			if particle.Layer == layer {
				drawParticles(screen, ecs, particle)
			}
		})
	}
}
