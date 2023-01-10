package systems

import (
	"image/color"
	"layla/pkg/components"
	"layla/pkg/tags"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func DrawWall(ecs *ecs.ECS, screen *ebiten.Image) {
	tags.Wall.Each(ecs.World, func(e *donburi.Entry) {
		o := components.GetObject(e)
		drawColor := color.RGBA{60, 60, 60, 255}
		camera := components.GetCamera(ecs)
		ebitenutil.DrawRect(screen, o.X-camera.X, o.Y-camera.Y, o.W, o.H, drawColor)
	})
}
