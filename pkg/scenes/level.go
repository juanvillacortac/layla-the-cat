package scenes

import (
	"layla/pkg/factory"
	"layla/pkg/layers"
	"layla/pkg/systems"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type LevelScene struct {
	ecs *ecs.ECS
}

func NewLevelScene() *LevelScene {
	ecs := ecs.NewECS(donburi.NewWorld())

	ecs.AddSystem(systems.UpdatePlayer)
	ecs.AddSystem(systems.UpdateParticles)
	ecs.AddSystem(systems.UpdateCamera)

	ecs.AddSystem(systems.UpdateInput)

	ecs.AddRenderer(layers.Default, systems.DrawWall)
	ecs.AddRenderer(layers.Default, systems.DrawLevel)
	ecs.AddRenderer(layers.Default, systems.DrawParticles)
	ecs.AddRenderer(layers.Default, systems.DrawPlayer)
	ecs.AddRenderer(layers.Default, systems.DrawCamera)
	ecs.AddRenderer(layers.Default, systems.DrawCamera)

	ecs.AddRenderer(layers.Input, systems.DrawInput)

	ls := &LevelScene{ecs: ecs}

	factory.CreateLevel(ls.ecs, "another")

	return ls
}

func (ps *LevelScene) Update() {
	ps.ecs.Update()
}

func (ps *LevelScene) Draw(screen *ebiten.Image) {
	ps.ecs.Draw(screen)
}

func (ps *LevelScene) init() {

	// Define the world's Space. Here, a Space is essentially a grid (the game's width and height, or 640x360), made up of 16x16 cells. Each cell can have 0 or more Objects within it,
	// and collisions can be found by checking the Space to see if the Cells at specific positions contain (or would contain) Objects. This is a broad, simplified approach to collision
	// detection.
	// space := factory.CreateSpace(ps.ecs)

	// dresolv.Add(space,
	// 	// Construct the solid level geometry. Note that the simple approach of checking cells in a Space for collision works simply when the geometry is aligned with the cells,
	// 	// as it all is in this platformer example.
	// 	factory.CreateWall(ps.ecs, resolv.NewObject(0, 0, 16, gh, "solid")),
	// 	factory.CreateWall(ps.ecs, resolv.NewObject(gw-16, 0, 16, gh, "solid")),
	// 	factory.CreateWall(ps.ecs, resolv.NewObject(0, 0, gw, 16, "solid")),
	// 	factory.CreateWall(ps.ecs, resolv.NewObject(0, gh-24, gw, 32, "solid")),
	// 	factory.CreateWall(ps.ecs, resolv.NewObject(160, gh-56, 160, 32, "solid")),
	// 	factory.CreateWall(ps.ecs, resolv.NewObject(320, 64, 32, 160, "solid")),
	// 	factory.CreateWall(ps.ecs, resolv.NewObject(64, 128, 16, 160, "solid")),
	// 	factory.CreateWall(ps.ecs, resolv.NewObject(gw-128, 64, 128, 16, "solid")),
	// 	factory.CreateWall(ps.ecs, resolv.NewObject(gw-128, gh-88, 128, 16, "solid")),
	//
	// 	factory.CreatePlayer(ps.ecs),
	// )
}
