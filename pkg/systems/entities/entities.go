package esystems

import (
	"layla/pkg/components"
	"layla/pkg/entities"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func UpdateEntities(ecs *ecs.ECS) {
	components.Entity.Each(ecs.World, func(e *donburi.Entry) {
		ent := components.Entity.Get(e)
		switch entities.EntityType(ent.Identifier) {
		case entities.Player:
			UpdatePlayer(ecs, e)
		case entities.PlayerCorpse:
			UpdatePlayerCorpse(ecs, e)
		case entities.Collectable:
			UpdateCollectable(ecs, e)
		case entities.BrokenWall:
			UpdateBrokenWall(ecs, e)
		}
	})
}

func DrawEntities(layer components.EntityLayer) func(ecs *ecs.ECS, screen *ebiten.Image) {
	return func(ecs *ecs.ECS, screen *ebiten.Image) {
		DrawEnemy(ecs, screen)
		components.Entity.Each(ecs.World, func(e *donburi.Entry) {
			ent := components.Entity.Get(e)
			if ent.Layer == layer {
				switch entities.EntityType(ent.Identifier) {
				case entities.Player:
					DrawPlayer(ecs, e, screen)
				case entities.PlayerCorpse:
					DrawPlayerCorpse(ecs, e, screen)
				case entities.Collectable:
					DrawCollectable(ecs, e, screen)
				}
			}
		})
	}
}
