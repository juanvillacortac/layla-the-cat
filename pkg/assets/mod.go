package assets

import (
	"embed"
	"fmt"

	"github.com/solarlune/ldtkgo"
)

//go:embed maps/*.ldtk
var MapsFS embed.FS

//go:embed tilesets/*.png
var TilesetsFS embed.FS

//go:embed player.png
var PlayerPng []byte

//go:embed particles.png
var ParticlesPng []byte

//go:embed input.png
var InputPng []byte

func LoadMap(name string) (*ldtkgo.Project, error) {
	b, err := MapsFS.ReadFile(fmt.Sprintf("maps/%s.ldtk", name))
	if err != nil {
		return nil, err
	}
	return ldtkgo.Read(b)
}
