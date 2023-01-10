package systems

import (
	"layla/pkg/components"
	"layla/pkg/tags"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func DrawLevel(ecs *ecs.ECS, screen *ebiten.Image) {
	tags.Level.Each(ecs.World, func(e *donburi.Entry) {
		level := components.Level.Get(e)
		for _, layer := range level.Renderer.RenderedLayers {
			camera := components.GetCamera(ecs)
			opt := &ebiten.DrawImageOptions{}
			opt.GeoM.Translate(-camera.X, -camera.Y)
			screen.DrawImage(layer.Image, opt)
		}
	})
}
