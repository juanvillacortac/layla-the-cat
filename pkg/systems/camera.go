package systems

import (
	"layla/pkg/components"
	"layla/pkg/config"
	esystems "layla/pkg/systems/entities"
	"math"
	"math/rand"

	"github.com/yohamta/donburi/ecs"
)

func UpdateCamera(ecs *ecs.ECS) {
	entry, ok := components.Camera.First(ecs.World)
	if !ok {
		return
	}
	camera := components.Camera.Get(entry)
	obj := components.GetObject(entry)

	if camera.ShakeMagnitude != 0 {
		min := int(-camera.ShakeMagnitude)
		max := int(camera.ShakeMagnitude)
		camera.ShakeValueX = float64(rand.Intn(max-min) + min)
		camera.ShakeValueY = float64(rand.Intn(max-min) + min)
	} else {
		camera.ShakeValueX = 0
		camera.ShakeValueY = 0
	}

	px, py := obj.X+obj.W/2, obj.Y+obj.H/2

	if entry.HasComponent(components.Player) {
		py -= obj.H / 2
		playerData := components.Player.Get(entry)
		if playerData.SpeedY >= esystems.GRAVITY {
			py += playerData.SpeedY * 4
		}
	}

	dx := px - float64(config.Width)/2
	dy := py - float64(config.Height)/2
	// if camera.Lerp {
	dx = Lerp(camera.X, dx, 0.10)
	dy = Lerp(camera.Y, dy, 0.10)
	// }

	if camera.W > config.Width {
		camera.X = math.Max(0, math.Min(dx, float64(camera.W-config.Width)))
	} else {
		camera.X = -(float64(config.Width-camera.W) - obj.W) / 2
	}
	camera.Y = math.Max(0, math.Min(dy, float64(camera.H-config.Height)))

	camera.X += camera.ShakeValueX
	camera.Y += camera.ShakeValueY
}

func Lerp(a, b, t float64) float64 {
	return (1-t)*a + t*b
}
