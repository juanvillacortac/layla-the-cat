package systems

import (
	"layla/pkg/assets"
	"layla/pkg/platform"
	"layla/pkg/tags"
	"layla/pkg/touch"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func DrawCursorUi(ecs *ecs.ECS, screen *ebiten.Image) {
	tags.CursorUi.Each(ecs.World, func(e *donburi.Entry) {
		if platform.Platform() == platform.Mobile {
			return
		}
		x, y := touch.MousePos()
		opt := &ebiten.DrawImageOptions{}
		opt.GeoM.Translate(float64(x), float64(y))

		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			opt.GeoM.Translate(1, 1)
		}

		sopt := &ebiten.DrawImageOptions{}
		sopt.GeoM.Translate(float64(x)+1, float64(y)+1)
		sopt.ColorScale.Scale(0, 0, 0, 1)
		sopt.ColorScale.ScaleAlpha(0.4)

		screen.DrawImage(assets.CursorUiSprite, sopt)
		screen.DrawImage(assets.CursorUiSprite, opt)
	})
}
