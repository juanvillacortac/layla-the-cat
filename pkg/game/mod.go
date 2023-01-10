package game

import (
	"fmt"
	"layla/pkg/config"
	"layla/pkg/platform"
	"layla/pkg/scenes"
	"layla/pkg/shaders"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/math/f64"
)

type ShaderParam struct {
	Name string
	Val  float32
	Min  float32
	Max  float32
	Step float32
}

type Game struct {
	size  f64.Vec2
	scene scenes.Scene

	vscreen      *ebiten.Image
	ShaderParams []ShaderParam
	shader       *ebiten.Shader
}

func NewGame() *Game {
	g := &Game{
		scene:   scenes.NewLevelScene(),
		vscreen: ebiten.NewImage(config.C.Width, config.C.Height),
	}

	if platform.Platform() == platform.Desktop {
		ebiten.SetFullscreen(true)
	}

	shader, err := ebiten.NewShader(shaders.ShaderCrtLotterSrc)
	if err != nil {
		log.Fatal(err)
	}
	g.shader = shader

	g.ShaderParams = []ShaderParam{
		{"HardScan", -10., -20., 0., 1.},
		{"HardPix", -4., -20., 0., 1.},
		{"WarpX", 0.03, 0.0, 0.125, 0.01},
		{"WarpY", 0.03, 0.0, 0.125, 0.01},
		{"MaskDark", 0.5, 0.0, 2.0, 0.1},
		{"MaskLight", 1.1, 0.0, 2.0, 0.1},
		{"ShadowMask", 2.0, 0.0, 4.0, 1.0},
		{"BrightBoost", 1.0, 0.0, 2.0, 0.05},
		{"HardBloomPix", -1.5, -2.0, -0.5, 0.1},
		{"HardBloomScan", -2.0, -4.0, -1.0, 0.1},
		{"BloomAmount", 0.2, 0.0, 1.0, 0.05},
		{"Shape", 2.0, 0.0, 10.0, 0.05},
	}

	return g
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		config.C.CrtShader = !config.C.CrtShader
	}

	g.scene.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	g.vscreen.Clear()

	g.scene.Draw(g.vscreen)
	ebitenutil.DebugPrint(g.vscreen, fmt.Sprintf("FPS: %.2f\nPlatform: %s", ebiten.ActualFPS(), platform.Platform()))

	if config.C.CrtShader {
		g.drawCRTImage(screen, g.vscreen)
	} else {
		opt := &ebiten.DrawImageOptions{}
		opt.GeoM.Scale(2, 2)
		screen.DrawImage(g.vscreen, opt)
	}
}

func (g *Game) Layout(width, height int) (int, int) {
	g.size = f64.Vec2{float64(width), float64(height)}
	return config.C.Width * 2, config.C.Height * 2
}

func (g *Game) drawCRTImage(screen *ebiten.Image, source *ebiten.Image) {
	tw, th := source.Size()
	l := CalculateLetterBox(f64.Vec2{float64(config.C.Width * 2), float64(config.C.Height * 2)}, f64.Vec2{float64(tw), float64(th)})

	// draw at pixel perfect scale
	// even if we have a bigger screen, we want to stay pixel perfect
	w := tw * int(l.Scale)
	h := th * int(l.Scale)

	sop := &ebiten.DrawRectShaderOptions{}
	sop.GeoM.Scale(l.Scale, l.Scale)
	sop.Uniforms = map[string]any{
		"ScreenSize":  []float32{float32(w), float32(h)},
		"TextureSize": []float32{float32(tw), float32(th)},
	}
	for _, p := range g.ShaderParams {
		sop.Uniforms[p.Name] = p.Val
	}

	sop.Images[0] = source

	// recalculate letter box - we are drawing a scaled texture now
	l = CalculateLetterBox(f64.Vec2{float64(config.C.Width * 2), float64(config.C.Height * 2)}, f64.Vec2{float64(w), float64(h)})
	sop.GeoM.Concat(l.GetTransform())

	screen.Clear()
	screen.DrawRectShader(config.C.Width, config.C.Height, g.shader, sop)
}
