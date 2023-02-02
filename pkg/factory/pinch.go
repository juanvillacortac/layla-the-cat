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
	PinchGrid *ganim8.Grid
	PinchAnim *ganim8.Animation
)

func init() {
	w, h := assets.PinchSprite.Size()
	PinchGrid = ganim8.NewGrid(18, 9, w, h)
}

func CreatePinch(ecs *ecs.ECS, x, y float64, flipY bool) *donburi.Entry {
	pinch := archetypes.NewPinch(ecs)

	if !flipY {
		y += 8
	}
	obj := resolv.NewObject(x, y+2, 16, 6, "enemy")

	components.SetObject(pinch, obj)
	components.Entity.Set(pinch, &components.EntityData{Identifier: string(entities.Pinch)})
	components.AnimatedSprite.Set(pinch, &components.AnimatedSpriteData{
		X:          x - 1,
		Y:          y,
		FlipY:      flipY,
		HookCamera: true,
		Anim:       ganim8.New(assets.PinchSprite, PinchGrid.Frames("1-2", 1), time.Millisecond*600),
	})

	return pinch
}
