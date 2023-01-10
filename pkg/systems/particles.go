package systems

import (
	"layla/pkg/components"
	"layla/pkg/tags"

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

func DrawParticles(ecs *ecs.ECS, screen *ebiten.Image) {
	tags.Particles.Each(ecs.World, func(e *donburi.Entry) {
		camera := components.GetCamera(ecs)
		p := components.Particles.Get(e)
		x, y := p.X-camera.X, p.Y-camera.Y
		opt := ganim8.DrawOpts(x, y)
		p.Anim.Draw(screen, opt)
	})
}
