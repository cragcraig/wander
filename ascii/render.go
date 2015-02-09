package ascii

// Creates a vt100 (ANSI) Esc code.
func vt100(s string) []byte {
	r := make([]byte, 2+len(s))
	r = append(r, vt100Esc...)
	return append(r, []byte(s)...)
}

// vt100 codes
var (
	upperLeft       []byte = vt100("2J")
	clearScreen     []byte = vt100(";H")
	hideCursor      []byte = vt100("?25l")
	showCursor      []byte = vt100("?25h")
	clearFormatting []byte = vt100("0m")
	newLine         []byte = []byte{'\r', '\n'}
)

// Constants
var (
	vt100Esc       []byte   = []byte{0x1b, '['}
	renderPrefixes [][]byte = [][]byte{hideCursor, clearScreen, upperLeft}
	renderPrefix   []byte
)

// Lazy initializes the render prefix.
func getRenderPrefix() []byte {
	if renderPrefix == nil {
		for i := range renderPrefixes {
			renderPrefix = append(renderPrefix, renderPrefixes[i]...)
		}
	}
	return renderPrefix
}

// Renders a buffer of printable bytes to a vt100 terminal.
func Render(w int, h int, buf []byte) []byte {
	if len(buf) != w*h {
		panic("Length of render buffer was not width * hight!")
	}
	prefix := getRenderPrefix()
	out := make([]byte, (w+len(newLine))*h+len(prefix))
	out = append(out, prefix...)
	for l := 0; l < h; l++ {
		if l != 0 {
			out = append(out, newLine...)
		}
		p := w * l
		out = append(out, buf[p:p+w]...)
	}
	return out
}
