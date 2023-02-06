package systems

import (
	"layla/pkg/components"
	"layla/pkg/config"
	"layla/pkg/events"
	"layla/pkg/factory"
	"layla/pkg/input"
	"layla/pkg/platform"
	esystems "layla/pkg/systems/entities"
	"layla/pkg/text"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	// etext "github.com/hajimehoshi/ebiten/v2/text"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func UpdatePauseScreen(ecs *ecs.ECS) {
	components.UpdatePauseScreenLayout()
	components.PauseScreen.Each(ecs.World, func(e *donburi.Entry) {
		p := components.PauseScreen.Get(e)
		spr := components.AnimatedSprite.Get(e)

		opacity, _ := p.BgTween.Update(1.0 / 60.0)
		p.BgOpacity = float64(opacity)

		offset, _ := p.OffsetTween.Update(1.0 / 60.0)
		p.OffsetY = float64(offset)

		if p.IsRequestResume() {
			events.PauseLevelEvents.Publish(ecs.World, events.PauseLevelEvent{})
		}

		if p.IsRequestRestart() {
			factory.CreateTransition(ecs, true, func() {
				events.RestartLevelEvents.Publish(ecs.World, events.RestartLevelEvent{})
			})
		}

		if p.IsRequestExit() {
			factory.CreateTransition(ecs, true, func() {
				events.ExitLevelEvents.Publish(ecs.World, struct{}{})
			})
		}

		if input.Handler.ActionIsJustPressed(input.ActionRight) {
			p.Selected++
		} else if input.Handler.ActionIsJustPressed(input.ActionLeft) {
			p.Selected--
		}

		p.Selected = components.PauseScreenAction(math.Max(float64(components.PauseScreenExit), math.Min(float64(components.PauseScreenRestart), float64(p.Selected))))

		switch p.Selected {
		case components.PauseScreenExit:
			spr.X = Lerp(spr.X, float64(components.PAUSE_UI_EXIT_X), 0.2)
			spr.Y = float64(components.PAUSE_UI_EXIT_Y)
		case components.PauseScreenResume:
			spr.X = Lerp(spr.X, float64(components.PAUSE_UI_RESUME_X), 0.2)
			spr.Y = float64(components.PAUSE_UI_RESUME_Y)
		case components.PauseScreenRestart:
			spr.X = Lerp(spr.X, float64(components.PAUSE_UI_RESTART_X), 0.2)
			spr.Y = float64(components.PAUSE_UI_RESTART_Y)
		}

		spr.OffsetX = -2
		spr.OffsetY = -2 - p.OffsetY
	})
}

func DrawPauseScreen(ecs *ecs.ECS, screen *ebiten.Image) {
	components.PauseScreen.Each(ecs.World, func(e *donburi.Entry) {
		p := components.PauseScreen.Get(e)
		opt := &ebiten.DrawImageOptions{}
		opt.ColorScale.ScaleAlpha(float32(p.BgOpacity) / 100)
		screen.DrawImage(components.PauseScreenBgImage, opt)

		ropt := &ebiten.DrawImageOptions{}
		ropt.GeoM.Translate(float64(components.PAUSE_UI_RESTART_X), float64(components.PAUSE_UI_RESTART_Y)-p.OffsetY)
		screen.DrawImage(components.PauseScreenImage[components.PauseScreenResetImage], ropt)

		rropt := &ebiten.DrawImageOptions{}
		rropt.GeoM.Translate(float64(components.PAUSE_UI_RESUME_X), float64(components.PAUSE_UI_RESUME_Y)-p.OffsetY)
		screen.DrawImage(components.PauseScreenImage[components.PauseScreenResumeImage], rropt)

		eopt := &ebiten.DrawImageOptions{}
		eopt.GeoM.Translate(float64(components.PAUSE_UI_EXIT_X), float64(components.PAUSE_UI_EXIT_Y)-p.OffsetY)
		screen.DrawImage(components.PauseScreenImage[components.PauseScreenExitImage], eopt)

		text.DrawShadowedText(screen, "Paused", float64(config.Width)/2, float64(components.PAUSE_UI_GAP)-p.OffsetY, true)
		//
		// f := text.LoadFont("ExpressionPro", 16)
		// txt := fmt.Sprintf("Paused")
		//
		// b := etext.BoundString(*f, "Paused")
		//
		// stopt := &ebiten.DrawImageOptions{}
		// stopt.GeoM.Translate(float64(config.Width)/2-float64(b.Dx())/2, float64(components.PAUSE_UI_GAP)+float64(b.Dy())-p.OffsetY)
		// stopt.GeoM.Translate(1, 1)
		// stopt.ColorScale.Scale(0, 0, 0, 1)
		// stopt.ColorScale.ScaleAlpha(0.8)
		// etext.DrawWithOptions(screen, txt, *f, stopt)
		// topt := &ebiten.DrawImageOptions{}
		// topt.GeoM.Translate(float64(config.Width)/2-float64(b.Dx())/2, float64(components.PAUSE_UI_GAP)+float64(b.Dy())-p.OffsetY)
		// etext.DrawWithOptions(screen, txt, *f, topt)

		// ebitenutil.DebugPrint(screen, fmt.Sprintf("%.2f lalala", p.OffsetY))
	})

	if platform.Platform() != platform.Mobile {
		esystems.DrawAnimatedSprites(ecs, components.PauseScreen, screen)
	}
}
