package factory

import (
	"bytes"
	"image"
	"layla/pkg/archetypes"
	"layla/pkg/assets"
	"layla/pkg/components"
	"layla/pkg/config"
	"layla/pkg/data"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/ganim8/v2"
)

func CreateStageCleared(ecs *ecs.ECS) *donburi.Entry {
	p := archetypes.NewStageCleared(ecs)

	if levelEntry, ok := components.Level.First(ecs.World); ok {
		level := components.Level.Get(levelEntry)
		data.SavedData.Levels[level.Number].Cleared = true
		nextLevel := level.Number + 1
		if nextLevel > len(data.Levels)-1 {
			nextLevel = len(data.Levels) - 1
		}
		if lvl, ok := data.SavedData.Levels[nextLevel]; ok && !lvl.Unlocked {
			data.SavedData.LastPlayedLevel = nextLevel
			data.SavedData.Levels[nextLevel].Unlocked = true
		}
	}

	img, _, err := image.Decode(bytes.NewReader(assets.PauseFramePng))
	if err != nil {
		panic(err)
	}
	frameSprite := ebiten.NewImageFromImage(img)

	frameGrid := ganim8.NewGrid(28, 28, 28*4, 28)
	frameAnim := ganim8.New(frameSprite, frameGrid.Frames("1-4", 1), time.Millisecond*80)

	components.StageCleared.SetValue(p, components.StageClearedData{
		BgTween:     gween.New(0, 80, 0.1, ease.Linear),
		OffsetTween: gween.New(float32(config.Height)/1.5, 0, 1, ease.OutBounce),
		Selected:    components.StageClearedResume,
	})

	components.AnimatedSprite.Set(p, &components.AnimatedSpriteData{
		X:       float64(components.STAGE_CLEARED_UI_RESUME_X),
		Y:       float64(components.STAGE_CLEARED_UI_RESUME_Y),
		OffsetX: -2,
		OffsetY: -2 - 10,
		Anim:    frameAnim.Clone(),
	})

	return p
}
