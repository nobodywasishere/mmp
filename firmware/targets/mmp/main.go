package main

import (
	"context"
	_ "embed"
	"machine"

	keyboard "github.com/sago35/tinygo-keyboard"
	"github.com/sago35/tinygo-keyboard/keycodes/jp"
)

func main() {
	d := keyboard.New()

	rowPins := []machine.Pin{
		machine.D3,
		machine.D2,
		machine.D1,
		machine.D0,
	}

	colPins := []machine.Pin{
		machine.D6,
		machine.D9,
		machine.D10,
	}

	d.AddMatrixKeyboard(colPins, rowPins, [][]keyboard.Keycode{
		{
			jp.KeyA, jp.KeyB, jp.KeyC, jp.KeyD,
			jp.KeyA, jp.KeyB, jp.KeyC, jp.KeyD,
			jp.KeyA, jp.KeyB, jp.KeyC, jp.KeyD,
		},
	})

	loadKeyboardDef()
	d.Loop(context.Background())
}
