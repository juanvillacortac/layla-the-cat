package systems

import (
	"fmt"
	"layla/pkg/assets"
	"layla/pkg/components"
	"layla/pkg/config"
	"layla/pkg/events"
	"layla/pkg/input"
	"layla/pkg/platform"
	"layla/pkg/text"

	// "layla/pkg/input"
	// "layla/pkg/platform"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func UpdateInput(ecs *ecs.ECS) {
	components.UpdateInputLayout()

	components.Input.Each(ecs.World, func(e *donburi.Entry) {
		input := components.Input.Get(e)
		input.GamepadIDs = ebiten.AppendGamepadIDs(input.GamepadIDs[:0])
		input.TouchIDs = inpututil.AppendJustPressedTouchIDs(input.TouchIDs[:0])
		if len(input.TouchIDs) > 0 {
			config.C.Touch = true
		}

		if input.IsRequestPause() {
			events.PauseLevelEvents.Publish(ecs.World, events.PauseLevelEvent{})
		}
	})
}

func DrawInput(ecs *ecs.ECS, screen *ebiten.Image) {
	components.Input.Each(ecs.World, func(e *donburi.Entry) {
		if e.HasComponent(components.Player) {
			p := components.Player.Get(e)
			tOpt := &ebiten.DrawImageOptions{}
			tOpt.GeoM.Translate(float64(components.INPUT_UI_GAP), float64(components.INPUT_UI_GAP))
			screen.DrawImage(assets.ClockSprite, tOpt)
			text.DrawShadowedText(screen, fmt.Sprint(p.Time), float64(components.INPUT_UI_GAP)+8+6, float64(components.INPUT_UI_GAP)+1, false)
		}

		if !config.C.Touch || platform.Platform() == platform.Desktop || input.Handler.GamepadConnected() {
			return
		}

		input := components.Input.Get(e)

		s := config.C.InputScale

		pOpt := &ebiten.DrawImageOptions{}
		pOpt.GeoM.Scale(s, s)
		pOpt.GeoM.Translate(float64(components.INPUT_UI_PAUSE_X), float64(components.INPUT_UI_PAUSE_Y))
		screen.DrawImage(components.InputUiPause, pOpt)

		player := components.Player.Get(e)
		if player != nil && player.Die {
			return
		}

		lOpt := &ebiten.DrawImageOptions{}
		lOpt.GeoM.Scale(s, s)
		lOpt.GeoM.Translate(float64(components.INPUT_UI_LEFT_X), float64(components.INPUT_UI_LEFT_Y))
		lOpt.ColorScale.ScaleAlpha(0.4)
		screen.DrawImage(components.InputImage[components.InputImageMovement][input.IsRunning(components.InputDirectionLeft)], lOpt)

		rOpt := &ebiten.DrawImageOptions{}
		rOpt.GeoM.Scale(-s, s)
		rOpt.GeoM.Translate(float64(components.INPUT_UI_RIGHT_X), float64(components.INPUT_UI_RIGHT_Y))
		rOpt.ColorScale.ScaleAlpha(0.4)
		rOpt.GeoM.Translate(float64(components.INPUT_UI_CELL_SIZE), 0)
		screen.DrawImage(components.InputImage[components.InputImageMovement][input.IsRunning(components.InputDirectionRight)], rOpt)

		jOpt := &ebiten.DrawImageOptions{}
		jOpt.GeoM.Scale(s, s)
		jOpt.GeoM.Translate(float64(components.INPUT_UI_JUMP_X), float64(components.INPUT_UI_JUMP_Y))
		jOpt.ColorScale.ScaleAlpha(0.4)
		screen.DrawImage(components.InputImage[components.InputImageJump][input.IsLongJumping()], jOpt)

		if input.Sliding {
			rrOpt := &ebiten.DrawImageOptions{}
			rrOpt.GeoM.Scale(s, s)
			rrOpt.GeoM.Translate(float64(components.INPUT_UI_RELEASE_X), float64(components.INPUT_UI_RELEASE_Y))
			rrOpt.ColorScale.ScaleAlpha(0.4)
			screen.DrawImage(components.InputImage[components.InputImageRelease][false], rrOpt)
		}
	})
}
