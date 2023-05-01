package taco

import "github.com/veandco/go-sdl2/sdl"

type Game struct {
}

type PlayerController struct {
	entity IPhysicalEntity
}

func (p *PlayerController) Update(state *WorldState) {
	movement := Vector2{}
	if state.InputState.KeyPressed(sdl.SCANCODE_LEFT) {
		movement.X -= 1
	}
	if state.InputState.KeyPressed(sdl.SCANCODE_RIGHT) {
		movement.X += 1
	}
	if state.InputState.KeyPressed(sdl.SCANCODE_UP) {
		movement.Y -= 1
	}
	if state.InputState.KeyPressed(sdl.SCANCODE_DOWN) {
		movement.Y += 1
	}

	p.entity.Move(movement)
}

func (g *Game) Run(renderer *sdl.Renderer, scene Scene) {
	engine := NewEngine(renderer, scene)

	rect := NewRect(scene)
	engine.AddDrawable(&rect)

	fpsCounter := NewFPSCounter()
	engine.AddUpdateable(&fpsCounter)
	engine.AddDrawable(&fpsCounter)

	player := PlayerController{entity: &rect}
	engine.AddUpdateable(&player)

	engine.Run()
}
