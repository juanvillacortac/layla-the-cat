package systems

import (
	"image/color"
	"layla/pkg/tags"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func DrawFlash(ecs *ecs.ECS, screen *ebiten.Image) {
	tags.Flash.Each(ecs.World, func(e *donburi.Entry) {
		screen.Fill(color.White)
	})
}
