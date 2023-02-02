package components

import (
	"bytes"
	"image"
	_ "image/png"
	"layla/pkg/assets"
	"layla/pkg/config"
	"layla/pkg/touch"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
)

var (
	INPUT_UI_GAP             int
	INPUT_UI_PIXEL_CELL_SIZE int
	INPUT_UI_CELL_SIZE       int

	INPUT_UI_LEFT_X  int
	INPUT_UI_LEFT_Y  int
	INPUT_UI_RIGHT_X int
	INPUT_UI_RIGHT_Y int

	INPUT_UI_JUMP_X int
	INPUT_UI_JUMP_Y int

	INPUT_UI_RELEASE_X int
	INPUT_UI_RELEASE_Y int

	INPUT_UI_PAUSE_X int
	INPUT_UI_PAUSE_Y int
)

func UpdateInputLayout() {
	INPUT_UI_GAP = 8 * int(config.C.InputScale)
	INPUT_UI_PIXEL_CELL_SIZE = 24
	INPUT_UI_CELL_SIZE = INPUT_UI_PIXEL_CELL_SIZE * int(config.C.InputScale)

	INPUT_UI_LEFT_X = INPUT_UI_GAP
	INPUT_UI_LEFT_Y = config.Height - INPUT_UI_GAP - INPUT_UI_CELL_SIZE
	INPUT_UI_RIGHT_X = INPUT_UI_LEFT_X + INPUT_UI_GAP + INPUT_UI_CELL_SIZE
	INPUT_UI_RIGHT_Y = INPUT_UI_LEFT_Y

	INPUT_UI_JUMP_X = config.Width - INPUT_UI_GAP - INPUT_UI_CELL_SIZE
	INPUT_UI_JUMP_Y = INPUT_UI_LEFT_Y

	INPUT_UI_RELEASE_X = config.Width - INPUT_UI_GAP - INPUT_UI_CELL_SIZE
	INPUT_UI_RELEASE_Y = INPUT_UI_LEFT_Y - INPUT_UI_GAP - INPUT_UI_CELL_SIZE

	INPUT_UI_PAUSE_X = config.Width - INPUT_UI_GAP - 12
	INPUT_UI_PAUSE_Y = INPUT_UI_GAP
}

var (
	InputUiImage *ebiten.Image
	InputUiPause *ebiten.Image
)

type InputImageKey int

const (
	InputImageMovement InputImageKey = iota
	InputImageJump
	InputImageRelease
)

var InputImage map[InputImageKey]map[bool]*ebiten.Image = map[InputImageKey]map[bool]*ebiten.Image{}

func init() {
	UpdateInputLayout()

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
				X: INPUT_UI_PIXEL_CELL_SIZE, Y: INPUT_UI_PIXEL_CELL_SIZE,
			},
		}).(*ebiten.Image),
		true: InputUiImage.SubImage(image.Rectangle{
			Min: image.Point{
				X: INPUT_UI_PIXEL_CELL_SIZE,
				Y: 0,
			},
			Max: image.Point{
				X: INPUT_UI_PIXEL_CELL_SIZE * 2, Y: INPUT_UI_PIXEL_CELL_SIZE,
			},
		}).(*ebiten.Image),
	}
	InputImage[InputImageJump] = map[bool]*ebiten.Image{
		false: InputUiImage.SubImage(image.Rectangle{
			Min: image.Point{
				X: 0,
				Y: INPUT_UI_PIXEL_CELL_SIZE,
			},
			Max: image.Point{
				X: INPUT_UI_PIXEL_CELL_SIZE, Y: INPUT_UI_PIXEL_CELL_SIZE * 2,
			},
		}).(*ebiten.Image),
		true: InputUiImage.SubImage(image.Rectangle{
			Min: image.Point{
				X: INPUT_UI_PIXEL_CELL_SIZE,
				Y: INPUT_UI_PIXEL_CELL_SIZE,
			},
			Max: image.Point{
				X: INPUT_UI_PIXEL_CELL_SIZE * 2, Y: INPUT_UI_PIXEL_CELL_SIZE * 2,
			},
		}).(*ebiten.Image),
	}
	InputImage[InputImageRelease] = map[bool]*ebiten.Image{
		false: InputUiImage.SubImage(image.Rectangle{
			Min: image.Point{
				X: 0,
				Y: INPUT_UI_PIXEL_CELL_SIZE * 2,
			},
			Max: image.Point{
				X: INPUT_UI_PIXEL_CELL_SIZE, Y: INPUT_UI_PIXEL_CELL_SIZE * 3,
			},
		}).(*ebiten.Image),
		true: InputUiImage.SubImage(image.Rectangle{
			Min: image.Point{
				X: INPUT_UI_PIXEL_CELL_SIZE,
				Y: INPUT_UI_PIXEL_CELL_SIZE * 2,
			},
			Max: image.Point{
				X: INPUT_UI_PIXEL_CELL_SIZE * 2, Y: INPUT_UI_PIXEL_CELL_SIZE * 3,
			},
		}).(*ebiten.Image),
	}
	InputUiPause = InputUiImage.SubImage(image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: INPUT_UI_PIXEL_CELL_SIZE * 3,
		},
		Max: image.Point{
			X: INPUT_UI_PIXEL_CELL_SIZE + 12, Y: INPUT_UI_PIXEL_CELL_SIZE*3 + 13,
		},
	}).(*ebiten.Image)
}

type InputData struct {
	TouchIDs   []ebiten.TouchID
	GamepadIDs []ebiten.GamepadID
	Sliding    bool
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
			tx, ty := touch.TouchPos(id)
			if tx >= INPUT_UI_LEFT_X && ty >= INPUT_UI_LEFT_Y && tx <= INPUT_UI_LEFT_X+INPUT_UI_CELL_SIZE && ty <= INPUT_UI_LEFT_Y+INPUT_UI_CELL_SIZE {
				return true
			}
		}
		return ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA)
	default:
		for _, id := range i.TouchIDs {
			tx, ty := touch.TouchPos(id)
			if tx >= INPUT_UI_RIGHT_X && ty >= INPUT_UI_RIGHT_Y && tx <= INPUT_UI_RIGHT_X+INPUT_UI_CELL_SIZE && ty <= INPUT_UI_RIGHT_Y+INPUT_UI_CELL_SIZE {
				return true
			}
		}
		return ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD)
	}
}

func (i *InputData) IsReleasing() bool {
	i.TouchIDs = inpututil.AppendJustPressedTouchIDs(i.TouchIDs[:0])
	for _, id := range i.TouchIDs {
		tx, ty := touch.TouchPos(id)
		if tx >= INPUT_UI_RELEASE_X && ty >= INPUT_UI_RELEASE_Y && tx <= INPUT_UI_RELEASE_X+INPUT_UI_CELL_SIZE && ty <= INPUT_UI_RELEASE_Y+INPUT_UI_CELL_SIZE {
			return true
		}
	}
	keyboardDown := ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS)
	gamepadDown := len(i.GamepadIDs) > 0 && ebiten.IsGamepadButtonPressed(i.GamepadIDs[0], ebiten.GamepadButton(ebiten.StandardGamepadButtonLeftBottom))
	return (keyboardDown || gamepadDown) && i.IsJustJumping()
}

func (i *InputData) IsLongJumping() bool {
	i.TouchIDs = ebiten.AppendTouchIDs(i.TouchIDs[:0])
	for _, id := range i.TouchIDs {
		tx, ty := touch.TouchPos(id)
		if tx >= INPUT_UI_JUMP_X && ty >= INPUT_UI_JUMP_Y && tx <= INPUT_UI_JUMP_X+INPUT_UI_CELL_SIZE && ty <= INPUT_UI_JUMP_Y+INPUT_UI_CELL_SIZE {
			return true
		}
	}
	return ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsKeyPressed(ebiten.KeyW)
}
func (i *InputData) IsJustJumping() bool {
	i.TouchIDs = inpututil.AppendJustPressedTouchIDs(i.TouchIDs[:0])
	for _, id := range i.TouchIDs {
		tx, ty := touch.TouchPos(id)
		if tx >= INPUT_UI_JUMP_X && ty >= INPUT_UI_JUMP_Y && tx <= INPUT_UI_JUMP_X+INPUT_UI_CELL_SIZE && ty <= INPUT_UI_JUMP_Y+INPUT_UI_CELL_SIZE {
			return true
		}
	}

	return inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsKeyJustPressed(ebiten.KeyW)
}

func (i *InputData) IsRequestPause() bool {
	i.TouchIDs = inpututil.AppendJustPressedTouchIDs(i.TouchIDs[:0])
	for _, id := range i.TouchIDs {
		tx, ty := touch.TouchPos(id)
		if tx >= INPUT_UI_PAUSE_X && ty >= INPUT_UI_PAUSE_Y && tx <= INPUT_UI_PAUSE_X+12 && ty <= INPUT_UI_PAUSE_Y+13 {
			return true
		}
	}

	return inpututil.IsKeyJustPressed(ebiten.KeyEscape)
}
