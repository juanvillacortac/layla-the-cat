package factory

import (
	"layla/pkg/archetypes"
	"layla/pkg/components"
	"layla/pkg/entities"

	"github.com/solarlune/resolv"
	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreateCollectable(ecs *ecs.ECS, x, y float64) *donburi.Entry {
	collectable := archetypes.NewCollectable(ecs)

	obj := resolv.NewObject(x+3, y+3, 13, 13, "collectable")
	obj.SetShape(resolv.NewRectangle(0, 0, 13, 13))

	components.SetObject(collectable, obj)

	components.Entity.Set(collectable, &components.EntityData{Identifier: string(entities.Collectable)})

	tween := gween.NewSequence(
		gween.New(-2, 2, 1, ease.OutSine),
		gween.New(2, -2, 1, ease.OutSine),
	)

	components.TweenSeq.Set(collectable, tween)

	return collectable
}
