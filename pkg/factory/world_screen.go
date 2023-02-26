package factory

import (
	"bytes"
	"image"
	_ "image/png"
	"layla/pkg/archetypes"
	"layla/pkg/assets"
	"layla/pkg/components"
	"layla/pkg/config"
	"layla/pkg/data"
	"layla/pkg/maps"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/ganim8/v2"
)

func CreateWorldScreen(ecs *ecs.ECS) *donburi.Entry {
	p := archetypes.NewWorldScreen(ecs)

	img, _, err := image.Decode(bytes.NewReader(assets.LevelFramePng))
	if err != nil {
		panic(err)
	}
	frameSprite := ebiten.NewImageFromImage(img)

	frameGrid := ganim8.NewGrid(20, 20, 20*4, 20)
	frameAnim := ganim8.New(frameSprite, frameGrid.Frames("1-4", 1), time.Millisecond*80)

	project, err := assets.LoadMap("world-1")
	if err != nil {
		panic(err)
	}

	renderer := maps.NewEbitenRenderer(maps.NewDiskLoader())
	lvl := project.Levels[0]
	renderer.Render(lvl)

	components.MapRenderer.Set(p, &components.MapRendererData{
		Renderer: renderer,
	})

	components.WorldScreen.Set(p, &components.WorldScreenData{
		Selected: data.SavedData.LastPlayedLevel,
	})

	ctx := maps.NewLevelCtx(project)
	components.Level.Set(p, &components.LevelData{
		LdtkProject: ctx,
		Deaths:      0,
	})

	InitLevelEntities(ecs, p)

	obj := resolv.NewObject(0, float64(lvl.Height-config.Height)/2, float64(config.Width), float64(config.Height))
	obj.SetShape(resolv.NewRectangle(0, 0, float64(config.Width), float64(config.Height)))

	var itemObj *resolv.Object
	donburi.NewQuery(filter.Contains(components.LevelItem)).Each(ecs.World, func(e *donburi.Entry) {
		if item := components.LevelItem.Get(e); item.Number == data.SavedData.LastPlayedLevel {
			itemObj = components.GetObject(e)
		}
	})
	if itemObj != nil {
		obj.X = itemObj.X - float64(config.Width)/2 + itemObj.W/2
		obj.Y = itemObj.Y - float64(config.Height)/2 + itemObj.H/2
	}

	components.SetObject(p, obj)

	components.Camera.SetValue(p, components.CameraData{
		X: obj.X,
		Y: obj.Y,
		W: lvl.Width,
		H: lvl.Height,
	})

	components.AnimatedSprite.Set(p, &components.AnimatedSpriteData{
		X:          itemObj.X,
		Y:          itemObj.Y,
		HookCamera: true,
		OffsetX:    -2,
		OffsetY:    -2,
		Anim:       frameAnim.Clone(),
	})

	return p
}
