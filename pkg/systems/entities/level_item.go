package esystems

import (
	"layla/pkg/audio"
	"layla/pkg/components"
	"layla/pkg/data"
	"layla/pkg/factory"

	"layla/pkg/events"
	"layla/pkg/touch"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

func UpdateLevelItem(ecs *ecs.ECS, e *donburi.Entry) {
	if !e.HasComponent(components.LevelItem) {
		return
	}
	item := components.LevelItem.Get(e)
	obj := components.GetObject(e)
	camera := components.GetCamera(ecs)
	if camera == nil {
		return
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		item.Stroke = touch.NewStroke(&touch.MouseStrokeSource{})
	}
	item.TouchIDs = inpututil.AppendJustPressedTouchIDs(item.TouchIDs[:0])
	for _, id := range item.TouchIDs {
		item.Stroke = touch.NewStroke(&touch.TouchStrokeSource{ID: id})
	}

	spr := components.AnimatedSprite.Get(e)

	savedData, ok := data.SavedData.Levels[item.Number]

	if !ok || !savedData.Unlocked {
		spr.Anim.GoToFrame(3)
	} else {
		spr.Anim.GoToFrame(1)
	}

	if item.Stroke != nil && ok && savedData.Unlocked {
		item.Stroke.Update()
		mx, my := item.Stroke.Position()
		dx, dy := item.Stroke.PositionDiff()
		x, y := float64(mx), float64(my)
		count := query.NewQuery(filter.Contains(components.Transition)).Count(ecs.World)
		if item.Stroke.IsReleased() {
			item.Stroke = nil
			if x >= obj.X-camera.X && y >= obj.Y-camera.Y && x <= obj.Right()-camera.X && y <= obj.Bottom()-camera.Y && count == 0 && dx == 0 && dy == 0 {
				audio.StopBGM()
				factory.CreateTransition(ecs, true, func() {
					events.LoadLevelEvents.Publish(ecs.World, item.Number)
				})
			}
		} else {
			if x >= obj.X-camera.X && y >= obj.Y-camera.Y && x <= obj.Right()-camera.X && y <= obj.Bottom()-camera.Y && count == 0 && dx == 0 && dy == 0 {
				spr.Anim.GoToFrame(2)
			} else {
				spr.Anim.GoToFrame(1)
			}
		}
	}
}
