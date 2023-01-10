package main

import (
	"layla/pkg/config"
	"layla/pkg/game"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(config.C.BaseWidth*2, config.C.BaseHeight*2)
	ebiten.SetWindowTitle("Layla the Cat")
	if err := ebiten.RunGame(game.NewGame()); err != nil {
		log.Fatal(err)
	}
}
