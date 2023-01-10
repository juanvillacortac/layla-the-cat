package components

import (
	"layla/pkg/maps"

	"github.com/solarlune/ldtkgo"
	"github.com/yohamta/donburi"
)

type LevelData struct {
	LdtkProject *ldtkgo.Project
	Renderer    *maps.EbitenRenderer
}

var Level = donburi.NewComponentType[LevelData]()
