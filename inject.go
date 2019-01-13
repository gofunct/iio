//+build wireinject

package iio

import "github.com/google/wire"

func Inject() (*Service, error) {
	wire.Build(Set)
	return &Service{}, nil
}
