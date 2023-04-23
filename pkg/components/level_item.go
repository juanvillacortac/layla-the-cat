package components

import (
	"layla/pkg/touch"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type LevelItemPressedAction int

const (
	LevelItemNotPressed LevelItemPressedAction = iota
	LevelItemKeyPressed
	LevelItemTouchPressed
)

type LevelItemData struct {
	Map      string
	Number   int
	TouchIDs []ebiten.TouchID
	Stroke   *touch.Stroke
	Pressed  LevelItemPressedAction
}

var LevelItem = donburi.NewComponentType[LevelItemData]()
