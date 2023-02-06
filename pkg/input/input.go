package input

import (
	input "github.com/quasilyte/ebitengine-input"
)

const (
	ActionUnknown input.Action = iota
	ActionLeft
	ActionRight
	ActionUp
	ActionDown
	ActionJump
	ActionSelect
	ActionExit
)

var InputSystem input.System
var Handler *input.Handler

var SimulatedExit = input.SimulatedEvent{
	Key: input.KeyEscape,
}

func init() {
	InputSystem.Init(input.SystemConfig{
		DevicesEnabled: input.KeyboardDevice | input.GamepadDevice,
	})
	keymap := input.Keymap{
		ActionLeft:  {input.KeyGamepadLeft, input.KeyGamepadLStickLeft, input.KeyA, input.KeyLeft},
		ActionRight: {input.KeyGamepadRight, input.KeyGamepadLStickRight, input.KeyD, input.KeyRight},
		ActionUp:    {input.KeyGamepadUp, input.KeyGamepadLStickUp, input.KeyW, input.KeyUp},
		ActionDown:  {input.KeyGamepadDown, input.KeyGamepadLStickDown, input.KeyS, input.KeyDown},

		ActionJump:   {input.KeyGamepadA, input.KeyGamepadB, input.KeySpace, input.KeyW, input.KeyUp},
		ActionExit:   {input.KeyGamepadStart, input.KeyEscape},
		ActionSelect: {input.KeyGamepadA, input.KeyEnter, input.KeySpace},
	}
	Handler = InputSystem.NewHandler(0, keymap)
}
