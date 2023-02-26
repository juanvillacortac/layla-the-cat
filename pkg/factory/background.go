package factory

import (
	"layla/pkg/assets"
	"layla/pkg/components"
	"layla/pkg/layers"
	"strings"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func AppendBackground(e *donburi.Entry, path string) {
	path = strings.Replace(path, "../", "", 1)
	img, _, err := ebitenutil.NewImageFromFileSystem(assets.BgFS, path)
	if err != nil {
		panic(err)
	}

	components.Background.SetValue(e, components.BackgroundData{
		Image: img,
	})
}

func CreateBackground(ecs *ecs.ECS, path string) *donburi.Entry {
	e := ecs.World.Entry(ecs.Create(
		layers.Background,
		components.Background,
	))

	AppendBackground(e, path)

	return e
}
