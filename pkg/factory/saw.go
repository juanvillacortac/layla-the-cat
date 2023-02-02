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

var (
	SawGrid *ganim8.Grid
	SawAnim *ganim8.Animation
)

func init() {
	SawGrid = ganim8.NewGrid(34, 34, 34*3, 34)
	SawAnim = ganim8.New(assets.SawSprite, SawGrid.Frames("1-3", 1), time.Millisecond*50)
}

func CreateSaw(ecs *ecs.ECS, x, y float64) *donburi.Entry {
	saw := archetypes.NewSaw(ecs)

	obj := resolv.NewObject((x-32/2)+4, (y-32/2)+2, 32-8, 32-8, "enemy")

	components.SetObject(saw, obj)
	components.Entity.Set(saw, &components.EntityData{Identifier: string(entities.Saw)})
	components.AnimatedSprite.Set(saw, &components.AnimatedSpriteData{
		X:          obj.X - 1 - 4,
		Y:          obj.Y - 1 - 4,
		HookCamera: true,
		Anim:       SawAnim.Clone(),
	})

	return saw
}
