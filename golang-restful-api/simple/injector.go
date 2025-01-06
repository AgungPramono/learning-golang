//go:build wireinject
// +build wireinject

//command generate dependecy injection: wire gen [namapackage]

package simple

import (
	"github.com/google/wire"
)

func InitializeService(isError bool) (*SimpleService, error) {
	wire.Build(
		NewSimpleRepository, NewSimpleService,
	)
	return nil, nil
}

func InitializeDatabaseRepository() *DatabaseRepository {
	wire.Build(
		NewDatabasePostgreSQL,
		NewDatabaseMysql,
		NewDatabaseRepository,
	)
	return nil
}

var FooSet = wire.NewSet(NewFooRepository, NewFooService)

var BarSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializeFooBarService() *FooBarService {
	wire.Build(FooSet, BarSet, NewFooBarService)
	return nil
}
