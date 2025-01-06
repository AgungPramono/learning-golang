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
