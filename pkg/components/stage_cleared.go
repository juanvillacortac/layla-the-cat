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

type StageClearedAction int

const (
	StageClearedExit StageClearedAction = iota
	StageClearedResume
	StageClearedRestart
)

type StageClearedData struct {
	BgOpacity   float64
	OffsetY     float64
	BgTween     *gween.Tween
	OffsetTween *gween.Tween

	TouchIDs []ebiten.TouchID

	Selected StageClearedAction
}

var StageClearedBgImage *ebiten.Image

func init() {
	StageClearedBgImage = ebiten.NewImage(config.Width*2, config.Height*2)
	StageClearedBgImage.Fill(color.RGBA{24, 20, 37, 255})
}

var StageCleared = donburi.NewComponentType[StageClearedData]()

type StageClearedImgKey int

const (
	StageClearedResumeImage StageClearedImgKey = iota
	StageClearedResetImage
	StageClearedExitImage
)

var StageClearedImage map[StageClearedImgKey]*ebiten.Image = map[StageClearedImgKey]*ebiten.Image{}

var (
	STAGE_CLEARED_UI_GAP             int
	STAGE_CLEARED_UI_PIXEL_CELL_SIZE int
	STAGE_CLEARED_UI_CELL_SIZE       int

	STAGE_CLEARED_UI_EXIT_X int
	STAGE_CLEARED_UI_EXIT_Y int

	STAGE_CLEARED_UI_RESUME_X int
	STAGE_CLEARED_UI_RESUME_Y int

	STAGE_CLEARED_UI_RESTART_X int
	STAGE_CLEARED_UI_RESTART_Y int
)

func UpdateStageClearedLayout() {
	cellsQuantity := 3
	STAGE_CLEARED_UI_GAP = 12
	STAGE_CLEARED_UI_PIXEL_CELL_SIZE = 24
	STAGE_CLEARED_UI_CELL_SIZE = STAGE_CLEARED_UI_PIXEL_CELL_SIZE

	STAGE_CLEARED_UI_EXIT_X = config.Width/2 - (STAGE_CLEARED_UI_CELL_SIZE*cellsQuantity+STAGE_CLEARED_UI_GAP*int(math.Max(0, float64(cellsQuantity)-1)))/2
	STAGE_CLEARED_UI_EXIT_Y = config.Height - STAGE_CLEARED_UI_CELL_SIZE - STAGE_CLEARED_UI_GAP

	STAGE_CLEARED_UI_RESUME_X = STAGE_CLEARED_UI_EXIT_X + STAGE_CLEARED_UI_PIXEL_CELL_SIZE + STAGE_CLEARED_UI_GAP
	STAGE_CLEARED_UI_RESUME_Y = config.Height - STAGE_CLEARED_UI_CELL_SIZE - STAGE_CLEARED_UI_GAP

	STAGE_CLEARED_UI_RESTART_X = STAGE_CLEARED_UI_RESUME_X + STAGE_CLEARED_UI_PIXEL_CELL_SIZE + STAGE_CLEARED_UI_GAP
	STAGE_CLEARED_UI_RESTART_Y = config.Height - STAGE_CLEARED_UI_CELL_SIZE - STAGE_CLEARED_UI_GAP
}

func init() {
	StageClearedImage[StageClearedResumeImage] = InputUiImage.SubImage(image.Rectangle{
		Min: image.Point{
			X: INPUT_UI_CELL_SIZE,
			Y: INPUT_UI_PIXEL_CELL_SIZE * 3,
		},
		Max: image.Point{
			X: INPUT_UI_PIXEL_CELL_SIZE * 2, Y: INPUT_UI_PIXEL_CELL_SIZE * 4,
		},
	}).(*ebiten.Image)
	StageClearedImage[StageClearedResetImage] = InputUiImage.SubImage(image.Rectangle{
		Min: image.Point{
			X: INPUT_UI_CELL_SIZE * 2,
			Y: INPUT_UI_PIXEL_CELL_SIZE * 3,
		},
		Max: image.Point{
			X: INPUT_UI_PIXEL_CELL_SIZE * 3, Y: INPUT_UI_PIXEL_CELL_SIZE * 4,
		},
	}).(*ebiten.Image)
	StageClearedImage[StageClearedExitImage] = InputUiImage.SubImage(image.Rectangle{
		Min: image.Point{
			X: INPUT_UI_CELL_SIZE * 3,
			Y: INPUT_UI_PIXEL_CELL_SIZE * 3,
		},
		Max: image.Point{
			X: INPUT_UI_PIXEL_CELL_SIZE * 4, Y: INPUT_UI_PIXEL_CELL_SIZE * 4,
		},
	}).(*ebiten.Image)
}

func (i *StageClearedData) IsRequestResume() bool {
	i.TouchIDs = inpututil.AppendJustPressedTouchIDs(i.TouchIDs[:0])
	for _, id := range i.TouchIDs {
		tx, ty := touch.TouchPos(id)
		if tx >= STAGE_CLEARED_UI_RESUME_X && ty >= STAGE_CLEARED_UI_RESUME_Y && tx <= STAGE_CLEARED_UI_RESUME_X+STAGE_CLEARED_UI_CELL_SIZE && ty <= STAGE_CLEARED_UI_RESUME_Y+STAGE_CLEARED_UI_CELL_SIZE {
			return true
		}
	}

	return input.Handler.ActionIsJustPressed(input.ActionExit) || (input.Handler.ActionIsJustPressed(input.ActionSelect) && i.Selected == StageClearedResume)
}

func (i *StageClearedData) IsRequestRestart() bool {
	i.TouchIDs = inpututil.AppendJustPressedTouchIDs(i.TouchIDs[:0])
	for _, id := range i.TouchIDs {
		tx, ty := touch.TouchPos(id)
		if tx >= STAGE_CLEARED_UI_RESTART_X && ty >= STAGE_CLEARED_UI_RESTART_Y && tx <= STAGE_CLEARED_UI_RESTART_X+STAGE_CLEARED_UI_CELL_SIZE && ty <= STAGE_CLEARED_UI_RESTART_Y+STAGE_CLEARED_UI_CELL_SIZE {
			return true
		}
	}

	return input.Handler.ActionIsJustPressed(input.ActionSelect) && i.Selected == StageClearedRestart
}

func (i *StageClearedData) IsRequestExit() bool {
	i.TouchIDs = inpututil.AppendJustPressedTouchIDs(i.TouchIDs[:0])
	for _, id := range i.TouchIDs {
		tx, ty := touch.TouchPos(id)
		if tx >= STAGE_CLEARED_UI_EXIT_X && ty >= STAGE_CLEARED_UI_EXIT_Y && tx <= STAGE_CLEARED_UI_EXIT_X+STAGE_CLEARED_UI_CELL_SIZE && ty <= STAGE_CLEARED_UI_EXIT_Y+STAGE_CLEARED_UI_CELL_SIZE {
			return true
		}
	}

	return input.Handler.ActionIsJustPressed(input.ActionSelect) && i.Selected == StageClearedExit
}
