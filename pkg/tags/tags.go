package tags

import "github.com/yohamta/donburi"

var (
	Level     = donburi.NewTag().SetName("Level")
	Wall      = donburi.NewTag().SetName("Wall")
	Player    = donburi.NewTag().SetName("Player")
	Camera    = donburi.NewTag().SetName("Camera")
	Particles = donburi.NewTag().SetName("Particles")
)
