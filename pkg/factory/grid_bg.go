package factory

import (
	"bytes"
	"image"
	"layla/pkg/archetypes"
	"layla/pkg/assets"
	"layla/pkg/components"
	"layla/pkg/config"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/ganim8/v2"
)

var gridBgSprite *ebiten.Image
var gridAnimGrid *ganim8.Grid

func init() {
	img, _, err := image.Decode(bytes.NewReader(assets.GridPng))
	if err != nil {
		panic(err)
	}
	gridBgSprite = ebiten.NewImageFromImage(img)
	gridAnimGrid = ganim8.NewGrid(32, 32, 32*16, 32)
}

func CreateGridBg(ecs *ecs.ECS) *donburi.Entry {
	anim := ganim8.New(gridBgSprite, gridAnimGrid.Frames("1-16", 1), 80*time.Millisecond)

	bg := archetypes.NewGridBackground(ecs)

	components.AnimatedTilesGroup.Set(bg, components.NewAnimatedTilesGroup())

	qx, qy := math.Ceil(float64(config.Width*2)/32), math.Ceil(float64(config.Height*2)/32)
	for y := 0.0; y <= qy; y++ {
		for x := 0.0; x <= qx; x++ {
			tileEntry := archetypes.NewAnimTile(ecs)
			components.AnimatedTile.Set(tileEntry, &components.AnimatedTileData{
				X: x * 32, Y: y * 32,
				Anim: anim.Clone(),
			})
			components.AddToAnimTilesGroup(bg, tileEntry)
		}
	}

	return bg
}
