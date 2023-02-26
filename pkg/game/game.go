package game

import (
	"fmt"
	"image"
	"layla/pkg/audio"
	"layla/pkg/config"
	"layla/pkg/events"
	"layla/pkg/input"
	"layla/pkg/platform"
	"layla/pkg/scenes"
	"layla/pkg/shaders"
	"layla/pkg/text"
	"log"
	"sync"

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

	loaded               bool
	initializedResources sync.Once
	resourceLoadedCh     chan error
}

func NewGame() *Game {
	g := &Game{
		vscreen:          ebiten.NewImage(config.Width, config.Height),
		ecs:              ecs.NewECS(donburi.NewWorld()),
		resourceLoadedCh: make(chan error),
	}

	if err := g.InitResources(); err != nil {
		log.Fatal(err)
	}

	return g
}

func (g *Game) init() {
	// g.scene = scenes.NewLevelScene(g.ecs, "grass")
	g.scene = scenes.NewTitleScreenScene(g.ecs)
	// g.scene = scenes.NewWorldScreenScene(g.ecs)

	events.SwitchSceneEvents.Subscribe(g.ecs.World, func(w donburi.World, scene events.SceneEvent) {
		g.scene = scene.Scene
	})

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

func (g *Game) InitResources() error {
	var err error
	g.initializedResources.Do(func() {
		go func() {
			defer close(g.resourceLoadedCh)
			text.LoadFont("default", 12)
			if err := audio.Load(); err != nil {
				g.resourceLoadedCh <- err
				return
			}
		}()
		if audioErr := audio.Finalize(); audioErr != nil {
			err = audioErr
		}
	})
	return err
}

func (g *Game) Update() error {
	select {
	case err := <-g.resourceLoadedCh:
		if err != nil {
			log.Fatal(err)
		}
		if !g.loaded {
			g.init()
			g.loaded = true
		}
	default:
		return nil
	}
	input.InputSystem.Update()
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
	if !g.loaded {
		return
	}
	g.vscreen.Clear()

	f := text.LoadFont("min", 16)
	txt := fmt.Sprintf("FPS: %.2f\n", ebiten.ActualFPS())
	b := etext.BoundString(*f, txt)

	opt := &ebiten.DrawImageOptions{}
	opt.ColorScale.ScaleAlpha(0.2)
	opt.GeoM.Translate(0, float64(b.Dy())-float64(b.Max.Y))

	if config.C.CrtQuality != config.CrtQualityOff {
		g.scene.Draw(g.vscreen)
		etext.DrawWithOptions(g.vscreen, txt, *f, opt)
		g.drawCRTImage(screen, g.vscreen)
	} else {
		if platform.Platform() == platform.Mobile {
			g.scene.Draw(screen)
			etext.DrawWithOptions(screen, txt, *f, opt)
		} else {
			g.scene.Draw(g.vscreen)
			etext.DrawWithOptions(g.vscreen, txt, *f, opt)
			screen.DrawImage(g.vscreen, &ebiten.DrawImageOptions{})
		}
	}
}

func (g *Game) Layout(width, height int) (int, int) {
	l := CalculateLetterBox(f64.Vec2{float64(width), float64(height)}, f64.Vec2{float64(config.BaseWidth), float64(config.BaseHeight)})

	config.Scale = l.Scale
	config.Width, config.Height = width/int(l.Scale), height/int(config.Scale)

	sw, sh := g.vscreen.Size()
	if sw != config.Width || sh != config.Height {
		g.vscreen = ebiten.NewImageWithOptions(image.Rect(0, 0, config.Width, config.Height), &ebiten.NewImageOptions{
			Unmanaged: false,
		})
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
