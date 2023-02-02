package systems

import (
	"fmt"
	"layla/pkg/components"
	"layla/pkg/config"
	"layla/pkg/events"
	"layla/pkg/factory"
	"layla/pkg/text"

	"github.com/hajimehoshi/ebiten/v2"
	etext "github.com/hajimehoshi/ebiten/v2/text"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func UpdatePauseScreen(ecs *ecs.ECS) {
	components.UpdatePauseScreenLayout()
	components.PauseScreen.Each(ecs.World, func(e *donburi.Entry) {
		p := components.PauseScreen.Get(e)

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

		f := text.LoadFont("ExpressionPro", 16)
		txt := fmt.Sprintf("Paused")

		b := etext.BoundString(*f, "Paused")

		stopt := &ebiten.DrawImageOptions{}
		stopt.GeoM.Translate(float64(config.Width)/2-float64(b.Dx())/2, float64(components.PAUSE_UI_GAP)+float64(b.Dy())-p.OffsetY)
		stopt.GeoM.Translate(1, 1)
		stopt.ColorScale.Scale(0, 0, 0, 1)
		stopt.ColorScale.ScaleAlpha(0.8)
		etext.DrawWithOptions(screen, txt, *f, stopt)
		topt := &ebiten.DrawImageOptions{}
		topt.GeoM.Translate(float64(config.Width)/2-float64(b.Dx())/2, float64(components.PAUSE_UI_GAP)+float64(b.Dy())-p.OffsetY)
		etext.DrawWithOptions(screen, txt, *f, topt)

		// ebitenutil.DebugPrint(screen, fmt.Sprintf("%.2f lalala", p.OffsetY))
	})
}
