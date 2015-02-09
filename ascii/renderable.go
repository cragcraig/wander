package ascii

type Renderable struct {
	bytes [][]byte
}

func asRenderable(s ...string) Renderable {
	r := make([][]byte, len(s))
	for i := range s {
		r[i] = []byte(s[i])
	}
	return Renderable{r}
}

func (r *Renderable) Render(x, y int, b *Buffer) {
	for iy := range r.bytes {
		if y+iy >= b.H || y+iy < 0 {
			continue
		}
		row := r.bytes[iy]
		for i := 0; i < len(row); i++ {
			if x+i >= b.W || x+i < 0 {
				continue
			}
			if row[i] != 0x00 {
				b.Set(x+i, y, row[i])
			}
		}
	}
}

var (
	TieFighter Renderable = asRenderable("|o|")
	PlayerShip Renderable = asRenderable("{^}")
	Asteroid   Renderable = asRenderable("{#}")
)
