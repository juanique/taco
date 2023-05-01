package taco

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Text struct {
	PhysicalEntity
	Message string
	Color   Color

	font *ttf.Font
}

type TextOpts struct {
	FontFilename string
	Color        Color
	Size         int
}

func NewText(message string, opts TextOpts) Text {
	if opts.FontFilename == "" {
		opts.FontFilename = "default_text.ttf"
	}

	if opts.Color == (Color{}) {
		opts.Color = Black
	}

	if opts.Size == 0 {
		opts.Size = 12
	}

	font, err := ttf.OpenFont(opts.FontFilename, opts.Size)
	if err != nil {
		panic("Could not load font: " + err.Error())
	}

	return Text{
		Message: message,
		font:    font,
	}
}

func (t *Text) Destroy() {
	t.font.Close()
}

func (t *Text) Draw(renderer *sdl.Renderer) {
	surface, err := t.font.RenderUTF8Solid(
		t.Message,
		sdl.Color{R: t.Color.R, G: t.Color.G, B: t.Color.B, A: t.Color.Alpha},
	)

	if err != nil {
		panic("Could not render font.")
	}
	defer surface.Free()

	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		panic("Could not create texture for text.")
	}
	defer texture.Destroy()

	position := sdl.Rect{X: t.Pos.X, Y: t.Pos.Y, W: surface.W, H: surface.H}
	renderer.Copy(texture, nil, &position)
}
