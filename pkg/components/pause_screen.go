package components

import (
	"image"
	"image/color"
	"layla/pkg/config"
	"layla/pkg/input"
	"layla/pkg/touch"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/tanema/gween"
	"github.com/yohamta/donburi"
)

type PauseScreenAction int

const (
	PauseScreenExit PauseScreenAction = iota
	PauseScreenResume
	PauseScreenRestart
)

type PauseScreenData struct {
	BgOpacity   float64
	OffsetY     float64
	BgTween     *gween.Tween
	OffsetTween *gween.Tween

	TouchIDs []ebiten.TouchID

	Selected PauseScreenAction
}

var PauseScreenBgImage *ebiten.Image

func init() {
	PauseScreenBgImage = ebiten.NewImage(config.Width*2, config.Height*2)
	PauseScreenBgImage.Fill(color.RGBA{24, 20, 37, 255})
}

var PauseScreen = donburi.NewComponentType[PauseScreenData]()

type PauseScreenImageKey int

const (
	PauseScreenResumeImage PauseScreenImageKey = iota
	PauseScreenResetImage
	PauseScreenExitImage
)

var PauseScreenImage map[PauseScreenImageKey]*ebiten.Image = map[PauseScreenImageKey]*ebiten.Image{}

var (
	PAUSE_UI_GAP             int
	PAUSE_UI_PIXEL_CELL_SIZE int
	PAUSE_UI_CELL_SIZE       int

	PAUSE_UI_EXIT_X int
	PAUSE_UI_EXIT_Y int

	PAUSE_UI_RESUME_X int
	PAUSE_UI_RESUME_Y int

	PAUSE_UI_RESTART_X int
	PAUSE_UI_RESTART_Y int
)

func UpdatePauseScreenLayout() {
	cellsQuantity := 3
	PAUSE_UI_GAP = 12
	PAUSE_UI_PIXEL_CELL_SIZE = 24
	PAUSE_UI_CELL_SIZE = PAUSE_UI_PIXEL_CELL_SIZE

	PAUSE_UI_EXIT_X = config.Width/2 - (PAUSE_UI_CELL_SIZE*cellsQuantity+PAUSE_UI_GAP*int(math.Max(0, float64(cellsQuantity)-1)))/2
	PAUSE_UI_EXIT_Y = config.Height - PAUSE_UI_CELL_SIZE - PAUSE_UI_GAP

	PAUSE_UI_RESUME_X = PAUSE_UI_EXIT_X + PAUSE_UI_PIXEL_CELL_SIZE + PAUSE_UI_GAP
	PAUSE_UI_RESUME_Y = config.Height - PAUSE_UI_CELL_SIZE - PAUSE_UI_GAP

	PAUSE_UI_RESTART_X = PAUSE_UI_RESUME_X + PAUSE_UI_PIXEL_CELL_SIZE + PAUSE_UI_GAP
	PAUSE_UI_RESTART_Y = config.Height - PAUSE_UI_CELL_SIZE - PAUSE_UI_GAP
}

func init() {
	PauseScreenImage[PauseScreenResumeImage] = InputUiImage.SubImage(image.Rectangle{
		Min: image.Point{
			X: INPUT_UI_CELL_SIZE,
			Y: INPUT_UI_PIXEL_CELL_SIZE * 3,
		},
		Max: image.Point{
			X: INPUT_UI_PIXEL_CELL_SIZE * 2, Y: INPUT_UI_PIXEL_CELL_SIZE * 4,
		},
	}).(*ebiten.Image)
	PauseScreenImage[PauseScreenResetImage] = InputUiImage.SubImage(image.Rectangle{
		Min: image.Point{
			X: INPUT_UI_CELL_SIZE * 2,
			Y: INPUT_UI_PIXEL_CELL_SIZE * 3,
		},
		Max: image.Point{
			X: INPUT_UI_PIXEL_CELL_SIZE * 3, Y: INPUT_UI_PIXEL_CELL_SIZE * 4,
		},
	}).(*ebiten.Image)
	PauseScreenImage[PauseScreenExitImage] = InputUiImage.SubImage(image.Rectangle{
		Min: image.Point{
			X: INPUT_UI_CELL_SIZE * 3,
			Y: INPUT_UI_PIXEL_CELL_SIZE * 3,
		},
		Max: image.Point{
			X: INPUT_UI_PIXEL_CELL_SIZE * 4, Y: INPUT_UI_PIXEL_CELL_SIZE * 4,
		},
	}).(*ebiten.Image)
}

func (i *PauseScreenData) IsRequestResume() bool {
	i.TouchIDs = inpututil.AppendJustPressedTouchIDs(i.TouchIDs[:0])
	for _, id := range i.TouchIDs {
		tx, ty := touch.TouchPos(id)
		if tx >= PAUSE_UI_RESUME_X && ty >= PAUSE_UI_RESUME_Y && tx <= PAUSE_UI_RESUME_X+PAUSE_UI_CELL_SIZE && ty <= PAUSE_UI_RESUME_Y+PAUSE_UI_CELL_SIZE {
			return true
		}
	}

	return input.Handler.ActionIsJustPressed(input.ActionExit) || (input.Handler.ActionIsJustPressed(input.ActionSelect) && i.Selected == PauseScreenResume)
}

func (i *PauseScreenData) IsRequestRestart() bool {
	i.TouchIDs = inpututil.AppendJustPressedTouchIDs(i.TouchIDs[:0])
	for _, id := range i.TouchIDs {
		tx, ty := touch.TouchPos(id)
		if tx >= PAUSE_UI_RESTART_X && ty >= PAUSE_UI_RESTART_Y && tx <= PAUSE_UI_RESTART_X+PAUSE_UI_CELL_SIZE && ty <= PAUSE_UI_RESTART_Y+PAUSE_UI_CELL_SIZE {
			return true
		}
	}

	return input.Handler.ActionIsJustPressed(input.ActionSelect) && i.Selected == PauseScreenRestart
}

func (i *PauseScreenData) IsRequestExit() bool {
	i.TouchIDs = inpututil.AppendJustPressedTouchIDs(i.TouchIDs[:0])
	for _, id := range i.TouchIDs {
		tx, ty := touch.TouchPos(id)
		if tx >= PAUSE_UI_EXIT_X && ty >= PAUSE_UI_EXIT_Y && tx <= PAUSE_UI_EXIT_X+PAUSE_UI_CELL_SIZE && ty <= PAUSE_UI_EXIT_Y+PAUSE_UI_CELL_SIZE {
			return true
		}
	}

	return input.Handler.ActionIsJustPressed(input.ActionSelect) && i.Selected == PauseScreenExit
}
