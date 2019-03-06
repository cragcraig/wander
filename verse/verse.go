package verse

import (
	"github.com/cragcraig/wander/ascii"
)

type VerseView interface {
	Render(x, y int, b *ascii.Buffer)
}

type PosRenderable struct {
	X, Y int
	R    ascii.Renderable
}

type SimpleVerse struct {
	Renderables []PosRenderable
}

func (v *SimpleVerse) Render(x, y int, b *ascii.Buffer) {
	for i := range v.Renderables {
		r := v.Renderables[i]
		r.R.Render(r.X-x, r.Y-y, b)
	}
}
