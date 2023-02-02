package systems

import (
	"layla/pkg/components"
	"layla/pkg/maps"
	"layla/pkg/tags"
	"regexp"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

var parallaxRegex = regexp.MustCompile(`psf[-]?\d[\d,]*[\.]?[\d{2}]*`)

func drawLevelLayer(ecs *ecs.ECS, layer *maps.RenderedLayer, screen *ebiten.Image) {
	if layer.Image != nil {
		psf := 100.0
		if match := strings.TrimPrefix(string(parallaxRegex.Find([]byte(layer.Layer.Identifier))), "psf"); match != "" {
			if v, err := strconv.ParseFloat(match, 64); err == nil {
				psf = v
			}
		}
		opt := &ebiten.DrawImageOptions{}
		if camera := components.GetCamera(ecs); camera != nil && !strings.Contains(layer.Layer.Identifier, "_static") {
			opt.GeoM.Translate(-camera.X/(psf/100), -camera.Y/(psf/100))
		}
		screen.DrawImage(layer.Image, opt)
	}
}

func DrawLevelBackground(ecs *ecs.ECS, screen *ebiten.Image) {
	tags.Level.Each(ecs.World, func(e *donburi.Entry) {
		level := components.Level.Get(e)
		for _, layer := range level.Renderer.RenderedLayers {
			if strings.Contains(layer.Layer.Identifier, "_bg") {
				drawLevelLayer(ecs, layer, screen)
			}
		}
	})

}
func DrawLevelForeground(ecs *ecs.ECS, screen *ebiten.Image) {
	tags.Level.Each(ecs.World, func(e *donburi.Entry) {
		level := components.Level.Get(e)
		for _, layer := range level.Renderer.RenderedLayers {
			if strings.Contains(layer.Layer.Identifier, "_fg") {
				drawLevelLayer(ecs, layer, screen)
			}
		}
	})
}
