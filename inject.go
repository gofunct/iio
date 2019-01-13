//+build wireinject

package iio

import (
	"github.com/google/wire"
	"github.com/mattn/go-colorable"
	"io"
	"os"
)

var Provider = wire.NewSet(
	NewStdIO,
)

var FlexProvider = wire.NewSet(
	New,
)

func Inject() (*Service, error) {
	wire.Build(Provider)
	return &Service{}, nil
}

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
