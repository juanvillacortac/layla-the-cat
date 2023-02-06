package components

import (
	"bytes"
	"image"
	_ "image/png"
	"layla/pkg/assets"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/ebitick"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/ganim8/v2"
)

var (
	PlayerSpriteSheet *ebiten.Image
	PlayerSpriteGrid  *ganim8.Grid
	PlayerIdleAnim    *ganim8.Animation
	PlayerRunningAnim *ganim8.Animation
	PlayerDieAnim     *ganim8.Animation
)

type PlayerState int

const (
	PlayerIdle PlayerState = iota
	PlayerRunning
	PlayerJumping
	PlayerWallSliding
	PlayerDie
)

var PlayerAnimations map[PlayerState]*ganim8.Animation

func init() {
	playerDecoded, _, err := image.Decode(bytes.NewReader(assets.PlayerPng))
	if err != nil {
		log.Fatal(err)
	}
	PlayerSpriteSheet = ebiten.NewImageFromImage(playerDecoded)
	PlayerSpriteGrid = ganim8.NewGrid(18, 18, 18*5, 18*4)

	PlayerAnimations = map[PlayerState]*ganim8.Animation{
		PlayerIdle:        ganim8.New(PlayerSpriteSheet, PlayerSpriteGrid.Frames("1-5", 1), 400*time.Millisecond),
		PlayerRunning:     ganim8.New(PlayerSpriteSheet, PlayerSpriteGrid.Frames(3, 1, 1, 2), 100*time.Millisecond),
		PlayerJumping:     ganim8.New(PlayerSpriteSheet, PlayerSpriteGrid.Frames(1, 2), 100*time.Millisecond),
		PlayerDie:         ganim8.New(PlayerSpriteSheet, PlayerSpriteGrid.Frames(1, 4), 100*time.Millisecond),
		PlayerWallSliding: ganim8.New(PlayerSpriteSheet, PlayerSpriteGrid.Frames(1, 3), 100*time.Millisecond),
	}
}

type PlayerData struct {
	SpeedX         float64
	SpeedY         float64
	FacingRight    bool
	Landed         bool
	OnGround       *resolv.Object
	WallSliding    *resolv.Object
	Pushing        *resolv.Object
	State          PlayerState
	CoyoteTime     *ebitick.Timer
	CountdownTimer *ebitick.Timer
	Jumped         int
	Die            bool
	Time           int
}

var Player = donburi.NewComponentType[PlayerData]()
