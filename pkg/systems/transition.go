package systems

import (
	"layla/pkg/components"
	"layla/pkg/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func UpdateTransitions(ecs *ecs.ECS) {
	components.Transition.Each(ecs.World, func(e *donburi.Entry) {
		t := components.Transition.Get(e)
		t.X += t.Speed
		w := config.Width
		if t.Outro {
			w += 64
		}
		if t.X > float64(w) {
			if t.EndAction != nil {
				(*t.EndAction)()
			}
			e.Remove()
		}
	})
}

func DrawTransitions(ecs *ecs.ECS, screen *ebiten.Image) {
	components.Transition.Each(ecs.World, func(e *donburi.Entry) {
		t := components.Transition.Get(e)

		t.RenderImage()

		opt := &ebiten.DrawImageOptions{}
		if t.Outro {
			opt.GeoM.Scale(-1, 1)
		}
		opt.GeoM.Translate(t.X, 0)
		screen.DrawImage(t.Image, opt)
	})
}
