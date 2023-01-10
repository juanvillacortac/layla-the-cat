package systems

import (
	"layla/pkg/components"
	"layla/pkg/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func UpdateInput(ecs *ecs.ECS) {
	components.Input.Each(ecs.World, func(e *donburi.Entry) {
		input := components.Input.Get(e)
		input.TouchIDs = inpututil.AppendJustPressedTouchIDs(input.TouchIDs[:0])
		if len(input.TouchIDs) > 0 {
			config.C.Touch = true
		}
		// for _, id := range input.TouchIDs {
		// 	_, y := ebiten.TouchPosition(id)
		// 	if y >= components.INPUT_UI_CELL_SIZE {
		// 		ebiten.SetFullscreen(true)
		// 	}
		// }
	})
}

func DrawInput(ecs *ecs.ECS, screen *ebiten.Image) {
	if !config.C.Touch {
		return
	}
	components.Input.Each(ecs.World, func(e *donburi.Entry) {
		input := components.Input.Get(e)

		lOpt := &ebiten.DrawImageOptions{}
		lOpt.GeoM.Scale(2, 2)
		lOpt.GeoM.Translate(float64(components.INPUT_UI_LEFT_X), float64(components.INPUT_UI_LEFT_Y))
		lOpt.ColorM.Scale(1, 1, 1, 0.4)
		screen.DrawImage(components.InputImage[components.InputImageMovement][input.IsRunning(components.InputDirectionLeft)], lOpt)

		rOpt := &ebiten.DrawImageOptions{}
		rOpt.GeoM.Scale(-2, 2)
		rOpt.GeoM.Translate(float64(components.INPUT_UI_RIGHT_X)+float64(components.INPUT_UI_CELL_SIZE), float64(components.INPUT_UI_RIGHT_Y))
		rOpt.ColorM.Scale(1, 1, 1, 0.4)
		rOpt.GeoM.Translate(float64(components.INPUT_UI_CELL_SIZE), 0)
		screen.DrawImage(components.InputImage[components.InputImageMovement][input.IsRunning(components.InputDirectionRight)], rOpt)

		jOpt := &ebiten.DrawImageOptions{}
		jOpt.GeoM.Scale(2, 2)
		jOpt.GeoM.Translate(float64(components.INPUT_UI_JUMP_X), float64(components.INPUT_UI_JUMP_Y))
		jOpt.ColorM.Scale(1, 1, 1, 0.4)
		screen.DrawImage(components.InputImage[components.InputImageJump][input.IsLongJumping()], jOpt)
	})
}
