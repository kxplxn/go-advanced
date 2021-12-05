package _020506_errorsNullsAndUnknownColumns_

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"

	_ "github.com/go-sql-driver/mysql"

	dbi "go-advanced/05-databases/00-shared"
)

func DemoHandlingNulls() {
	fmt.Println("\n\n020506 Databases: Errors, Nulls, and Unknown Columns — 02 Handling Nulls")

	db, err := sql.Open("mysql", dbi.DSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// DB column name is nullable and contains a null value
	//var name, city, dept string
	var city, dept string
	var name sql.NullString
	fmt.Println("Type of name:", reflect.TypeOf(name))
	fmt.Println("Type of name.String:", reflect.TypeOf(name.String))
	id := 1060
	err = db.QueryRow("SELECT name, city, dept FROM employees WHERE id = ?", id).Scan(&name, &city, &dept)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no employee found for id", id)
		} else {
			log.Fatal(err)
		}
	} else {
		if name.Valid {
			fmt.Printf("name is %s\n", name.String)
		} else {
			fmt.Printf("name is Null {%T}\n", name.String)
		}
	}
}
