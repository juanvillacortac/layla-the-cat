package esystems

import (
	"layla/pkg/components"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/ganim8/v2"
)

func UpdatePlayerCorpse(ecs *ecs.ECS, e *donburi.Entry) {
	corpse := components.PlayerCorpse.Get(e)
	if corpse.Fall {
		corpse.SpeedY += GRAVITY
		corpse.Y += corpse.SpeedY
	}
}

func DrawPlayerCorpse(ecs *ecs.ECS, e *donburi.Entry, screen *ebiten.Image) {
	camera := components.GetCamera(ecs)
	corpse := components.PlayerCorpse.Get(e)
	x, y := corpse.X-math.Round(camera.X)-1, corpse.Y-math.Round(camera.Y)-1

	var opts *ganim8.DrawOptions
	opts = ganim8.DrawOpts(x, y)
	components.PlayerAnimations[components.PlayerDie].Draw(screen, opts)
}
