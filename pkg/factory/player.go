package factory

import (
	"layla/pkg/archetypes"
	"layla/pkg/components"
	"layla/pkg/config"
	"layla/pkg/entities"
	"math"
	"time"

	"github.com/solarlune/ebitick"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreatePlayer(ecs *ecs.ECS, x, y float64, viewportW, viewportH int) *donburi.Entry {
	player := archetypes.NewPlayer(ecs)

	obj := resolv.NewObject(x, y, 16, 16, "player")
	obj.SetShape(resolv.NewRectangle(0, 0, 16, 16))

	wobj := resolv.NewObject(x, y+4, 16, 16-4, "wallsliding")
	wobj.SetShape(resolv.NewRectangle(0, 0, 16, 16-4))
	p := components.PlayerData{
		FacingRight: true,
		Time:        120,
	}

	ts := ebitick.NewTimerSystem()
	components.TimerSystem.Set(player, ts)
	loop := ts.After(time.Second, func() {
		p.Time = int(math.Max(0, float64(p.Time-1)))
	})
	loop.Loop = true
	p.CountdownTimer = loop

	components.SetObject(player, obj)
	components.Input.Set(player, &components.InputData{})
	components.Entity.Set(player, &components.EntityData{Identifier: string(entities.Player)})
	components.Player.Set(player, &p)
	components.Camera.SetValue(player, components.CameraData{
		X: x - float64(config.Width)/2,
		Y: y - float64(config.Height)/2,
		W: viewportW,
		H: viewportH,
	})

	return player
}
