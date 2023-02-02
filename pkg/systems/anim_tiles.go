package systems

import (
	"layla/pkg/components"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/ganim8/v2"
)

func UpdateAnimTiles(ecs *ecs.ECS) {
	components.AnimatedTile.Each(ecs.World, func(e *donburi.Entry) {
		tile := components.AnimatedTile.Get(e)
		tile.Anim.Update()
	})
}

func DrawAnimTiles(ecs *ecs.ECS, screen *ebiten.Image) {
	components.AnimatedTile.Each(ecs.World, func(e *donburi.Entry) {
		tile := components.AnimatedTile.Get(e)
		if camera := components.GetCamera(ecs); camera != nil {
			x, y := tile.X-math.Round(camera.ShakeValueX), tile.Y-math.Round(camera.ShakeValueY)
			tile.Anim.Draw(screen, ganim8.DrawOpts(x, y))
		} else {
			tile.Anim.Draw(screen, ganim8.DrawOpts(tile.X, tile.Y))
		}
	})
}
