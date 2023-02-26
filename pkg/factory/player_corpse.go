package factory

import (
	"layla/pkg/archetypes"
	"layla/pkg/audio"
	"layla/pkg/components"
	"layla/pkg/entities"
	"time"

	"github.com/solarlune/ebitick"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func CreatePlayerCorpse(ecs *ecs.ECS, obj *resolv.Object) *donburi.Entry {
	audio.StopBGM()
	corpse := archetypes.NewPlayerCorpse(ecs)
	ts := ebitick.NewTimerSystem()
	components.TimerSystem.Set(corpse, ts)
	data := &components.PlayerCorpseData{
		SpeedY: -2.9,
		X:      obj.X,
		Y:      obj.Y,
	}
	components.PlayerCorpse.Set(corpse, data)
	components.Entity.Set(corpse, &components.EntityData{Identifier: string(entities.PlayerCorpse), Layer: components.EntityFrontLayer})
	ts.After(time.Millisecond*500, func() {
		data.Fall = true
		audio.PlaySE("die.wav")
	})
	return corpse
}
