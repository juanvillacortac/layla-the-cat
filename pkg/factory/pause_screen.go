package factory

import (
	"bytes"
	"image"
	"layla/pkg/archetypes"
	"layla/pkg/assets"
	"layla/pkg/components"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/ganim8/v2"
)

func CreatePauseScreen(ecs *ecs.ECS) *donburi.Entry {
	p := archetypes.NewPauseScreen(ecs)

	img, _, err := image.Decode(bytes.NewReader(assets.PauseFramePng))
	if err != nil {
		panic(err)
	}
	frameSprite := ebiten.NewImageFromImage(img)

	frameGrid := ganim8.NewGrid(28, 28, 28*4, 28)
	frameAnim := ganim8.New(frameSprite, frameGrid.Frames("1-4", 1), time.Millisecond*80)

	components.PauseScreen.SetValue(p, components.PauseScreenData{
		BgTween:     gween.New(0, 80, 0.1, ease.Linear),
		OffsetTween: gween.New(10, 0, 0.4, ease.OutElastic),
		Selected:    components.PauseScreenResume,
	})

	components.AnimatedSprite.Set(p, &components.AnimatedSpriteData{
		X:       float64(components.PAUSE_UI_RESUME_X),
		Y:       float64(components.PAUSE_UI_RESUME_Y),
		OffsetX: -2,
		OffsetY: -2 - 10,
		Anim:    frameAnim.Clone(),
	})

	CreateCursorUi(ecs)

	return p
}
