package tags

import "github.com/yohamta/donburi"

var (
	CursorUi    = donburi.NewTag().SetName("CursorUi")
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
	Goal        = donburi.NewTag().SetName("Goal")
	Confetti    = donburi.NewTag().SetName("Confetti")
)
