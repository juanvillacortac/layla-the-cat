package factory

import (
	"layla/pkg/archetypes"
	"layla/pkg/components"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/ganim8/v2"
)

func CreateParticles(ecs *ecs.ECS, particleType components.ParticlesType, x, y float64, flipX bool) *donburi.Entry {
	particles := archetypes.NewParticles(ecs)

	animKind := components.ParticlesAnimations[particleType]
	anim := ganim8.New(components.ParticlesSpriteSheet, animKind.Frames, animKind.Duration, func(anim *ganim8.Animation, loops int) {
		particles.Remove()
	})

	if flipX {
		anim.Sprite().FlipH()
	}

	components.Particles.Set(particles, &components.ParticlesData{
		Type: particleType,
		X:    x,
		Y:    y,
		Anim: anim,
	})

	return particles
}
