package esystems

import (
	"layla/pkg/components"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/ecs"
)

func DrawEnemy(ecs *ecs.ECS, screen *ebiten.Image) {
	DrawAnimatedSprites(ecs, components.AnimatedSprite, screen)
}
