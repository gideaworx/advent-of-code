package colors

import (
	"errors"
	"io"
)

// Colorizer provides an interface to send go template strings and return strings that will be rendered in color
// for different output mechanisms (for example, HTML or an ANSI terminal). Strings are of the format
//
//	{{ colorize "modifier1;modifier2;..." "text to colorize" }}
//
// Currently supported modifiers are
//   - bold
//   - faint
//   - italic
//   - strikethrough
//   - underline
//   - black
//   - red
//   - green
//   - yellow
//   - blue
//   - magenta
//   - cyan
//   - white
//
// All 8 colors have "bright-", "bg-", and "bright-bg-" variants to use bright versions of colors or to affect the background
// color of the text. The function name "colorizer" is configurable
type Colorizer interface {
	Format(string) string
	setCustomFunction(string)
}

// ColorizerOption is used to affect the instantiation of a new Colorizer
type ColorizerOption func(Colorizer)

// CustomFunctionName will change the name of the function from the default "colorizer"
func CustomFunctionName(name string) ColorizerOption {
	return func(c Colorizer) {
		c.setCustomFunction(name)
	}
}

type colorWriter struct {
	out io.Writer
	c   Colorizer
}

// NewColorWriter will create a new io.Writer instance that colorizes text sent to it before sending
// it on to the target writer
func NewColorWriter(out io.Writer, c Colorizer) io.Writer {
	return &colorWriter{
		out,
		c,
	}
}

func (c *colorWriter) Write(b []byte) (int, error) {
	if c.out == nil || c.c == nil {
		return 0, errors.New("neither the out writer nor colorizer can be nil")
	}

	s := string(b)
	colorized := c.c.Format(s)

	return c.out.Write([]byte(colorized))
}
