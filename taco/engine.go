package taco

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const screenTicksPerFrame = time.Microsecond * 6944

type Drawable interface {
	Draw(*sdl.Renderer)
}

type Updateable interface {
	Update(*WorldState)
}

type InputState struct {
	keyboardState []uint8
}

func (s InputState) KeyPressed(keyCode uint8) bool {
	return s.keyboardState[keyCode] == 1
}

func (s *InputState) Update(state []uint8) {
	s.keyboardState = state
}

type WorldState struct {
	InputState InputState
}

type Engine struct {
	renderer *sdl.Renderer
	scene    Scene

	// Cap to ~144fps
	capTimer Timer

	// Renderable entities
	drawableEntities []Drawable

	// Objects that are updated each frame
	updatebleEntities []Updateable

	Stopped    bool
	WorldState WorldState
}

func NewEngine(renderer *sdl.Renderer, scene Scene) Engine {
	return Engine{
		renderer: renderer,
		scene:    scene,
	}
}

func (eng *Engine) AddDrawable(entity Drawable) {
	eng.drawableEntities = append(eng.drawableEntities, entity)
}

func (eng *Engine) AddUpdateable(entity Updateable) {
	eng.updatebleEntities = append(eng.updatebleEntities, entity)
}

func (eng *Engine) Update() {
	eng.capTimer.Start()
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			eng.Stopped = true
			return
		}
	}

	eng.WorldState.InputState.Update(sdl.GetKeyboardState())
	eng.renderer.SetDrawColor(255, 200, 200, 255)
	eng.renderer.Clear()

	for _, entity := range eng.updatebleEntities {
		entity.Update(&eng.WorldState)
	}

	for _, entity := range eng.drawableEntities {
		entity.Draw(eng.renderer)
	}

	eng.renderer.Present()

	frameTicks := eng.capTimer.GetTicks()
	if frameTicks < screenTicksPerFrame {
		delay := screenTicksPerFrame - frameTicks
		sdl.Delay(uint32(delay.Milliseconds()))
	}

}

func (eng *Engine) Run() {
	for {
		eng.Update()
		if eng.Stopped {
			return
		}
	}
}
