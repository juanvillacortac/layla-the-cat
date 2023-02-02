package layers

import "github.com/yohamta/donburi/ecs"

const (
	Background ecs.LayerID = iota
	Default    ecs.LayerID = iota
	Input      ecs.LayerID = iota
	Transition ecs.LayerID = iota
)
