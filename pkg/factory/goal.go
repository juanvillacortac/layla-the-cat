package factory

import (
	"layla/pkg/archetypes"
	"layla/pkg/assets"
	"layla/pkg/components"
	"layla/pkg/entities"
	"time"

	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/ganim8/v2"
)

func CreateGoal(ecs *ecs.ECS, x, y float64) *donburi.Entry {
	goal := archetypes.NewGoal(ecs)

	obj := resolv.NewObject(x-4, y, 24, 16, "goal")
	components.Entity.Set(goal, &components.EntityData{Identifier: string(entities.Goal), Layer: components.EntityFrontLayer})
	components.SetObject(goal, obj)

	obj.Data = false

	w, h := assets.GoalSprite.Size()
	g := ganim8.NewGrid(32, 33, w, h)
	anim := ganim8.New(assets.GoalSprite, g.Frames("1-10", 1), time.Millisecond*40, func(anim *ganim8.Animation, loops int) {
		anim.PauseAtEnd()
	})
	anim.Pause()
	components.AnimatedSprite.Set(goal, &components.AnimatedSpriteData{
		X:          obj.X - 4,
		Y:          obj.Y - 16,
		FlipY:      false,
		HookCamera: true,
		Anim:       anim,
	})

	return goal
}
