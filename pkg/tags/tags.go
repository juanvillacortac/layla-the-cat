package tags

import "github.com/yohamta/donburi"

var (
	Level       = donburi.NewTag().SetName("Level")
	Background  = donburi.NewTag().SetName("Background")
	Wall        = donburi.NewTag().SetName("Wall")
	BrokenWall  = donburi.NewTag().SetName("BrokenWall")
	Collectable = donburi.NewTag().SetName("Collectable")
	Player      = donburi.NewTag().SetName("Player")
	Enemy       = donburi.NewTag().SetName("Enemy")
	Pushable    = donburi.NewTag().SetName("Pushable")
	Camera      = donburi.NewTag().SetName("Camera")
	Particles   = donburi.NewTag().SetName("Particles")
	Flash       = donburi.NewTag().SetName("Flash")
)
