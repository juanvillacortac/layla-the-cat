package systems

import (
	"layla/pkg/components"
	"layla/pkg/layers"
	"layla/pkg/systems/entities"

	"github.com/yohamta/donburi/ecs"
)

func AddSystems(ecs *ecs.ECS) {
	ecs.AddSystem(UpdateTimers)
	ecs.AddSystem(UpdateAnimTiles)
	ecs.AddSystem(UpdateParticles)
	ecs.AddSystem(esystems.UpdateEntities)
	ecs.AddSystem(esystems.UpdateAnimatedSprites)
	ecs.AddSystem(UpdateCamera)
	ecs.AddSystem(UpdateObjects)
	ecs.AddSystem(UpdateTransitions)

	ecs.AddSystem(UpdatePauseScreen)
	ecs.AddSystem(UpdateTitleScreen)

	ecs.AddSystem(UpdateInput)

	ecs.AddRenderer(layers.Input, DrawInput)

	ecs.AddRenderer(layers.Background, DrawAnimTiles)

	ecs.AddRenderer(layers.Background, DrawLevelBackground)

	ecs.AddRenderer(layers.Default, DrawParticles(components.ParticlesBackLayer))
	ecs.AddRenderer(layers.Default, esystems.DrawEntities(components.EntityBackLayer))
	ecs.AddRenderer(layers.Default, DrawParticles(components.ParticlesFrontLayer))
	ecs.AddRenderer(layers.Default, DrawLevelForeground)
	ecs.AddRenderer(layers.Default, esystems.DrawEntities(components.EntityFrontLayer))

	ecs.AddRenderer(layers.Default, DrawPauseScreen)
	ecs.AddRenderer(layers.Default, DrawTitleScreen)

	ecs.AddRenderer(layers.Transition, DrawFlash)

	ecs.AddRenderer(layers.Transition, DrawTransitions)
}
