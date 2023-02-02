package layla

import (
	"github.com/hajimehoshi/ebiten/v2/mobile"

	"layla/pkg/config"
	"layla/pkg/game"
)

func init() {
	game := game.NewGame()
	mobile.SetGame(game)
}

func SetDataPath(path string) {
	config.DataDir = path
}
func BackButton() {}
