package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type BackgroundData struct {
	Image *ebiten.Image
}

var Background = donburi.NewComponentType[BackgroundData]()
