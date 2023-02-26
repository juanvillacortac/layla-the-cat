package components

import (
	"time"

	"github.com/solarlune/ebitick"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type CameraData struct {
	X              float64
	Y              float64
	W              int
	H              int
	ShakeMagnitude float64
	Shaking        bool
	ShakeTimer     *ebitick.Timer
	ShakeValueX    float64
	ShakeValueY    float64
	Lerp           bool
}

var Camera = donburi.NewComponentType[CameraData]()

func GetCamera(ecs *ecs.ECS) *CameraData {
	entry, ok := Camera.First(ecs.World)
	if !ok {
		return &CameraData{}
	}
	return Camera.Get(entry)
}

func ShakeCamera(ecs *ecs.ECS, magnitude float64, duration time.Duration) {
	entry, ok := Camera.First(ecs.World)
	if !ok {
		return
	}
	camera := Camera.Get(entry)
	ts := TimerSystem.Get(entry)
	ts.After(duration, func() {
		camera.ShakeMagnitude = 0
		camera.ShakeValueX = 0
		camera.ShakeValueY = 0
	})
	camera.ShakeMagnitude = magnitude
	camera.Shaking = true
}

func StartShakingCamera(ecs *ecs.ECS, magnitude float64) {
	entry, ok := Camera.First(ecs.World)
	if !ok {
		return
	}
	camera := Camera.Get(entry)
	camera.ShakeMagnitude = magnitude
	camera.Shaking = true
}

func StopShakingCamera(ecs *ecs.ECS) {
	entry, ok := Camera.First(ecs.World)
	if !ok {
		return
	}
	camera := Camera.Get(entry)
	camera.ShakeMagnitude = 0
	camera.ShakeValueX = 0
	camera.ShakeValueY = 0
	camera.Shaking = false
}
