package iio

import (
	"github.com/mattn/go-colorable"
	"io"
	"os"
	"runtime"
)

// IO contains an input reader, an output writer and an error writer.
type IOService interface {
	In() io.Reader
	Out() io.Writer
	Err() io.Writer
}

// IO contains an input reader, an output writer and an error writer.
type IO struct {
	In  io.Reader
	Out io.Writer
	Err io.Writer
}

// IOContainer is a basic implementation of the IO interface.
type IOContainer struct {
	InR  io.Reader
	OutW io.Writer
	ErrW io.Writer
}

func (i *IOContainer) In() io.Reader  { return i.InR }
func (i *IOContainer) Out() io.Writer { return i.OutW }
func (i *IOContainer) Err() io.Writer { return i.ErrW }

// Stdio returns a standard IO object.
func Stdio() IOService {
	io := &IOContainer{
		InR:  os.Stdin,
		OutW: os.Stdout,
		ErrW: os.Stderr,
	}
	if runtime.GOOS == "windows" {
		io.OutW = colorable.NewColorableStdout()
		io.ErrW = colorable.NewColorableStderr()
	}
	return io
}

// DefaultIO returns a standard IO object.
func DefaultIO() *IO {
	io := &IO{
		In:  os.Stdin,
		Out: os.Stdout,
		Err: os.Stderr,
	}
	if runtime.GOOS == "windows" {
		io.Out = colorable.NewColorableStdout()
		io.Err = colorable.NewColorableStderr()
	}
	return io
}

var closers []func()

// Close closes cli utilities.
func Close() {
	for _, f := range closers {
		f()
	}
}

func AddCloseFunc(f func()) {
	closers = append(closers, f)
}
