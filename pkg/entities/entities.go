package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/ecs"
)

type EntityType string
type EntityRenderer func(ecs *ecs.ECS, screen *ebiten.Image)

const (
	Goal         EntityType = "Goal"
	Player       EntityType = "Player"
	PlayerCorpse EntityType = "PlayerCorpse"
	Collectable  EntityType = "Collectable"

	Saw   EntityType = "Saw"
	Pinch EntityType = "Pinch"

	BrokenWall EntityType = "BrokenWall"

	LevelItem EntityType = "Level"
)
