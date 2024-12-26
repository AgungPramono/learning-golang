package golang_databases

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
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

func TestSqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	scriptSql := "SELECT id,name,email,balance,rating,birth_date,married,create_at FROM customer"
	rows, err := db.QueryContext(ctx, scriptSql)
	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate, createdAt time.Time
		var married bool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("======================================")
		fmt.Println("id: ", id)
		fmt.Println("name : ", name)
		if email.Valid {
			fmt.Println("email : ", email.String)
		}
		fmt.Println("balance :", balance)
		fmt.Println("rating:", rating)
		fmt.Println("birth_date :", birthDate)
		fmt.Println("married :", married)
		fmt.Println("create_at :", createdAt)
		fmt.Println("======================================")
	}
	defer rows.Close()
}
