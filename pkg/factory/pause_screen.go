package factory

import (
	"layla/pkg/archetypes"
	"layla/pkg/components"

	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreatePauseScreen(ecs *ecs.ECS) *donburi.Entry {
	p := archetypes.NewPauseScreen(ecs)

	components.PauseScreen.SetValue(p, components.PauseScreenData{
		BgTween:     gween.New(0, 80, 0.1, ease.Linear),
		OffsetTween: gween.New(10, 0, 0.4, ease.OutElastic),
	})

	return p
}
