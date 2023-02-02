package components

import (
	"bytes"
	"image"
	_ "image/png"
	"layla/pkg/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
)

var BoxImage *ebiten.Image

func init() {
	img, _, err := image.Decode(bytes.NewReader(assets.BoxPng))
	if err != nil {
		panic(err)
	}
	BoxImage = ebiten.NewImageFromImage(img)
}

type PushableData struct {
	SpeedX   float64
	SpeedY   float64
	OnGround *resolv.Object
	Tile     *ebiten.Image
}

var Pushable = donburi.NewComponentType[PushableData]()
