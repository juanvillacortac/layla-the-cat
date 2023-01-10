package components

import (
	"bytes"
	"image"
	_ "image/png"
	"layla/pkg/assets"
	"layla/pkg/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
)

var (
	INPUT_UI_GAP       = 8
	INPUT_UI_CELL_SIZE = 24

	INPUT_UI_LEFT_X  = INPUT_UI_GAP * 2
	INPUT_UI_LEFT_Y  = config.C.Height - INPUT_UI_GAP*2 - INPUT_UI_CELL_SIZE*2
	INPUT_UI_RIGHT_X = (INPUT_UI_GAP) + INPUT_UI_CELL_SIZE*3
	INPUT_UI_RIGHT_Y = config.C.Height - INPUT_UI_GAP*2 - INPUT_UI_CELL_SIZE*2

	INPUT_UI_JUMP_X = config.C.Width - INPUT_UI_GAP*2 - INPUT_UI_CELL_SIZE*2
	INPUT_UI_JUMP_Y = config.C.Height - INPUT_UI_GAP*2 - INPUT_UI_CELL_SIZE*2
)

var (
	InputUiImage *ebiten.Image
)

type InputImageKey int

const (
	InputImageMovement InputImageKey = iota
	InputImageJump
)

var InputImage map[InputImageKey]map[bool]*ebiten.Image = map[InputImageKey]map[bool]*ebiten.Image{}

func init() {
	img, _, err := image.Decode(bytes.NewReader(assets.InputPng))
	if err != nil {
		panic(err)
	}
	InputUiImage = ebiten.NewImageFromImage(img)
	InputImage[InputImageMovement] = map[bool]*ebiten.Image{
		false: InputUiImage.SubImage(image.Rectangle{
			Min: image.Point{
				X: 0,
				Y: 0,
			},
			Max: image.Point{
				X: INPUT_UI_CELL_SIZE, Y: INPUT_UI_CELL_SIZE,
			},
		}).(*ebiten.Image),
		true: InputUiImage.SubImage(image.Rectangle{
			Min: image.Point{
				X: INPUT_UI_CELL_SIZE,
				Y: 0,
			},
			Max: image.Point{
				X: INPUT_UI_CELL_SIZE * 2, Y: INPUT_UI_CELL_SIZE,
			},
		}).(*ebiten.Image),
	}
	InputImage[InputImageJump] = map[bool]*ebiten.Image{
		false: InputUiImage.SubImage(image.Rectangle{
			Min: image.Point{
				X: 0,
				Y: INPUT_UI_CELL_SIZE,
			},
			Max: image.Point{
				X: INPUT_UI_CELL_SIZE, Y: INPUT_UI_CELL_SIZE * 2,
			},
		}).(*ebiten.Image),
		true: InputUiImage.SubImage(image.Rectangle{
			Min: image.Point{
				X: INPUT_UI_CELL_SIZE,
				Y: INPUT_UI_CELL_SIZE,
			},
			Max: image.Point{
				X: INPUT_UI_CELL_SIZE * 2, Y: INPUT_UI_CELL_SIZE * 2,
			},
		}).(*ebiten.Image),
	}
}

type InputData struct {
	TouchIDs []ebiten.TouchID
}

var Input = donburi.NewComponentType[InputData]()

type InputDirection int

const (
	InputDirectionLeft InputDirection = iota
	InputDirectionRight
)

func (i *InputData) IsRunning(dir InputDirection) bool {
	i.TouchIDs = ebiten.AppendTouchIDs(i.TouchIDs[:0])
	switch dir {
	case InputDirectionLeft:
		for _, id := range i.TouchIDs {
			tx, ty := ebiten.TouchPosition(id)
			tx, ty = tx/2, ty/2
			if tx >= INPUT_UI_LEFT_X && ty >= INPUT_UI_LEFT_Y && tx <= INPUT_UI_LEFT_X+INPUT_UI_CELL_SIZE*2 && ty <= INPUT_UI_LEFT_Y+INPUT_UI_CELL_SIZE*2 {
				return true
			}
		}
		return ebiten.IsKeyPressed(ebiten.KeyLeft)
	default:
		for _, id := range i.TouchIDs {
			tx, ty := ebiten.TouchPosition(id)
			tx, ty = tx/2, ty/2
			if tx >= INPUT_UI_RIGHT_X && ty >= INPUT_UI_RIGHT_Y && tx <= INPUT_UI_RIGHT_X+INPUT_UI_CELL_SIZE*2 && ty <= INPUT_UI_RIGHT_Y+INPUT_UI_CELL_SIZE*2 {
				return true
			}
		}
		return ebiten.IsKeyPressed(ebiten.KeyRight)
	}
}

func (i *InputData) IsLongJumping() bool {
	i.TouchIDs = ebiten.AppendTouchIDs(i.TouchIDs[:0])
	for _, id := range i.TouchIDs {
		tx, ty := ebiten.TouchPosition(id)
		tx, ty = tx/2, ty/2
		if tx >= INPUT_UI_JUMP_X && ty >= INPUT_UI_JUMP_Y && tx <= INPUT_UI_JUMP_X+INPUT_UI_CELL_SIZE*2 && ty <= INPUT_UI_JUMP_Y+INPUT_UI_CELL_SIZE*2 {
			return true
		}
	}
	return ebiten.IsKeyPressed(ebiten.KeySpace)
}
func (i *InputData) IsJustJumping() bool {
	i.TouchIDs = inpututil.AppendJustPressedTouchIDs(i.TouchIDs[:0])
	for _, id := range i.TouchIDs {
		tx, ty := ebiten.TouchPosition(id)
		tx, ty = tx/2, ty/2
		if tx >= INPUT_UI_JUMP_X && ty >= INPUT_UI_JUMP_Y && tx <= INPUT_UI_JUMP_X+INPUT_UI_CELL_SIZE*2 && ty <= INPUT_UI_JUMP_Y+INPUT_UI_CELL_SIZE*2 {
			return true
		}
	}

	return inpututil.IsKeyJustPressed(ebiten.KeySpace)
}
