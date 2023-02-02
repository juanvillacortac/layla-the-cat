package assets

import (
	"bytes"
	"embed"
	"fmt"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/ldtkgo"
)

//go:embed maps/*.ldtk
var MapsFS embed.FS

//go:embed tilesets/*.png
var TilesetsFS embed.FS

//go:embed fonts/*.ttf
var FontsFS embed.FS

//go:embed player.png
var PlayerPng []byte

//go:embed particles.png
var ParticlesPng []byte

//go:embed box.png
var BoxPng []byte

//go:embed input.png
var InputPng []byte

//go:embed logo.png
var LogoPng []byte

//go:embed grid.png
var GridPng []byte

//go:embed dithering.png
var DitheringPng []byte

//go:embed collectable.png
var CollectablePng []byte

var LogoSprite *ebiten.Image
var DitheringSprite *ebiten.Image
var CollectableSprite *ebiten.Image

func init() {
	img, _, err := image.Decode(bytes.NewReader(LogoPng))
	if err != nil {
		panic(err)
	}
	LogoSprite = ebiten.NewImageFromImage(img)
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(DitheringPng))
	if err != nil {
		panic(err)
	}
	DitheringSprite = ebiten.NewImageFromImage(img)
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(CollectablePng))
	if err != nil {
		panic(err)
	}
	CollectableSprite = ebiten.NewImageFromImage(img)
}

//go:embed saw.png
var SawPng []byte
var SawSprite *ebiten.Image

//go:embed pinch.png
var PinchPng []byte
var PinchSprite *ebiten.Image

func init() {
	img, _, err := image.Decode(bytes.NewReader(SawPng))
	if err != nil {
		panic(err)
	}
	SawSprite = ebiten.NewImageFromImage(img)
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(PinchPng))
	if err != nil {
		panic(err)
	}
	PinchSprite = ebiten.NewImageFromImage(img)
}

func LoadMap(name string) (*ldtkgo.Project, error) {
	b, err := MapsFS.ReadFile(fmt.Sprintf("maps/%s.ldtk", name))
	if err != nil {
		return nil, err
	}
	return ldtkgo.Read(b)
}
