package taco

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

var Black = Color{0, 0, 0, 255}

type Scene struct {
	H int32
	W int32
}

type Color struct {
	R     uint8
	G     uint8
	B     uint8
	Alpha uint8
}

type Vector2 struct {
	X int32
	Y int32
}

func (v *Vector2) Add(other Vector2) {
	v.X += other.X
	v.Y += other.Y
}

func (v Vector2) String() string {
	return fmt.Sprintf("Vector2{%d, %d}", v.X, v.Y)
}

type IPhysicalEntity interface {
	Move(Vector2)
}

type PhysicalEntity struct {
	Pos Vector2
}

func (ent *PhysicalEntity) Move(vector Vector2) {
	ent.Pos.Add(vector)
}

type Rect struct {
	PhysicalEntity
	H     int32
	W     int32
	Color Color
}

func (ent *Rect) Draw(renderer *sdl.Renderer) {
	rect := sdl.Rect{X: ent.Pos.X, Y: ent.Pos.Y, W: ent.W, H: ent.H}
	renderer.SetDrawColor(ent.Color.R, ent.Color.G, ent.Color.B, ent.Color.Alpha)
	renderer.DrawRect(&rect)
}

func NewRect(scene Scene) Rect {
	r := Rect{}
	r.Pos.X = scene.W / 2
	r.Pos.Y = scene.H / 2
	r.H = 10
	r.W = 10
	r.Color = Black
	return r
}

type FPSCounter struct {
	frames int64
	timer  Timer
}

func (fpsCounter *FPSCounter) Update(state *WorldState) {
	if !fpsCounter.timer.Started {
		fpsCounter.timer.Start()
	}

	fpsCounter.frames += 1

	if fpsCounter.timer.GetTicks() > 1*time.Second {
		fps := float64(fpsCounter.frames) / fpsCounter.timer.GetTicks().Seconds()
		fmt.Printf("%.2f FPS\n", fps)
		fpsCounter.timer.Reset()
		fpsCounter.frames = 0
	}
}
