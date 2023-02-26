package components

import (
	"layla/pkg/maps"

	"github.com/yohamta/donburi"
)

type MapRendererData struct {
	Renderer *maps.EbitenRenderer
	Centered bool
}

var MapRenderer = donburi.NewComponentType[MapRendererData]()
