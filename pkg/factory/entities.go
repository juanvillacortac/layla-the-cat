package factory

import (
	"image"
	"layla/pkg/components"
	"layla/pkg/entities"
	"layla/pkg/maps"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/ldtkgo"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type EntityFactoryFn func(ecs *ecs.ECS, space *donburi.Entry, ctx *maps.LevelCtx, layer *ldtkgo.Layer, entity *ldtkgo.Entity)

var Entities map[entities.EntityType]EntityFactoryFn = map[entities.EntityType]EntityFactoryFn{
	entities.Player: func(ecs *ecs.ECS, space *donburi.Entry, ctx *maps.LevelCtx, layer *ldtkgo.Layer, entity *ldtkgo.Entity) {
		components.AddToSpace(space,
			CreatePlayer(ecs, float64(entity.Position[0]), float64(entity.Position[1]), ctx.Level.Width, ctx.Level.Height),
		)
	},
	entities.Collectable: func(ecs *ecs.ECS, space *donburi.Entry, ctx *maps.LevelCtx, layer *ldtkgo.Layer, entity *ldtkgo.Entity) {
		components.AddToSpace(space,
			CreateCollectable(ecs, float64(entity.Position[0]), float64(entity.Position[1])),
		)
	},
	entities.Saw: func(ecs *ecs.ECS, space *donburi.Entry, ctx *maps.LevelCtx, layer *ldtkgo.Layer, entity *ldtkgo.Entity) {
		components.AddToSpace(space,
			CreateSaw(ecs, float64(entity.Position[0]), float64(entity.Position[1])),
		)
	},
	entities.Pinch: func(ecs *ecs.ECS, space *donburi.Entry, ctx *maps.LevelCtx, layer *ldtkgo.Layer, entity *ldtkgo.Entity) {
		flipY := false
		for _, prop := range entity.Properties {
			if prop.Identifier == "FlipY" {
				flipY = prop.AsBool()
			}
		}
		components.AddToSpace(space,
			CreatePinch(ecs, float64(entity.Position[0]), float64(entity.Position[1]), flipY),
		)
	},
	entities.BrokenWall: func(ecs *ecs.ECS, space *donburi.Entry, ctx *maps.LevelCtx, layer *ldtkgo.Layer, entity *ldtkgo.Entity) {
		var tile *ebiten.Image
		for _, prop := range entity.Properties {
			if prop.Identifier != "Tile" {
				continue
			}
			data := prop.AsMap()
			tileset := ctx.Tilesets[int(data["tilesetUid"].(float64))]
			x, y := int(data["x"].(float64)), int(data["y"].(float64))
			w, h := int(data["w"].(float64)), int(data["h"].(float64))
			tile = ebiten.NewImageFromImage(tileset.SubImage(image.Rect(x, y, x+w, y+h)))
		}
		obj := resolv.NewObject(float64(entity.Position[0]), float64(entity.Position[1]), 16, 16, "solid", "broken")
		components.AddToSpace(space,
			CreateBrokenWall(ecs, obj, tile),
		)
	},
	// "Pushable": func(ecs *ecs.ECS, space *donburi.Entry, ctx *maps.LevelCtx, layer *ldtkgo.Layer, entity *ldtkgo.Entity) {
	// 	for _, prop := range entity.Properties {
	// 		if prop.Identifier != "Tile" {
	// 			continue
	// 		}
	// 		data := prop.AsMap()
	// 		tileset := ctx.Tilesets[int(data["tilesetUid"].(float64))]
	// 		tile := tileset.SubImage(image.Rectangle{
	// 			Min: image.Point{
	// 				X: int(data["x"].(float64)),
	// 				Y: int(data["y"].(float64)),
	// 			},
	// 			Max: image.Point{
	// 				X: int(data["x"].(float64)) + int(data["w"].(float64)),
	// 				Y: int(data["y"].(float64)) + int(data["h"].(float64)),
	// 			},
	// 		}).(*ebiten.Image)
	//
	// 		components.AddToSpace(space,
	// 			CreatePushable(ecs, float64(entity.Position[0]), float64(entity.Position[1]), tile),
	// 		)
	// 	}
	// },
}
