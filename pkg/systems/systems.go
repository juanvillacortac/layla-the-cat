package systems

import (
	"layla/pkg/components"
	"layla/pkg/layers"
	"layla/pkg/systems/entities"

	"github.com/yohamta/donburi/ecs"
)

func AddSystems(ecs *ecs.ECS) {
	ecs.AddSystem(UpdateTimers)
	ecs.AddSystem(esystems.UpdateAnimatedSprites)

	ecs.AddSystem(UpdateWorldScreen)

	ecs.AddSystem(UpdateParticles)
	ecs.AddSystem(esystems.UpdateEntities)
	ecs.AddSystem(UpdateObjects)
	ecs.AddSystem(UpdateTransitions)

	ecs.AddSystem(UpdateStageCleared)
	ecs.AddSystem(UpdatePauseScreen)
	ecs.AddSystem(UpdateTitleScreen)

	ecs.AddSystem(UpdateCamera)

	ecs.AddSystem(UpdateInput)

	ecs.AddRenderer(layers.Input, DrawInput)

	ecs.AddRenderer(layers.Background, esystems.DrawEntities(components.EntityBackgroundLayer))
	ecs.AddRenderer(layers.Background, DrawImageBackground)
	ecs.AddRenderer(layers.Background, DrawMapBackground)

	ecs.AddRenderer(layers.Default, DrawParticles(components.ParticlesBackLayer))
	ecs.AddRenderer(layers.Default, esystems.DrawEntities(components.EntityBackLayer))
	ecs.AddRenderer(layers.Default, DrawParticles(components.ParticlesFrontLayer))
	ecs.AddRenderer(layers.Foreground, DrawMapForeground)
	ecs.AddRenderer(layers.Foreground, esystems.DrawEntities(components.EntityFrontLayer))

	ecs.AddRenderer(layers.Input, DrawStageCleared)
	ecs.AddRenderer(layers.Input, DrawPauseScreen)
	ecs.AddRenderer(layers.Input, DrawTitleScreen)
	ecs.AddRenderer(layers.Input, DrawWorldScreen)

	ecs.AddRenderer(layers.Input, DrawCursorUi)

	ecs.AddRenderer(layers.Transition, DrawFlash)

	ecs.AddRenderer(layers.Transition, DrawTransitions)
}
