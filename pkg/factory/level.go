package factory

import (
	"layla/pkg/archetypes"
	"layla/pkg/assets"
	"layla/pkg/components"
	"layla/pkg/maps"

	"github.com/solarlune/ldtkgo"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreateLevel(ecs *ecs.ECS, name string) (level *donburi.Entry, space *donburi.Entry) {
	level = archetypes.NewLevel(ecs)

	project, err := assets.LoadMap(name)
	if err != nil {
		panic(err)
	}
	renderer := maps.NewEbitenRenderer(maps.NewDiskLoader())
	renderer.Render(project.Levels[0])
	components.Level.SetValue(level, components.LevelData{
		LdtkProject: project,
		Renderer:    renderer,
	})

	space = CreateSpace(ecs, project.Levels[0].Width, project.Levels[0].Height)

	InitLevelEntities(ecs, space, project.Levels[0])
	InitLevelGrid(ecs, space, project.Levels[0])

	return
}

func InitLevelEntities(ecs *ecs.ECS, space *donburi.Entry, level *ldtkgo.Level) {
	for _, layer := range level.Layers {
		switch layer.Type {
		case ldtkgo.LayerTypeEntity:
			for _, entity := range layer.Entities {
				if entity.Identifier == "Player" {
					components.AddToSpace(space,
						CreatePlayer(ecs, float64(entity.Position[0]), float64(entity.Position[1]), level.Width, level.Height),
					)
				}
			}
		}
	}
}

func InitLevelGrid(ecs *ecs.ECS, space *donburi.Entry, level *ldtkgo.Level) {
	for _, layer := range level.Layers {
		switch layer.Type {
		case ldtkgo.LayerTypeIntGrid:
			for _, tile := range layer.IntGrid {
				switch tile.Value {
				case 1:
					components.AddToSpace(space,
						CreateWall(ecs, resolv.NewObject(float64(tile.Position[0]), float64(tile.Position[1]), 16, 16, "solid")),
					)
				}
			}
		}
	}
}
