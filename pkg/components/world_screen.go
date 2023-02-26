package components

import (
	"layla/pkg/touch"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type WorldScreenData struct {
	Selected int
	Stroke   *touch.Stroke
	TouchIDs []ebiten.TouchID
}

var WorldScreen = donburi.NewComponentType[WorldScreenData]()
