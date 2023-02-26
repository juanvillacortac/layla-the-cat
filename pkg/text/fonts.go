package text

import (
	"fmt"
	"image/color"
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

func DrawBorderedText(screen *ebiten.Image, txt string, x, y float64, centered bool) {
	f := LoadFont("ExpressionPro", 16)

	b := text.BoundString(*f, txt)
	if centered {
		x -= float64(b.Dx()) / 2
	}

	y += 8

	tlOpt := &ebiten.DrawImageOptions{}
	tlOpt.GeoM.Translate(x, y)
	tlOpt.GeoM.Translate(-1, -1)
	tlOpt.ColorScale.ScaleWithColor(color.RGBA{24, 20, 37, 255})
	text.DrawWithOptions(screen, txt, *f, tlOpt)

	blOpt := &ebiten.DrawImageOptions{}
	blOpt.GeoM.Translate(x, y)
	blOpt.GeoM.Translate(-1, 1)
	blOpt.ColorScale.ScaleWithColor(color.RGBA{24, 20, 37, 255})
	text.DrawWithOptions(screen, txt, *f, blOpt)

	trOpt := &ebiten.DrawImageOptions{}
	trOpt.GeoM.Translate(x, y)
	trOpt.GeoM.Translate(1, -1)
	trOpt.ColorScale.ScaleWithColor(color.RGBA{24, 20, 37, 255})
	text.DrawWithOptions(screen, txt, *f, trOpt)

	brOpt := &ebiten.DrawImageOptions{}
	brOpt.GeoM.Translate(x, y)
	brOpt.GeoM.Translate(1, 1)
	brOpt.ColorScale.ScaleWithColor(color.RGBA{24, 20, 37, 255})
	text.DrawWithOptions(screen, txt, *f, brOpt)

	lOpt := &ebiten.DrawImageOptions{}
	lOpt.GeoM.Translate(x, y)
	lOpt.GeoM.Translate(-1, 0)
	lOpt.ColorScale.ScaleWithColor(color.RGBA{24, 20, 37, 255})
	text.DrawWithOptions(screen, txt, *f, lOpt)

	rOpt := &ebiten.DrawImageOptions{}
	rOpt.GeoM.Translate(x, y)
	rOpt.GeoM.Translate(1, 0)
	rOpt.ColorScale.ScaleWithColor(color.RGBA{24, 20, 37, 255})
	text.DrawWithOptions(screen, txt, *f, rOpt)

	topt := &ebiten.DrawImageOptions{}
	topt.GeoM.Translate(x, y)
	text.DrawWithOptions(screen, txt, *f, topt)
}
