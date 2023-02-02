package systems

import (
	"layla/pkg/components"
	"layla/pkg/config"
	"layla/pkg/tags"
	"math"
	"math/rand"

	"github.com/yohamta/donburi/ecs"
)

func UpdateCamera(ecs *ecs.ECS) {
	entry, ok := tags.Player.First(ecs.World)
	if !ok {
		return
	}
	camera := components.Camera.Get(entry)
	player := components.GetObject(entry)

	if camera.ShakeMagnitude != 0 {
		min := int(-camera.ShakeMagnitude)
		max := int(camera.ShakeMagnitude)
		camera.ShakeValueX = float64(rand.Intn(max-min) + min)
		camera.ShakeValueY = float64(rand.Intn(max-min) + min)
	} else {
		camera.ShakeValueX = 0
		camera.ShakeValueY = 0
	}

	px, py := player.X+player.W/2, player.Y

	dx := Lerp(camera.X, px-float64(config.Width)/2, 0.06)
	dy := Lerp(camera.Y, py-float64(config.Height)/2, 0.12)

	if camera.W > config.Width {
		camera.X = math.Max(0, math.Min(dx, float64(camera.W-config.Width)))
	} else {
		camera.X = -float64(config.Width-camera.W) / 2
	}
	camera.Y = math.Max(0, math.Min(dy, float64(camera.H-config.Height)))

	camera.X += camera.ShakeValueX
	camera.Y += camera.ShakeValueY
}

func Lerp(a, b, t float64) float64 {
	return (1-t)*a + t*b
}
