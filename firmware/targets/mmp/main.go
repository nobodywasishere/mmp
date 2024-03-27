package main

import (
	"context"
	_ "embed"
	"image/color"
	"machine"

	keyboard "github.com/sago35/tinygo-keyboard"
	"github.com/sago35/tinygo-keyboard/keycodes/jp"
	"tinygo.org/x/drivers/ws2812"
)

func main() {
	neo := machine.WS2812
	neo.Configure(machine.PinConfig{Mode: machine.PinOutput})

	wsLeds := [1]color.RGBA{}
	wsLeds[0] = color.RGBA{0x05, 0x00, 0x00, 0x05}

	ws := ws2812.New(neo)
	ws.WriteColors(wsLeds[:])

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
			jp.KeyA, jp.KeyB, jp.KeyC,
			jp.KeyD, jp.KeyE, jp.KeyF,
			jp.KeyG, jp.KeyH, jp.KeyI,
			jp.KeyJ, jp.KeyK, jp.KeyL,
		},
	})

	loadKeyboardDef()
	d.Loop(context.Background())
}
