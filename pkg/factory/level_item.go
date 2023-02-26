package factory

import (
	"layla/pkg/archetypes"
	"layla/pkg/assets"
	"layla/pkg/components"
	"layla/pkg/entities"
	"time"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/ganim8/v2"
)

func CreateLevelItem(ecs *ecs.ECS, m string, n int, x, y float64) *donburi.Entry {
	e := archetypes.NewLevelItem(ecs)

	img, _, err := ebitenutil.NewImageFromFileSystem(assets.TilesetsFS, "tilesets/maps_tileset.png")
	if err != nil {
		panic(err)
	}

	w, h := img.Size()
	grid := ganim8.NewGrid(16, 16, w, h)
	anim := ganim8.New(img, grid.Frames("1-3", 2), time.Millisecond*50)
	anim.Pause()

	obj := resolv.NewObject(x, y, 16, 16)

	components.SetObject(e, obj)
	components.Entity.Set(e, &components.EntityData{Identifier: string(entities.LevelItem), Layer: components.EntityFrontLayer})
	components.LevelItem.Set(e, &components.LevelItemData{
		Map:    m,
		Number: n,
	})
	components.AnimatedSprite.Set(e, &components.AnimatedSpriteData{
		X:          obj.X,
		Y:          obj.Y,
		HookCamera: true,
		Anim:       anim.Clone(),
	})

	return e
}
