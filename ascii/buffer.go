package ascii

type Buffer struct {
	Buffer []byte
	W, H   int
}

func CreateBuffer(w, h int) *Buffer {
	b := Buffer{make([]byte, w*h), w, h}
	b.Clear()
	return &b
}

func (buf *Buffer) Get(x, y int) byte {
	return buf.Buffer[buf.W*y+x]
}

func (buf *Buffer) Set(x, y int, b byte) {
	buf.Buffer[buf.W*y+x] = b
}

func (buf *Buffer) Render() []byte {
	return Render(buf.W, buf.H, buf.Buffer)
}

func (buf *Buffer) Clear() {
	for i := range buf.Buffer {
		buf.Buffer[i] = ' '
	}
}
