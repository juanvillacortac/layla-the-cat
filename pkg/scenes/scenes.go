package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/ecs"
)

type Scene interface {
	Ecs() *ecs.ECS
	Update()
	Draw(screen *ebiten.Image)
}
