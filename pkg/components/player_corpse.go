package components

import "github.com/yohamta/donburi"

type PlayerCorpseData struct {
	SpeedY float64
	Y      float64
	X      float64
	Fall   bool
}

var PlayerCorpse = donburi.NewComponentType[PlayerCorpseData]()
