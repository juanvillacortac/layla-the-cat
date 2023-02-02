package maps

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/ldtkgo"
)

type LevelCtx struct {
	Project  *ldtkgo.Project
	Tilesets map[int]*ebiten.Image
	Level    *ldtkgo.Level
}

func NewLevelCtx(project *ldtkgo.Project) *LevelCtx {
	tilesetLoader := NewDiskLoader()
	tilesets := map[int]*ebiten.Image{}
	for _, tileset := range project.Tilesets {
		tilesets[tileset.ID] = tilesetLoader.LoadTileset(tileset.Path)
	}
	return &LevelCtx{
		Project:  project,
		Tilesets: tilesets,
		Level:    project.Levels[0],
	}
}
