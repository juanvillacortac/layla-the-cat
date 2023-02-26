package factory

import (
	"layla/pkg/archetypes"
	"layla/pkg/assets"
	"layla/pkg/audio"
	"layla/pkg/components"
	"layla/pkg/data"
	"layla/pkg/entities"
	"layla/pkg/maps"

	"github.com/solarlune/ldtkgo"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreateLevel(ecs *ecs.ECS, levelNumber int, deaths int) *donburi.Entry {
	level := archetypes.NewLevel(ecs)

	data.SavedData.LastPlayedLevel = levelNumber

	var name string
	if levelNumber > len(data.Levels)-1 {
		name = data.Levels[0]
		data.SavedData.LastPlayedLevel = 0
	} else {
		name = data.Levels[levelNumber]
	}

	project, err := assets.LoadMap(name)
	if err != nil {
		panic(err)
	}

	renderer := maps.NewEbitenRenderer(maps.NewDiskLoader())
	lvl := project.Levels[0]
	renderer.Render(lvl)

	ctx := maps.NewLevelCtx(project)

	components.Level.Set(level, &components.LevelData{
		LdtkProject: ctx,
		Number:      levelNumber,
		Deaths:      deaths,
	})
	components.MapRenderer.Set(level, &components.MapRendererData{
		Renderer: renderer,
	})

	components.Space.Set(level, resolv.NewSpace(ctx.Level.Width, ctx.Level.Height, 4, 4))

	if lvl.BGImage != nil {
		AppendBackground(level, lvl.BGImage.Path)
	}

	CreateGridBg(ecs)
	InitLevelEntities(ecs, level)
	InitLevelGrid(ecs, level)
	CreateTransition(ecs, false, func() {
		audio.PlayBGM("green-green.mp3")
		audio.SetBGMVolume(0.7)
	})

	return level
}

func InitLevelEntities(ecs *ecs.ECS, level *donburi.Entry) {
	ctx := components.Level.Get(level).LdtkProject
	for _, layer := range ctx.Level.Layers {
		switch layer.Type {
		case ldtkgo.LayerTypeEntity:
			for _, entity := range layer.Entities {
				if fn, ok := Entities[entities.EntityType(entity.Identifier)]; ok {
					fn(ecs, level, ctx, layer, entity)
				}
			}
		}
	}
}

func InitLevelGrid(ecs *ecs.ECS, level *donburi.Entry) {
	ctx := components.Level.Get(level).LdtkProject
	for _, layer := range ctx.Level.Layers {
		switch layer.Type {
		case ldtkgo.LayerTypeIntGrid:
			for _, tile := range layer.IntGrid {
				components.AddToSpace(level,
					CreateWall(ecs, resolv.NewObject(float64(tile.Position[0]), float64(tile.Position[1]), 16, 16, "solid")),
				)
			}
		}
	}
}
