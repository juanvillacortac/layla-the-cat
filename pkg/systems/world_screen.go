package systems

import (
	"fmt"
	"image"
	"image/color"
	"layla/pkg/assets"
	"layla/pkg/audio"
	"layla/pkg/components"
	"layla/pkg/config"
	"layla/pkg/data"
	"layla/pkg/events"
	"layla/pkg/factory"
	"layla/pkg/input"
	"layla/pkg/platform"
	esystems "layla/pkg/systems/entities"
	"layla/pkg/text"
	"layla/pkg/touch"
	"math"

	// "math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	etext "github.com/hajimehoshi/ebiten/v2/text"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

func UpdateWorldScreen(ecs *ecs.ECS) {
	components.WorldScreen.Each(ecs.World, func(e *donburi.Entry) {
		world := components.WorldScreen.Get(e)
		camera := components.Camera.Get(e)
		obj := components.GetObject(e)
		obj.H = float64(config.Height)

		spr := components.AnimatedSprite.Get(e)

		// obj.Y = (float64(camera.H) - obj.H) / 2

		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			world.Stroke = touch.NewStroke(&touch.MouseStrokeSource{})
		}
		world.TouchIDs = inpututil.AppendJustPressedTouchIDs(world.TouchIDs[:0])
		for _, id := range world.TouchIDs {
			world.Stroke = touch.NewStroke(&touch.TouchStrokeSource{ID: id})
		}

		if input.Handler.ActionIsJustPressed(input.ActionExit) {
			audio.StopBGM()
			factory.CreateTransition(ecs, true, func() {
				events.LoadLevelEvents.Publish(ecs.World, -1)
			})
		}

		dx, dy := 0, 0

		if world.Stroke != nil && !world.Stroke.IsReleased() {
			dx, dy = world.Stroke.PositionDiff()
		}

		isDragging := dx != 0 || dy != 0

		if input.Handler.ActionIsJustPressed(input.ActionRight) || input.Handler.ActionIsJustPressed(input.ActionLeft) {
			selected := world.Selected
			if input.Handler.ActionIsJustPressed(input.ActionRight) {
				selected++
			} else if input.Handler.ActionIsJustPressed(input.ActionLeft) {
				selected--
			}

			if selected < 0 {
				selected = 0
			} else if world.Selected > len(data.Levels)-1 {
				selected = len(data.Levels) - 1
			}

			if data.SavedData.Levels[selected].Unlocked {
				world.Selected = selected
			}

			var itemObj *resolv.Object
			donburi.NewQuery(filter.Contains(components.LevelItem)).Each(ecs.World, func(e *donburi.Entry) {
				if item := components.LevelItem.Get(e); item.Number == world.Selected {
					itemObj = components.GetObject(e)
				}
			})
			if itemObj != nil {
				obj.X = itemObj.X - float64(config.Width)/2 + itemObj.W/2
				obj.Y = itemObj.Y - float64(config.Height)/2 + itemObj.H/2
				spr.X = itemObj.X
				spr.Y = itemObj.Y
			}
		}

		if input.Handler.ActionIsJustPressed(input.ActionSelect) {
			audio.StopBGM()
			factory.CreateTransition(ecs, true, func() {
				events.LoadLevelEvents.Publish(ecs.World, world.Selected)
			})
		}

		if world.Stroke != nil {
			world.Stroke.Update()
			if world.Stroke.IsReleased() {
				x, y := world.Stroke.Position()
				if x >= components.INPUT_UI_GAP && x <= components.INPUT_UI_GAP+16 && y >= components.INPUT_UI_GAP && y <= components.INPUT_UI_GAP+16 && !isDragging {
					audio.StopBGM()
					factory.CreateTransition(ecs, true, func() {
						events.LoadLevelEvents.Publish(ecs.World, -1)
					})
				}
				world.Stroke = nil
			} else {
				x, y := world.Stroke.PositionDiff()
				if x > 0 {
					obj.X -= math.Min(float64(x)/16, 4)
				} else {
					obj.X -= math.Max(float64(x)/16, -4)
				}
				if y > 0 {
					obj.Y -= math.Min(float64(y)/16, 4)
				} else {
					obj.Y -= math.Max(float64(y)/16, -4)
				}
				// camera.X -= float64(x)
				obj.X = math.Max(0, math.Min(obj.X, float64(camera.W)-obj.W))
				obj.Y = math.Max(0, math.Min(obj.Y, float64(camera.H)-obj.H))
			}
		}
	})
}

func DrawWorldScreen(ecs *ecs.ECS, screen *ebiten.Image) {
	components.WorldScreen.Each(ecs.World, func(e *donburi.Entry) {
		world := components.WorldScreen.Get(e)

		checkImg := assets.WorldUiSprite.SubImage(image.Rect(16*2, 0, 16*3, 16)).(*ebiten.Image)

		if platform.Platform() != platform.Mobile {
			esystems.DrawAnimatedSprites(ecs, components.WorldScreen, screen)
		}

		components.LevelItem.Each(ecs.World, func(e *donburi.Entry) {
			item := components.LevelItem.Get(e)
			if data.SavedData.Levels[item.Number].Cleared {
				obj := components.GetObject(e)
				camera := components.GetCamera(ecs)
				checkOpt := &ebiten.DrawImageOptions{}
				checkOpt.GeoM.Translate(obj.X-math.Round(camera.X)-6, obj.Y-math.Round(camera.Y)-6)
				screen.DrawImage(checkImg, checkOpt)
			}
		})

		homeImg := assets.WorldUiSprite.SubImage(image.Rect(0, 0, 16, 16)).(*ebiten.Image)

		if world.Stroke != nil && !world.Stroke.IsReleased() {
			x, y := world.Stroke.Position()
			dx, dy := world.Stroke.PositionDiff()
			if x >= components.INPUT_UI_GAP && x <= components.INPUT_UI_GAP+16 && y >= components.INPUT_UI_GAP && y <= components.INPUT_UI_GAP+16 && dx == 0 && dy == 0 {
				homeImg = assets.WorldUiSprite.SubImage(image.Rect(0, 16, 16, 16*2)).(*ebiten.Image)
			}
		}

		dx, dy := 0, 0

		if world.Stroke != nil && !world.Stroke.IsReleased() {
			dx, dy = world.Stroke.PositionDiff()
		}

		isDragging := dx != 0 || dy != 0

		if !isDragging {
			opt := &ebiten.DrawImageOptions{}
			opt.GeoM.Translate(float64(components.INPUT_UI_GAP), float64(components.INPUT_UI_GAP))
			screen.DrawImage(homeImg, opt)
		}

		if !isDragging {
			cleared := 0
			donburi.NewQuery(filter.Contains(components.LevelItem)).Each(ecs.World, func(e *donburi.Entry) {
				if item := components.LevelItem.Get(e); data.SavedData.Levels[item.Number].Cleared {
					cleared++
				}
			})

			checkTxt := fmt.Sprintf("%v/%v", cleared, len(data.Levels))

			f := text.LoadFont("ExpressionPro", 16)
			b := etext.BoundString(*f, checkTxt)

			bgW, bgH := 16+4+b.Dx(), 16-3

			checkBgImg := ebiten.NewImage(bgW, bgH)
			checkBgImg.Clear()
			checkBgImg.Fill(color.RGBA{0, 0, 0, 255})

			checkBgImg.Set(0, 0, color.Transparent)
			checkBgImg.Set(0, bgH-1, color.Transparent)
			checkBgImg.Set(0, 0, color.Transparent)
			checkBgImg.Set(bgW-1, bgH-1, color.Transparent)
			checkBgImg.Set(bgW-1, 0, color.Transparent)

			checkBgOpt := &ebiten.DrawImageOptions{}
			checkBgOpt.GeoM.Translate(float64(components.INPUT_UI_GAP)+16+4, float64(components.INPUT_UI_GAP)+2)
			checkBgOpt.ColorScale.ScaleAlpha(0.7)
			screen.DrawImage(checkBgImg, checkBgOpt)

			checkOpt := &ebiten.DrawImageOptions{}
			checkOpt.GeoM.Translate(float64(components.INPUT_UI_GAP)+16+4, float64(components.INPUT_UI_GAP))
			screen.DrawImage(checkImg, checkOpt)

			text.DrawBorderedText(screen, checkTxt, float64(components.INPUT_UI_GAP)+32+4+1, float64(components.INPUT_UI_GAP)+4, false)

			if platform.Platform() != platform.Mobile && !isDragging {
				text.DrawBorderedText(screen, fmt.Sprintf("Level %v", world.Selected+1), float64(components.INPUT_UI_GAP)+1, float64(config.Height-components.INPUT_UI_GAP)-8, false)
			}
		}
	})
}
