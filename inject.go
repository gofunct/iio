//+build wireinject

package iio

import "github.com/google/wire"

func NewIO() (*Service, error) {
	wire.Build(Set)
	return &Service{}, nil
}
