package components

import (
	"github.com/yohamta/donburi"
)

type EntityLayer int

const (
	EntityBackLayer EntityLayer = iota
	EntityFrontLayer
	EntityBackgroundLayer
)

type EntityData struct {
	Identifier string
	Layer      EntityLayer
}

var Entity = donburi.NewComponentType[EntityData]()
