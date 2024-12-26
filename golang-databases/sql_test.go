package golang_databases

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	scriptSql := "INSERT INTO customer (id, name) VALUES ('03','joko')"
	//query exec: untuk yang tidak mengembalikan result
	//seperti perintah: insert,update,delete
	_, err := db.ExecContext(ctx, scriptSql)
	if err != nil {
		panic(err)
	}
	fmt.Println("sukses insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	scriptSql := "SELECT id,name FROM customer"
	//QueryContext: execute sql yang berupa query
	rows, err := db.QueryContext(ctx, scriptSql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	//iterasi data
	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("id: ", id)
		fmt.Println("name: ", name)
	}
}
