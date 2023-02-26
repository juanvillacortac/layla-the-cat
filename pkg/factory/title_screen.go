package factory

import (
	"layla/pkg/archetypes"
	"layla/pkg/assets"
	"layla/pkg/components"
	// "layla/pkg/maps"

	"github.com/solarlune/ebitick"
	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreateTitleScreen(ecs *ecs.ECS) *donburi.Entry {
	p := archetypes.NewTitleScreen(ecs)

	_, h := assets.LogoSprite.Size()

	// project, err := assets.LoadMap("world-1")
	// if err != nil {
	// 	panic(err)
	// }

	// renderer := maps.NewEbitenRenderer(maps.NewDiskLoader())
	// lvl := project.Levels[0]
	// renderer.Render(lvl)

	// components.MapRenderer.Set(p, &components.MapRendererData{
	// 	Renderer: renderer,
	// 	Centered: true,
	// })

	// if lvl.BGImage != nil {
	// 	AppendBackground(p, lvl.BGImage.Path)
	// }

	components.TitleScreen.SetValue(p, components.TitleScreenData{
		BgTween: gween.New(0, 80, 0.1, ease.Linear),
		OffsetTween: gween.NewSequence(
			gween.New(float32(h-components.PAUSE_UI_GAP), 0, 0.8, ease.OutBounce),
		),
	})
	components.TimerSystem.Set(p, ebitick.NewTimerSystem())

	return p
}
