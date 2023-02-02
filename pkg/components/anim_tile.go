package components

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/ganim8/v2"
)

type AnimatedTileData struct {
	X, Y  float64
	Anim  *ganim8.Animation
	Layer ecs.LayerID
}

type AnimatedTilesGroupData struct {
	tiles []*AnimatedTileData
}

func NewAnimatedTilesGroup() *AnimatedTilesGroupData {
	tiles := make([]*AnimatedTileData, 0)
	group := &AnimatedTilesGroupData{
		tiles: tiles,
	}
	return group
}

func (g *AnimatedTilesGroupData) Add(tile *AnimatedTileData) {
	g.tiles = append(g.tiles, tile)
}

var AnimatedTile = donburi.NewComponentType[AnimatedTileData]()
var AnimatedTilesGroup = donburi.NewComponentType[AnimatedTilesGroupData]()

func AddToAnimTilesGroup(group *donburi.Entry, objects ...*donburi.Entry) {
	for _, obj := range objects {
		AnimatedTilesGroup.Get(group).Add(AnimatedTile.Get(obj))
	}
}
