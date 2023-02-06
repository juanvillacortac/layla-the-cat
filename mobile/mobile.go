package layla

import (
	"github.com/hajimehoshi/ebiten/v2/mobile"

	"layla/pkg/config"
	"layla/pkg/game"
	"layla/pkg/input"
)

func init() {
	game := game.NewGame()
	mobile.SetGame(game)
}

func SetDataPath(path string) {
	config.DataDir = path
}

func BackButton() {
	input.Handler.EmitEvent(input.SimulatedExit)
}
