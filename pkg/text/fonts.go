package text

import (
	"fmt"
	"layla/pkg/assets"
	"log"

	"github.com/hajimehoshi/bitmapfont/v2"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const Dpi = 72

var fontFaces map[string]*font.Face = map[string]*font.Face{
	fmt.Sprintf("default___%v", 12): &bitmapfont.Face,
}

func LoadFont(name string, size float64) *font.Face {
	key := fmt.Sprintf("%s___%v", name, size)
	if face, ok := fontFaces[key]; ok {
		return face
	}
	bytes, err := assets.FontsFS.ReadFile(fmt.Sprintf("fonts/%s.ttf", name))
	if err != nil {
		log.Fatal(err)
	}
	tt, err := opentype.Parse(bytes)
	if err != nil {
		log.Fatal(err)
	}
	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    size,
		DPI:     Dpi,
		Hinting: font.HintingNone,
	})
	if err != nil {
		log.Fatal(err)
	}
	fontFaces[key] = &face
	return fontFaces[key]
}

func DrawShadowedText(screen *ebiten.Image, txt string, x, y float64, centered bool) {
	f := LoadFont("ExpressionPro", 16)

	b := text.BoundString(*f, txt)
	if centered {
		x -= float64(b.Dx()) / 2
	}

	y += 8

	stopt := &ebiten.DrawImageOptions{}
	stopt.GeoM.Translate(x, y)
	stopt.GeoM.Translate(1, 1)
	stopt.ColorScale.Scale(0, 0, 0, 1)
	stopt.ColorScale.ScaleAlpha(0.8)
	text.DrawWithOptions(screen, txt, *f, stopt)
	topt := &ebiten.DrawImageOptions{}
	topt.GeoM.Translate(x, y)
	text.DrawWithOptions(screen, txt, *f, topt)
}
