package golang_databases

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	//buat koneksi ke database
	db, err := sql.Open("mysql", "root:admin123@tcp(localhost:3306)/godb")
	if err != nil {
		panic(err)
	}

	//close koneksi setelah selesai digunakan
	defer db.Close()

}
