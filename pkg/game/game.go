package game

import (
	"fmt"
	"layla/pkg/config"
	"layla/pkg/events"
	"layla/pkg/platform"
	"layla/pkg/scenes"
	"layla/pkg/shaders"
	"layla/pkg/text"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	etext "github.com/hajimehoshi/ebiten/v2/text"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	e "github.com/yohamta/donburi/features/events"
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
	scene scenes.Scene
	ecs   *ecs.ECS

	vscreen      *ebiten.Image
	ShaderParams []ShaderParam
	shader       *ebiten.Shader
}

func NewGame() *Game {
	g := &Game{
		vscreen: ebiten.NewImage(config.Width, config.Height),
		ecs:     ecs.NewECS(donburi.NewWorld()),
	}

	g.init()

	return g
}

func (g *Game) init() {
	// g.scene = scenes.NewLevelScene(g.ecs, "grass")
	g.scene = scenes.NewTitleScreenScene(g.ecs)

	events.SwitchSceneEvents.Subscribe(g.ecs.World, func(w donburi.World, scene events.SceneEvent) {
		g.scene = scene.Scene
	})

	text.LoadFont("CompassPro", 16)

	if platform.Platform() == platform.Desktop {
		// ebiten.SetFullscreen(true)
		ebiten.SetScreenFilterEnabled(false)
	}

	shader, err := ebiten.NewShader(shaders.ShaderCrtLotterSrc)
	if err != nil {
		log.Fatal(err)
	}
	g.shader = shader

	g.ShaderParams = []ShaderParam{
		{"HardScan", -10., -20., 0., 1.},
		{"HardPix", -4., -20., 0., 1.},
		{"WarpX", 0.04, 0.0, 0.125, 0.01},
		{"WarpY", 0.04, 0.0, 0.125, 0.01},
		{"MaskDark", 0.5, 0.0, 2.0, 0.1},
		{"MaskLight", 1.1, 0.0, 2.0, 0.1},
		{"ShadowMask", 2.0, 0.0, 4.0, 1.0},
		{"BrightBoost", 1.0, 0.0, 2.0, 0.05},
		{"HardBloomPix", -1.5, -2.0, -0.5, 0.1},
		{"HardBloomScan", -2.0, -4.0, -1.0, 0.1},
		{"BloomAmount", 0.2, 0.0, 1.0, 0.05},
		{"Shape", 2.0, 0.0, 10.0, 0.05},
	}
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		config.C.ToggleCrtQuality()
	}

	g.scene.Update()

	e.ProcessAllEvents(g.ecs.World)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	g.vscreen.Clear()

	g.scene.Draw(g.vscreen)

	f := text.LoadFont("ExpressionPro", 16)
	txt := fmt.Sprintf("FPS: %.2f\n", ebiten.ActualFPS())
	x, y := 0., 8.

	opt := &ebiten.DrawImageOptions{}
	opt.ColorScale.ScaleAlpha(0.2)
	opt.GeoM.Translate(x, y)
	etext.DrawWithOptions(g.vscreen, txt, *f, opt)

	if config.C.CrtQuality != config.CrtQualityOff {
		g.drawCRTImage(screen, g.vscreen)
	} else {
		opt := &ebiten.DrawImageOptions{}
		// opt.GeoM.Scale(config.C.Scale, config.C.Scale)
		screen.DrawImage(g.vscreen, opt)
	}
}

func (g *Game) Layout(width, height int) (int, int) {
	l := CalculateLetterBox(f64.Vec2{float64(width), float64(height)}, f64.Vec2{float64(config.BaseWidth), float64(config.BaseHeight)})

	config.Scale = l.Scale
	config.Width, config.Height = width/int(l.Scale), height/int(config.Scale)

	sw, sh := g.vscreen.Size()
	if sw != config.Width || sh != config.Height {
		g.vscreen = ebiten.NewImage(config.Width, config.Height)
	}

	switch config.C.CrtQuality {
	case config.CrtQualityHigh:
		return config.Width * int(l.Scale), config.Height * int(l.Scale)
	default:
		return config.Width, config.Height
	}
}

func (g *Game) drawCRTImage(screen *ebiten.Image, source *ebiten.Image) {
	var screenSize f64.Vec2
	switch config.C.CrtQuality {
	case config.CrtQualityHigh:
		screenSize = f64.Vec2{float64(config.Width * int(config.Scale)), float64(config.Height * int(config.Scale))}
	default:
		screenSize = f64.Vec2{float64(config.Width), float64(config.Height)}
	}

	tw, th := source.Size()
	l := CalculateLetterBox(screenSize, f64.Vec2{float64(tw), float64(th)})

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
		if p.Name == "ShadowMask" {
			if config.C.CrtQuality == config.CrtQualityLow {
				sop.Uniforms[p.Name] = 2.0
			} else {
				sop.Uniforms[p.Name] = 0.0
			}
		} else {
			sop.Uniforms[p.Name] = p.Val
		}
	}

	sop.Images[0] = source

	// recalculate letter box - we are drawing a scaled texture now
	l = CalculateLetterBox(screenSize, f64.Vec2{float64(w), float64(h)})
	sop.GeoM.Concat(l.GetTransform())

	screen.Clear()
	screen.DrawRectShader(config.Width, config.Height, g.shader, sop)
}
