package layers

import "github.com/yohamta/donburi/ecs"

const (
	Background ecs.LayerID = iota
	Default
	Foreground
	Input
	Transition
)
