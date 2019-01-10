package iio

import (
	"github.com/google/wire"
	"github.com/mattn/go-colorable"
	"io"
	"os"
)

func New(o, e io.Writer, i io.Reader) *Service {
	return &Service{
		InR:  i,
		OutW: o,
		ErrW: e,
	}
}

// Stdio returns a standard IO object.
func NewStdIO() *Service {
	io := &Service{
		InR:  os.Stdin,
		OutW: colorable.NewColorableStdout(),
		ErrW: colorable.NewColorableStderr(),
	}

	return io
}

var Set = wire.NewSet(
	NewStdIO,
)

var FlexSet = wire.NewSet(
	New,
)
