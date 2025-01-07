//go:build wireinject
// +build wireinject

//command generate dependecy injection: wire gen [namapackage]

package simple

import (
	"github.com/google/wire"
	"io"
	"os"
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

//func InitializeHelloService() *HelloService {
//	wire.Build(NewHelloService, NewSayHelloImpl)
//	return nil
//}

var helloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

func InitializeHelloService() *HelloService {
	wire.Build(helloSet, NewHelloService)
	return nil
}

func InitializeFooBar() *FooBar {
	//wire.Build(NewFoo, NewBar, wire.Struct(new(FooBar), "Foo", "Bar"))
	wire.Build(NewFoo, NewBar, wire.Struct(new(FooBar), "*"))
	return nil
}

var fooValue = &Foo{}
var barValue = &Bar{}

func IntializeFooBarUsingValue() *FooBar {
	wire.Build(wire.Value(fooValue), wire.Value(barValue), wire.Struct(new(FooBar), "*"))
	return nil
}

func InitializedReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

func InitializedConfiguration() *Configuration {
	wire.Build(
		NewApplication,
		wire.FieldsOf(new(*Application), "Configuration"),
	)
	return nil
}
