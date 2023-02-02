package components

import (
	"image/color"
	"layla/pkg/assets"
	"layla/pkg/config"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type TransitionData struct {
	X         float64
	Outro     bool
	Speed     float64
	EndAction *func()
	Image     *ebiten.Image
}

func (t *TransitionData) RenderImage() {
	w, h := config.Width, config.Height
	iw, ih := 0, 0
	if t.Image != nil {
		iw, ih = t.Image.Size()
	}
	if t.Image == nil || iw > w-64 || iw < w-64 || ih > h || ih < h {
		img := ebiten.NewImage(w+64, h)

		img2 := ebiten.NewImage(w, h)
		// img2.Fill(color.RGBA{24, 20, 37, 255})
		img2.Fill(color.Black)
		opt := &ebiten.DrawImageOptions{}
		opt.GeoM.Translate(64, 0)
		img.DrawImage(img2, opt)

		qy := math.Ceil(float64(config.Height) / 64)
		for y := 0.0; y <= qy; y++ {
			opt := &ebiten.DrawImageOptions{}
			opt.GeoM.Translate(0, y*64)
			img.DrawImage(assets.DitheringSprite, opt)
		}
		t.Image = img
	}
}

var Transition = donburi.NewComponentType[TransitionData]()
