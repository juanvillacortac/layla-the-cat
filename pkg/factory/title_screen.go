package factory

import (
	"layla/pkg/archetypes"
	"layla/pkg/assets"
	"layla/pkg/components"

	"github.com/solarlune/ebitick"
	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreateTitleScreen(ecs *ecs.ECS) *donburi.Entry {
	p := archetypes.NewTitleScreen(ecs)

	_, h := assets.LogoSprite.Size()

	components.TitleScreen.SetValue(p, components.TitleScreenData{
		BgTween: gween.New(0, 80, 0.1, ease.Linear),
		OffsetTween: gween.NewSequence(
			gween.New(float32(h-components.PAUSE_UI_GAP), 0, 0.8, ease.OutBounce),
		),
	})
	components.TimerSystem.Set(p, ebitick.NewTimerSystem())

	return p
}
