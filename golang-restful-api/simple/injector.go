//go:build wireinject
// +build wireinject

//command generate dependecy injection: wire gen [namapackage]

package simple

import (
	"github.com/google/wire"
)

func InitializeService() (*SimpleService, error) {
	wire.Build(
		NewSimpleRepository, NewSimpleService,
	)
	return nil, nil
}
