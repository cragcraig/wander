package ascii

type Renderable struct {
	bytes []byte
}

var RenderableWidth int = 3
var RenderableHeight int = 1

func asRenderable(s string) Renderable {
	if len(s) != RenderableWidth {
		panic("Renderable is not 3 chars")
	}
	return Renderable{[]byte(s)}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func (r *Renderable) Render(x, y int, b *Buffer) {
	if y >= b.H || y < 0 {
		return
	}
	for i := 0; i < len(r.bytes); i++ {
		if x+i >= b.W || x+i < 0 {
			continue
		}
		ch := r.bytes[i]
		if ch != ' ' {
			b.Set(x+i, y, ch)
		}
	}
}

var (
	TieFighter Renderable = asRenderable("|o|")
	PlayerShip Renderable = asRenderable("{^}")
	Asteroid   Renderable = asRenderable("{#}")
)
