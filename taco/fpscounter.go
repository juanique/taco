package taco

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type FPSCounter struct {
	frames int64
	timer  Timer
	fps    float64
	period time.Duration
	text   Text
}

func NewFPSCounter() FPSCounter {
	return FPSCounter{
		text: NewText("0 FPS", TextOpts{}),
	}
}

func (fpsCounter *FPSCounter) Destroy() {
	fpsCounter.text.Destroy()
}

func (fpsCounter *FPSCounter) Update(state *WorldState) {
	if !fpsCounter.timer.Started {
		fpsCounter.timer.Start()
	}

	if fpsCounter.period == 0 {
		fpsCounter.period = time.Millisecond * 500
	}

	fpsCounter.frames += 1

	if fpsCounter.timer.GetTicks() > fpsCounter.period {
		fpsCounter.fps = float64(fpsCounter.frames) / fpsCounter.timer.GetTicks().Seconds()
		fpsCounter.timer.Reset()
		fpsCounter.frames = 0
	}
}

func (fpsCounter *FPSCounter) Draw(renderer *sdl.Renderer) {
	fpsCounter.text.Message = fmt.Sprintf("%.0f FPS", fpsCounter.fps)
	fpsCounter.text.Draw(renderer)
}
