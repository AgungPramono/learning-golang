package golang_databases

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	//buat koneksi ke database
	db, err := sql.Open("mysql", "root:admin123@tcp(localhost:3306)/godb")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
