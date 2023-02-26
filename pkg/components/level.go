package components

import (
	"layla/pkg/maps"

	"github.com/yohamta/donburi"
)

type LevelData struct {
	LdtkProject *maps.LevelCtx
	Number      int
	Deaths      int
}

var Level = donburi.NewComponentType[LevelData]()
