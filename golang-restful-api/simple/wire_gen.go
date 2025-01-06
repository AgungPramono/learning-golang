// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package simple

// Injectors from injector.go:

func InitializeService(isError bool) (*SimpleService, error) {
	simpleRepository := NewSimpleRepository(isError)
	simpleService, err := NewSimpleService(simpleRepository)
	if err != nil {
		return nil, err
	}
	return simpleService, nil
}

func InitializeDatabaseRepository() *DatabaseRepository {
	databasePostgreSQL := NewDatabasePostgreSQL()
	databaseMysql := NewDatabaseMysql()
	databaseRepository := NewDatabaseRepository(databasePostgreSQL, databaseMysql)
	return databaseRepository
}
