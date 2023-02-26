package systems

import (
	"layla/pkg/audio"
	"layla/pkg/components"
	"layla/pkg/config"
	"layla/pkg/events"
	"layla/pkg/factory"
	"layla/pkg/input"
	"layla/pkg/platform"
	esystems "layla/pkg/systems/entities"
	"layla/pkg/tags"
	"layla/pkg/text"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	// etext "github.com/hajimehoshi/ebiten/v2/text"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

func UpdateStageCleared(ecs *ecs.ECS) {
	components.UpdateStageClearedLayout()
	components.StageCleared.Each(ecs.World, func(e *donburi.Entry) {
		tags.Confetti.Each(ecs.World, func(e *donburi.Entry) {
			UpdateConfetti(ecs, e)
		})

		confettis := donburi.NewQuery(filter.Contains(tags.Confetti)).Count(ecs.World)
		if confettis < 360 {
			factory.CreateConfetti(ecs)
		}
		p := components.StageCleared.Get(e)
		spr := components.AnimatedSprite.Get(e)

		opacity, _ := p.BgTween.Update(1.0 / 60.0)
		p.BgOpacity = float64(opacity)

		offset, _ := p.OffsetTween.Update(1.0 / 60.0)
		p.OffsetY = float64(offset)

		if p.IsRequestResume() {
			audio.StopBGM()
			factory.CreateTransition(ecs, true, func() {
				events.WinLevelEvents.Publish(ecs.World, struct{}{})
			})
		}

		if p.IsRequestRestart() {
			audio.StopBGM()
			factory.CreateTransition(ecs, true, func() {
				events.RestartLevelEvents.Publish(ecs.World, events.RestartLevelEvent{})
			})
		}

		if p.IsRequestExit() {
			audio.StopBGM()
			factory.CreateTransition(ecs, true, func() {
				events.ExitLevelEvents.Publish(ecs.World, struct{}{})
			})
		}

		if input.Handler.ActionIsJustPressed(input.ActionRight) {
			p.Selected++
		} else if input.Handler.ActionIsJustPressed(input.ActionLeft) {
			p.Selected--
		}

		p.Selected = components.StageClearedAction(math.Max(float64(components.StageClearedExit), math.Min(float64(components.StageClearedRestart), float64(p.Selected))))

		switch p.Selected {
		case components.StageClearedExit:
			spr.X = Lerp(spr.X, float64(components.STAGE_CLEARED_UI_EXIT_X), 0.2)
			spr.Y = float64(components.STAGE_CLEARED_UI_EXIT_Y)
		case components.StageClearedResume:
			spr.X = Lerp(spr.X, float64(components.STAGE_CLEARED_UI_RESUME_X), 0.2)
			spr.Y = float64(components.STAGE_CLEARED_UI_RESUME_Y)
		case components.StageClearedRestart:
			spr.X = Lerp(spr.X, float64(components.STAGE_CLEARED_UI_RESTART_X), 0.2)
			spr.Y = float64(components.STAGE_CLEARED_UI_RESTART_Y)
		}

		spr.OffsetX = -2
		spr.OffsetY = -2 - p.OffsetY
	})
}

func DrawStageCleared(ecs *ecs.ECS, screen *ebiten.Image) {
	components.StageCleared.Each(ecs.World, func(e *donburi.Entry) {
		p := components.StageCleared.Get(e)
		opt := &ebiten.DrawImageOptions{}
		opt.ColorScale.ScaleAlpha(float32(p.BgOpacity) / 100)
		screen.DrawImage(components.StageClearedBgImage, opt)

		esystems.DrawAnimatedSprites(ecs, tags.Confetti, screen)

		ropt := &ebiten.DrawImageOptions{}
		ropt.GeoM.Translate(float64(components.STAGE_CLEARED_UI_RESTART_X), float64(components.STAGE_CLEARED_UI_RESTART_Y)-p.OffsetY)
		screen.DrawImage(components.StageClearedImage[components.StageClearedResetImage], ropt)

		rropt := &ebiten.DrawImageOptions{}
		rropt.GeoM.Translate(float64(components.STAGE_CLEARED_UI_RESUME_X), float64(components.STAGE_CLEARED_UI_RESUME_Y)-p.OffsetY)
		screen.DrawImage(components.StageClearedImage[components.StageClearedResumeImage], rropt)

		eopt := &ebiten.DrawImageOptions{}
		eopt.GeoM.Translate(float64(components.STAGE_CLEARED_UI_EXIT_X), float64(components.STAGE_CLEARED_UI_EXIT_Y)-p.OffsetY)
		screen.DrawImage(components.StageClearedImage[components.StageClearedExitImage], eopt)

		text.DrawShadowedText(screen, "Stage Cleared!", float64(config.Width)/2, float64(components.STAGE_CLEARED_UI_GAP)-p.OffsetY, true)
	})

	if platform.Platform() != platform.Mobile {
		esystems.DrawAnimatedSprites(ecs, components.StageCleared, screen)
	}
}
