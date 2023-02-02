package systems

import (
	"layla/pkg/components"
	"layla/pkg/tags"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

func UpdatePushable(ecs *ecs.ECS) {
	tags.Pushable.Each(ecs.World, func(e *donburi.Entry) {
		pushable := components.Pushable.Get(e)
		obj := components.GetObject(e)

		updatePushableMovement(e, pushable, obj, ecs)
	})
}

func updatePushableMovement(entry *donburi.Entry, pushable *components.PushableData, obj *resolv.Object, ecs *ecs.ECS) {
	// playerEntry, ok := tags.Player.First(ecs.World)
	// var player *components.PlayerData
	// if ok {
	// 	player = components.Player.Get(playerEntry)
	// }
	// pushable.SpeedY += GRAVITY
	//
	// if player != nil && player.Pushing == obj {
	// 	pushable.SpeedX = player.SpeedX
	// 	obj.X += player.SpeedX
	// 	if player.FacingRight {
	// 		obj.X += 4
	// 	} else {
	// 		obj.X -= 4
	// 	}
	// 	// fmt.Printf("%.2f\n", player.SpeedX)
	// } else {
	// 	pushable.SpeedX = 0
	// }

	// dx is the horizontal delta movement variable (which is the Player's horizontal speed). If we come into contact with layla, then it will
	// be that movement instead.
	dx := pushable.SpeedX

	// Then we just apply the horizontal movement to the Player's Object. Easy-peasy.

	// Now for the vertical movement; it's the most complicated because we can land on different types of objects and need
	// to treat them all differently, but overall, it's not bad.

	// First, we set OnGround to be nil, in case we don't end up standing on anything.
	pushable.OnGround = nil

	// dy is the delta movement downward, and is the vertical movement by default; similarly to dx, if we come into contact with
	// layla, this will be changed to move to contact instead.

	dy := pushable.SpeedY

	// We want to be sure to lock vertical movement to a maximum of the size of the Cells within the Space
	// so we don't miss any collisions by tunneling through.

	dy = math.Max(math.Min(dy, 16), -16)

	// We're going to check for collision using dy (which is vertical movement speed), but add one when moving downwards to look a bit deeper down
	// into the ground for solid objects to land on, specifically.
	checkDistance := dy
	if dy >= 0 {
		checkDistance++
	}

	// We check for any solid / stand-able objects. In actuality, there aren't any other Objects
	// with other tags in this Space, so we don't -have- to specify any tags, but it's good to be specific for clarity in this example.
	if check := obj.Check(0, checkDistance, "solid", "ramp"); check != nil {
		if ramps := check.ObjectsByTags("ramp"); len(ramps) > 0 {
			ramp := ramps[0]

			// For simplicity, this code assumes we can only stand on one ramp at a time as there is only one ramp in this example.
			// In actuality, if there was a possibility to have a potential collision with multiple ramps (i.e. a ramp that sits on another ramp, and the player running down
			// one onto the other), the collision testing code should probably go with the ramp with the highest confirmed intersection point out of the two.

			// Next, we see if there's been an intersection between the two Shapes using Shape.Intersection. We pass the ramp's shape, and also the movement
			// we're trying to make horizontally, as this makes Intersection return the next y-position while moving, not the one directly
			// underneath the Player. This would keep the player from getting "stuck" when walking up a ramp into the top of a solid block, if there weren't
			// a landing at the top and bottom of the ramp.

			// We use 8 here for the Y-delta so that we can easily see if you're running down the ramp (in which case you're probably in the air as you
			// move faster than you can fall in this example). This way we can maintain contact so you can always jump while running down a ramp. We only
			// continue with coming into contact with the ramp as long as you're not moving upwards (i.e. jumping).

			if contactSet := obj.Shape.Intersection(dx, 8, ramp.Shape); dy >= 0 && contactSet != nil {

				// If Intersection() is successful, a ContactSet is returned. A ContactSet contains information regarding where
				// two Shapes intersect, like the individual points of contact, the center of the contacts, and the MTV, or
				// Minimum Translation Vector, to move out of contact.

				// Here, we use ContactSet.TopmostPoint() to get the top-most contact point as an indicator of where
				// we want the player's feet to be. Then we just set that position, and we're done.

				dy = contactSet.TopmostPoint()[1] - obj.Bottom() + 0.1
				pushable.OnGround = ramp
				pushable.SpeedY = 0
			}
		}

		// Finally, we check for simple solid ground. If we haven't had any success in landing previously, or the solid ground
		// is higher than the existing ground (like if the platform passes underneath the ground, or we're walking off of solid ground
		// onto a ramp), we stand on it instead. We don't check for solid collision first because we want any ramps to override solid
		// ground (so that you can walk onto the ramp, rather than sticking to solid ground).

		// We use ContactWithObject() here because otherwise, we might come into contact with the moving platform's cells (which, naturally,
		// would be selected by a Collision.ContactWithCell() call because the cell is closest to the Player).

		if solids := check.ObjectsByTags("solid"); len(solids) > 0 && (pushable.OnGround == nil || pushable.OnGround.Y >= solids[0].Y) {
			dy = check.ContactWithObject(solids[0]).Y()
			pushable.SpeedY = 0

			// We're only on the ground if we land on it (if the object's Y is greater than the player's).
			if solids[0].Y > obj.Y {
				pushable.OnGround = solids[0]
			}

		}
	}

	// Move the object on dy.
	obj.Y += dy
}

func DrawPushable(ecs *ecs.ECS, screen *ebiten.Image) {
	tags.Pushable.Each(ecs.World, func(e *donburi.Entry) {
		camera := components.GetCamera(ecs)
		o := components.GetObject(e)
		pushable := components.Pushable.Get(e)

		x, y := o.X-math.Round(camera.X), o.Y-math.Round(camera.Y)

		opt := &ebiten.DrawImageOptions{}
		opt.GeoM.Translate(x, y)

		screen.DrawImage(pushable.Tile, opt)
	})
}
