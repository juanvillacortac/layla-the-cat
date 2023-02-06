package systems

import (
	"layla/pkg/assets"
	"layla/pkg/components"
	"layla/pkg/config"
	"layla/pkg/events"
	"layla/pkg/factory"
	"layla/pkg/platform"
	"time"

	"layla/pkg/text"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func UpdateTitleScreen(ecs *ecs.ECS) {
	components.TitleScreen.Each(ecs.World, func(e *donburi.Entry) {
		ts := components.TitleScreen.Get(e)
		timers := components.TimerSystem.Get(e)

		offset, _, done := ts.OffsetTween.Update(1.0 / 60.0)
		if done && ts.TextTimer == nil {
			ts.TextTimer = timers.After(time.Millisecond*500, func() {
				ts.ShowText = !ts.ShowText
			})
			ts.TextTimer.Loop = true
			ts.ShowText = true
		}
		ts.OffsetY = float64(offset)

		ts.TouchIDs = inpututil.AppendJustPressedTouchIDs(ts.TouchIDs[:0])

		if inpututil.IsKeyJustPressed(ebiten.KeySpace) || len(ts.TouchIDs) > 0 {
			factory.CreateTransition(ecs, true, func() {
				events.LoadLevelEvents.Publish(ecs.World, "grass")
			})
		}
	})
}

func DrawTitleScreen(ecs *ecs.ECS, screen *ebiten.Image) {
	components.TitleScreen.Each(ecs.World, func(e *donburi.Entry) {
		ts := components.TitleScreen.Get(e)
		w, _ := assets.LogoSprite.Size()
		logoOpt := &ebiten.DrawImageOptions{}
		logoOpt.GeoM.Translate(float64(config.Width/2-w/2), 8-ts.OffsetY)
		screen.DrawImage(assets.LogoSprite, logoOpt)

		if ts.ShowText {
			txt := "Press SPACE to start"
			if platform.Platform() == platform.Mobile {
				txt = "Tap to start"
			}

			text.DrawShadowedText(screen, txt, float64(config.Width)/2, float64(config.Height)-float64(components.PAUSE_UI_GAP)-4, true)
		}
		//
		// f := text.LoadFont("min", 16)
		// txt := fmt.Sprintf("Juan Villacorta - %v", time.Now().Year())
		//
		// b := etext.BoundString(*f, txt)
		// x, y := float64(config.Width)/2-float64(b.Dx())/2, float64(config.Height)-8
		// stopt := &ebiten.DrawImageOptions{}
		// stopt.GeoM.Translate(x, y)
		// stopt.GeoM.Translate(1, 1)
		// stopt.ColorScale.Scale(0, 0, 0, 1)
		// stopt.ColorScale.ScaleAlpha(0.8)
		// etext.DrawWithOptions(screen, txt, *f, stopt)
		// topt := &ebiten.DrawImageOptions{}
		// topt.GeoM.Translate(x, y)
		// etext.DrawWithOptions(screen, txt, *f, topt)
	})
}
