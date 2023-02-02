package esystems

import (
	"fmt"
	"layla/pkg/components"
	"layla/pkg/events"
	"layla/pkg/factory"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/solarlune/ebitick"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/ganim8/v2"
)

var (
	PLAYER_FRICTION     = 0.3
	PLAYER_ACCELERATION = 0.3 + PLAYER_FRICTION
	PLAYER_MAX_SPEED    = 2.0
	PLAYER_JUMP_SPEED   = 3.0
	GRAVITY             = 0.15
)

func UpdatePlayer(ecs *ecs.ECS, e *donburi.Entry) {
	player := components.Player.Get(e)
	if player.Die {
		return
	}
	playerObject := components.GetObject(e)

	components.PlayerAnimations[player.State].Update()

	updatePlayerMovement(e, player, playerObject, ecs)
}

func updatePlayerMovement(entry *donburi.Entry, player *components.PlayerData, playerObject *resolv.Object, ecs *ecs.ECS) {
	input := components.Input.Get(entry)
	input.Sliding = player.WallSliding != nil

	ts := components.TimerSystem.Get(entry)

	if player.OnGround == nil {
		player.Landed = false
	}

	// Now we update the Player's movement. This is the real bread-and-butter of this example, naturally.
	player.SpeedY += GRAVITY
	if player.WallSliding != nil && player.SpeedY > 0.1 {
		player.SpeedY = 0.1
	}

	// Horizontal movement is only possible when not wallsliding.
	if player.WallSliding == nil {
		if input.IsRunning(components.InputDirectionRight) {
			player.SpeedX += PLAYER_ACCELERATION
			player.FacingRight = true
		}

		if input.IsRunning(components.InputDirectionLeft) {
			player.SpeedX -= PLAYER_ACCELERATION
			player.FacingRight = false
		}
	}

	// Apply friction and horizontal speed limiting.
	if player.SpeedX > PLAYER_FRICTION {
		player.SpeedX -= PLAYER_FRICTION
	} else if player.SpeedX < -PLAYER_FRICTION {
		player.SpeedX += PLAYER_FRICTION
	} else {
		player.SpeedX = 0
	}

	if player.SpeedX > PLAYER_MAX_SPEED {
		player.SpeedX = PLAYER_MAX_SPEED
	} else if player.SpeedX < -PLAYER_MAX_SPEED {
		player.SpeedX = -PLAYER_MAX_SPEED
	}

	wallOnHead := playerObject.Check(0, -1, "solid")

	allowCoyote := player.CoyoteTime != nil && player.CoyoteTime.State == ebitick.StateRunning && player.Jumped == 0
	normalJump := player.OnGround != nil

	// Check for jumping.

	if input.IsReleasing() {
		player.WallSliding = nil
	} else if input.IsJustJumping() {
		if (normalJump || allowCoyote) && wallOnHead == nil && player.WallSliding == nil {
			player.Jumped += 1
			if player.CoyoteTime != nil {
				player.CoyoteTime.Cancel()
			}
			player.SpeedY = -PLAYER_JUMP_SPEED

			px := playerObject.X
			if player.FacingRight {
				px -= 4
			} else {
				px += 4
			}
			factory.CreateParticles(ecs, components.ParticlesBackLayer, components.ParticlesJump, px, playerObject.Y, !player.FacingRight)
		} else if player.WallSliding != nil {
			// WALLJUMPING
			player.SpeedY = -PLAYER_JUMP_SPEED

			if player.WallSliding.X > playerObject.X {
				player.SpeedX = -2
			} else {
				player.SpeedX = 2
			}

			player.WallSliding = nil
			player.FacingRight = !player.FacingRight
		}
	}

	// We handle horizontal movement separately from vertical movement. This is, conceptually, decomposing movement into two phases / axes.
	// By decomposing movement in this manner, we can handle each case properly (i.e. stop movement horizontally separately from vertical movement, as
	// necesseary). More can be seen on this topic over on this blog post on higherorderfun.com:
	// http://higherorderfun.com/blog/2012/05/20/the-guide-to-implementing-2d-platformers/

	// dx is the horizontal delta movement variable (which is the Player's horizontal speed). If we come into contact with something, then it will
	// be that movement instead.
	dx := player.SpeedX

	// Moving horizontally is done fairly simply; we just check to see if something solid is in front of us. If so, we move into contact with it
	// and stop horizontal movement speed. If not, then we can just move forward.

	if check := playerObject.Check(player.SpeedX, 0, "solid"); check != nil {
		dx = check.ContactWithObject(check.Objects[0]).X()
		player.SpeedX = 0

		sliding := check.Objects[0]
		// If you're in the air, then colliding with a wall object makes you start wall sliding.
		if player.OnGround == nil && playerObject.Y < sliding.Bottom()-10 {
			player.WallSliding = sliding
		}
	}

	if check := playerObject.Check(player.SpeedX, 0, "pushable"); check != nil && (player.SpeedX > 0 || player.SpeedX < 0) && player.OnGround != nil {
		fmt.Printf("%.2f\n", player.SpeedX)
		pushable := check.Objects[0]

		pushable.X += player.SpeedX

		// dx = check.ContactWithObject(pushable).X()
		player.Pushing = pushable

	} else {
		player.Pushing = nil
	}

	// Then we just apply the horizontal movement to the Player's Object. Easy-peasy.
	playerObject.X += dx

	// Now for the vertical movement; it's the most complicated because we can land on different types of objects and need
	// to treat them all differently, but overall, it's not bad.

	// First, we set OnGround to be nil, in case we don't end up standing on anything.
	player.OnGround = nil

	// dy is the delta movement downward, and is the vertical movement by default; similarly to dx, if we come into contact with
	// something, this will be changed to move to contact instead.

	dy := player.SpeedY

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
	if check := playerObject.Check(0, checkDistance, "solid", "ramp", "pushable", "enemy"); check != nil {
		// So! Firstly, we want to see if we jumped up into something that we can slide around horizontally to avoid bumping the Player's head.

		// Sliding around a misspaced jump is a small thing that makes jumping a bit more forgiving, and is something different polished platformers
		// (like the 2D Mario games) do to make it a smidge more comfortable to play. For a visual example of this, see this excellent devlog post
		// from the extremely impressive indie game, Leilani's Island: https://forums.tigsource.com/index.php?topic=46289.msg1387138#msg1387138

		// To accomplish this sliding, we simply call Collision.SlideAgainstCell() to see if we can slide.
		// We pass the first cell, and tags that we want to avoid when sliding (i.e. we don't want to slide into cells that contain other solid objects).

		if enemies := check.ObjectsByTags("enemy"); len(enemies) > 0 {
			KillPlayer(ecs, entry)
		}

		slide := check.SlideAgainstCell(check.Cells[0], "solid")

		// We further ensure that we only slide if:
		// 1) We're jumping up into something (dy < 0),
		// 2) If the cell we're bumping up against contains a solid object,
		// 3) If there was, indeed, a valid slide left or right, and
		// 4) If the proposed slide is less than 8 pixels in horizontal distance. (This is a relatively arbitrary number that just so happens to be half the
		// width of a cell. This is to ensure the player doesn't slide too far horizontally.)

		if dy < 0 && check.Cells[0].ContainsTags("solid") && slide != nil && math.Abs(slide.X()) <= 8 {

			// If we are able to slide here, we do so. No contact was made, and vertical speed (dy) is maintained upwards.
			playerObject.X += slide.X()

		} else {

			// If sliding -fails-, that means the Player is jumping directly onto or into something, and we need to do more to see if we need to come into
			// contact with it. Let's press on!

			// First, we check for ramps. For ramps, we can't simply check for collision with Check(), as that's not precise enough. We need to get a bit
			// more information, and so will do so by checking its Shape (a triangular ConvexPolygon, as defined in WorldPlatformer.Init()) against the
			// Player's Shape (which is also a rectangular ConvexPolygon).

			// We get the ramp by simply filtering out Objects with the "ramp" tag out of the objects returned in our broad Check(), and grabbing the first one
			// if there's any at all.
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

				if contactSet := playerObject.Shape.Intersection(dx, 8, ramp.Shape); dy >= 0 && contactSet != nil {

					// If Intersection() is successful, a ContactSet is returned. A ContactSet contains information regarding where
					// two Shapes intersect, like the individual points of contact, the center of the contacts, and the MTV, or
					// Minimum Translation Vector, to move out of contact.

					// Here, we use ContactSet.TopmostPoint() to get the top-most contact point as an indicator of where
					// we want the player's feet to be. Then we just set that position, and we're done.

					dy = contactSet.TopmostPoint()[1] - playerObject.Bottom() + 0.1
					player.OnGround = ramp
					player.SpeedY = 0
				}
			}

			if pushables := check.ObjectsByTags("pushable"); len(pushables) > 0 {
				pushable := pushables[0]

				if player.SpeedY >= 0 && playerObject.Bottom() < pushable.Y+4 {
					dy = check.ContactWithObject(pushable).Y()
					player.OnGround = pushable
					player.SpeedY = 0
				}
			}

			// Finally, we check for simple solid ground. If we haven't had any success in landing previously, or the solid ground
			// is higher than the existing ground (like if the platform passes underneath the ground, or we're walking off of solid ground
			// onto a ramp), we stand on it instead. We don't check for solid collision first because we want any ramps to override solid
			// ground (so that you can walk onto the ramp, rather than sticking to solid ground).

			// We use ContactWithObject() here because otherwise, we might come into contact with the moving platform's cells (which, naturally,
			// would be selected by a Collision.ContactWithCell() call because the cell is closest to the Player).

			if solids := check.ObjectsByTags("solid"); len(solids) > 0 && (player.OnGround == nil || player.OnGround.Y >= solids[0].Y) {
				dy = check.ContactWithObject(solids[0]).Y()
				// prevSpeed := player.SpeedY
				player.SpeedY = 0

				// We're only on the ground if we land on it (if the object's Y is greater than the player's).
				if solids[0].Y > playerObject.Y {
					player.OnGround = solids[0]
				}
			}

			if player.OnGround != nil {
				player.WallSliding = nil // Player's on the ground, so no wallsliding anymore.
				player.Jumped = 0
				if !player.Landed {
					factory.CreateParticles(ecs, components.ParticlesFrontLayer, components.ParticlesFall, playerObject.X, player.OnGround.Y-player.OnGround.H, !player.FacingRight)
					player.CoyoteTime = nil
					// ebiten.Vibrate(&ebiten.VibrateOptions{
					// 	Duration:  200 * time.Millisecond,
					// 	Magnitude: 0.2,
					// })
				}
				player.Landed = true
			}
		}
	}

	// Move the object on dy.
	playerObject.Y += dy

	wallNext := 1.0
	if !player.FacingRight {
		wallNext = -1
	}

	if player.OnGround == nil {
		if player.WallSliding != nil {
			player.State = components.PlayerWallSliding
		} else {
			player.State = components.PlayerJumping
		}
	} else if player.SpeedX > PLAYER_FRICTION || player.SpeedX < -PLAYER_FRICTION || player.Pushing != nil {
		player.State = components.PlayerRunning
	} else {
		player.State = components.PlayerIdle
	}

	// If the wall next to the Player runs out, stop wall sliding.
	if c := playerObject.Check(wallNext, 0, "solid"); player.WallSliding != nil && c == nil {
		player.WallSliding = nil
	}

	if player.WallSliding != nil {
		onBottom := player.WallSliding.Check(0, 16, "solid")
		if onBottom == nil && player.WallSliding.Bottom() < playerObject.Y+10 {
			player.WallSliding = nil
		}
	}

	if player.OnGround == nil && player.CoyoteTime == nil {
		player.CoyoteTime = ts.After(200*time.Millisecond, func() {})
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		KillPlayer(ecs, entry)
	}
}

func KillPlayer(ecs *ecs.ECS, e *donburi.Entry) {
	player := components.Player.Get(e)
	playerObject := components.GetObject(e)
	ts := components.TimerSystem.Get(e)
	if player.Die {
		return
	}
	player.Die = true
	factory.CreatePlayerCorpse(ecs, playerObject)
	components.ShakeCamera(ecs, 4, time.Millisecond*200)
	factory.CreateFlash(ecs, time.Millisecond*100)
	ts.After(time.Second*2, func() {
		factory.CreateTransition(ecs, true, func() {
			events.RestartLevelEvents.Publish(ecs.World, events.RestartLevelEvent{})
		})
	})
}

func DrawPlayer(ecs *ecs.ECS, e *donburi.Entry, screen *ebiten.Image) {
	player := components.Player.Get(e)
	if player.Die {
		return
	}
	camera := components.Camera.Get(e)

	o := components.GetObject(e)
	x, y := o.X-math.Round(camera.X)-1, o.Y-math.Round(camera.Y)-1

	var opts *ganim8.DrawOptions
	if !player.FacingRight {
		opts = ganim8.DrawOpts(x, y, 0, -1, 1, 1, 0)
	} else {
		opts = ganim8.DrawOpts(x, y)
	}
	components.PlayerAnimations[player.State].Draw(screen, opts)
}
