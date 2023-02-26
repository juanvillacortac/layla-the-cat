package components

import (
	"layla/pkg/touch"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type LevelItemData struct {
	Map      string
	Number   int
	TouchIDs []ebiten.TouchID
	Stroke   *touch.Stroke
}

var LevelItem = donburi.NewComponentType[LevelItemData]()
