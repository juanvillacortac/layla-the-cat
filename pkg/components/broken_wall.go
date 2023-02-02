package components

import "github.com/yohamta/donburi"

type BrokenWallData struct {
	OffsetX  float64
	Collided bool
}

var BrokenWall = donburi.NewComponentType[BrokenWallData]()
