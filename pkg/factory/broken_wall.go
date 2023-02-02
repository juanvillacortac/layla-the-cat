package factory

import (
	"layla/pkg/archetypes"
	"layla/pkg/components"
	"layla/pkg/entities"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/ebitick"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/ganim8/v2"
)

func CreateBrokenWall(ecs *ecs.ECS, obj *resolv.Object, tile *ebiten.Image) *donburi.Entry {
	wall := archetypes.NewBrokenWall(ecs)

	components.SetObject(wall, obj)

	ts := ebitick.NewTimerSystem()
	components.TimerSystem.Set(wall, ts)
	g := ganim8.NewGrid(16, 16, 16, 16)
	components.Entity.Set(wall, &components.EntityData{Identifier: string(entities.BrokenWall), Layer: components.EntityFrontLayer})
	components.AnimatedSprite.Set(wall, &components.AnimatedSpriteData{
		X:          obj.X,
		Y:          obj.Y,
		FlipY:      false,
		HookCamera: true,
		Anim:       ganim8.New(tile, g.Frames(1, 1), time.Second),
	})
	components.TweenSeq.Set(wall, nil)

	return wall
}
