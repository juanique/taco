package taco

import (
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

type Entity struct {
	Pos Vector2
}

func (ent *Entity) Move(x int32, y int32) {
	ent.Pos.X += x
	ent.Pos.Y += y
}

type Rect struct {
	Entity
	H     int32
	W     int32
	Color Color
}

func (ent *Rect) Draw(renderer *sdl.Renderer) {
	rect := sdl.Rect{X: ent.Pos.X, Y: ent.Pos.Y, W: ent.W, H: ent.H}
	renderer.SetDrawColor(ent.Color.R, ent.Color.G, ent.Color.B, ent.Color.Alpha)
	renderer.DrawRect(&rect)
}

func NewRect(scene *Scene) Rect {
	r := Rect{}
	r.Pos.X = scene.W / 2
	r.Pos.Y = scene.H / 2
	r.H = 10
	r.W = 10
	r.Color = Black
	return r
}
