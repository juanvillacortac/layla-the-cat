package touch

import (
	"layla/pkg/config"

	"github.com/hajimehoshi/ebiten/v2"
)

func TouchPos(touchId ebiten.TouchID) (int, int) {
	x, y := ebiten.TouchPosition(touchId)
	switch config.C.CrtQuality {
	case config.CrtQualityHigh:
		return x / int(config.Scale), y / int(config.Scale)
	default:
		return x, y
	}
}
