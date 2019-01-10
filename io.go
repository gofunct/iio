package iio

import (
	"io"
)

// Service is a basic implementation of the IO interface.
type Service struct {
	InR  io.Reader
	OutW io.Writer
	ErrW io.Writer
}

func (i *Service) In() io.Reader  { return i.InR }
func (i *Service) Out() io.Writer { return i.OutW }
func (i *Service) Err() io.Writer { return i.ErrW }
