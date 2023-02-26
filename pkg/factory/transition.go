package factory

import (
	"layla/pkg/archetypes"
	"layla/pkg/components"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreateTransition(ecs *ecs.ECS, outro bool, action func()) *donburi.Entry {
	transition := archetypes.NewTransition(ecs)

	components.Transition.SetValue(transition, components.TransitionData{
		X:         -64,
		Outro:     outro,
		EndAction: &action,
		Speed:     12,
	})

	return transition
}
