package systems

import (
	"layla/pkg/components"
	"layla/pkg/config"
	"layla/pkg/maps"
	"regexp"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

var parallaxRegex = regexp.MustCompile(`psf[-]?\d[\d,]*[\.]?[\d{2}]*`)

func drawMapLayer(ecs *ecs.ECS, layer *maps.RenderedLayer, screen *ebiten.Image, centered bool) {
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
		if centered {
			w, h := layer.Image.Size()
			x := float64(config.Width-w) / 2
			y := float64(config.Height-h) / 2
			opt.GeoM.Translate(x, y)
		}
		screen.DrawImage(layer.Image, opt)
	}
}

func DrawImageBackground(ecs *ecs.ECS, screen *ebiten.Image) {
	components.Background.Each(ecs.World, func(e *donburi.Entry) {
		bg := components.Background.Get(e)
		if bg.Image == nil {
			return
		}
		opt := &ebiten.DrawImageOptions{}
		w, h := bg.Image.Size()
		x := float64(config.Width-w) / 2
		y := float64(config.Height-h) / 2
		opt.GeoM.Translate(x, y)
		screen.DrawImage(bg.Image, opt)
	})
}

func DrawMapBackground(ecs *ecs.ECS, screen *ebiten.Image) {
	components.MapRenderer.Each(ecs.World, func(e *donburi.Entry) {
		m := components.MapRenderer.Get(e)
		for _, layer := range m.Renderer.RenderedLayers {
			if strings.Contains(layer.Layer.Identifier, "_bg") {
				drawMapLayer(ecs, layer, screen, m.Centered)
			}
		}
	})
}

func DrawMapForeground(ecs *ecs.ECS, screen *ebiten.Image) {
	components.MapRenderer.Each(ecs.World, func(e *donburi.Entry) {
		m := components.MapRenderer.Get(e)
		for _, layer := range m.Renderer.RenderedLayers {
			if strings.Contains(layer.Layer.Identifier, "_fg") {
				drawMapLayer(ecs, layer, screen, m.Centered)
			}
		}
	})
}
