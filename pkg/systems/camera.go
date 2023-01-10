package systems

import (
	"layla/pkg/components"
	"layla/pkg/config"
	"layla/pkg/tags"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/ecs"
)

func UpdateCamera(ecs *ecs.ECS) {
	camera := components.GetCamera(ecs)
	if camera == nil {
		return
	}
	player := components.GetObject(tags.Player.MustFirst(ecs.World))

	dx := Lerp(camera.X, player.X-float64(config.C.Width)/2, 0.05)
	dy := Lerp(camera.Y, player.Y-float64(config.C.Height)/2, 0.05)

	if camera.W > config.C.Width {
		camera.X = math.Max(0, math.Min(dx, float64(camera.W-config.C.Width)))
	} else {
		camera.X = -float64(config.C.Width-camera.W) / 2
	}
	camera.Y = math.Max(0, math.Min(dy, float64(camera.H-config.C.Height)))
}

func DrawCamera(ecs *ecs.ECS, screen *ebiten.Image) {
}

func Lerp(a, b, t float64) float64 {
	return (1-t)*a + t*b
}
