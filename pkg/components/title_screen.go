package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/ebitick"
	"github.com/tanema/gween"
	"github.com/yohamta/donburi"
)

type TitleScreenData struct {
	BgOpacity   float64
	OffsetY     float64
	BgTween     *gween.Tween
	OffsetTween *gween.Sequence
	TextTimer   *ebitick.Timer
	ShowText    bool

	TouchIDs []ebiten.TouchID
}

var TitleScreen = donburi.NewComponentType[TitleScreenData]()
