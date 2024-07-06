package main

import (
	"context"
	"machine"
	"machine/usb"

	//keyboard "github.com/sago35/tinygo-keyboard"
	//"github.com/sago35/tinygo-keyboard/keycodes/jp"
	keyboard "github.com/xcd0/tinygo-keyboard"
	"github.com/xcd0/tinygo-keyboard/keycodes/jp"
)

func main() {
	usb.Product = "test-0.0.1"
	d := keyboard.New()

	col := []machine.Pin{machine.GPIO0, machine.GPIO1}
	row := []machine.Pin{machine.GPIO2, machine.GPIO3}

	d.AddMatrixKeyboard(col, row, [][]keyboard.Keycode{
		{
			jp.KeyA, jp.KeyB,
			jp.KeyC, jp.KeyD,
		},
	})
	d.Loop(context.Background())
}
