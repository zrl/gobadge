package main

import (
	"image/color"
	"time"
)

const (
	WIDTH  = 160
	HEIGHT = 128
)

const (
	logoDisplayTime = 10 * time.Second
)

var rainbow []color.RGBA
var pressed uint8
var quit bool

func Badge() {
	quit = false
	display.FillScreen(colors[BLACK])

	rainbow = make([]color.RGBA, 256)
	for i := 0; i < 256; i++ {
		rainbow[i] = getRainbowRGB(uint8(i))
	}

	for {
		logoAnimated()
		if quit {
			break
		}
	}
}

func logoAnimated() {
	for _, frame := range logoRGBA_Animated {
		display.FillRectangleWithBuffer(0, 0, 100, 100, frame)
		time.Sleep(logoDisplayTime / 100)
	}
}
