package main

import (
	"layla/pkg/config"
	"layla/pkg/game"
	"layla/pkg/platform"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	if platform.Platform() == platform.Desktop {
		ebiten.SetWindowSize(config.Width*int(config.Scale), config.Height*int(config.Scale))
	}
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	ebiten.SetWindowTitle("Layla the Cat")
	if err := ebiten.RunGame(game.NewGame()); err != nil {
		log.Fatal(err)
	}
}
