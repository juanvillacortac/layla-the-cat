package text

import (
	"fmt"
	"layla/pkg/assets"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const Dpi = 72

var fontFaces map[string]*font.Face = map[string]*font.Face{}

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
