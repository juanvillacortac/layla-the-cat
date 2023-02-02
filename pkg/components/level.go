package components

import (
	"layla/pkg/maps"

	"github.com/yohamta/donburi"
)

type LevelData struct {
	LdtkProject *maps.LevelCtx
	Renderer    *maps.EbitenRenderer
}

var Level = donburi.NewComponentType[LevelData]()
