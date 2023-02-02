package components

import (
	"github.com/yohamta/donburi"
)

type CollectableData struct {
	OffsetY float64
}

var Collectable = donburi.NewComponentType[CollectableData]()
