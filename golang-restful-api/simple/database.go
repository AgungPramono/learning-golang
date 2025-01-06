package simple

type Database struct {
	Name string
}

type DatabasePostgreSQL Database
type DatabaseMysql Database

func NewDatabaseMysql() *DatabaseMysql {
	return &DatabaseMysql{Name: "mysql"}
}

func NewDatabasePostgreSQL() *DatabasePostgreSQL {
	return &DatabasePostgreSQL{Name: "PostgreSQL"}
}

type DatabaseRepository struct {
	DatabasePostgreSQL *Database
	DatabaseMysql      *Database
}

func NewDatabaseRepository(postgreSQL *DatabasePostgreSQL, mysql *DatabaseMysql) *DatabaseRepository {
	return &DatabaseRepository{
		DatabasePostgreSQL: (*Database)(postgreSQL),
		DatabaseMysql:      (*Database)(mysql),
	}
}
