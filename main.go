package main

import (
	"fmt"
	"time"

	"github.com/juanique/taco/taco"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth         = 600
	screenHeight        = 800
	screenTicksPerFrame = time.Microsecond * 6944
)

func main() {
	scene := taco.Scene{H: screenHeight, W: screenWidth}
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL:", err)
		return
	}

	window, err := sdl.CreateWindow(
		"Gaming in Go Episode 2",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("initializing window:", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing renderer:", err)
		return
	}
	defer renderer.Destroy()

	rect := taco.NewRect(&scene)
	frames := 1

	fpsTimer := taco.Timer{}
	capTimer := taco.Timer{}
	fpsTimer.Start()
	for {
		capTimer.Start()
		frames += 1
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		keys := sdl.GetKeyboardState()
		if keys[sdl.SCANCODE_LEFT] == 1 {
			rect.Move(-1, 0)
		}
		if keys[sdl.SCANCODE_RIGHT] == 1 {
			rect.Move(1, 0)
		}
		if keys[sdl.SCANCODE_UP] == 1 {
			rect.Move(0, -1)
		}
		if keys[sdl.SCANCODE_DOWN] == 1 {
			rect.Move(0, 1)
		}

		renderer.SetDrawColor(255, 200, 200, 255)
		renderer.Clear()
		rect.Draw(renderer)
		renderer.Present()
		frameTicks := capTimer.GetTicks()
		if frameTicks < screenTicksPerFrame {
			delay := screenTicksPerFrame - frameTicks
			sdl.Delay(uint32(delay.Milliseconds()))
		}

		if fpsTimer.GetTicks() > 1*time.Second {
			fps := float64(frames) / fpsTimer.GetTicks().Seconds()
			fmt.Printf("%.2f FPS\n", fps)
			fpsTimer.Reset()
			frames = 0
		}
	}
}
