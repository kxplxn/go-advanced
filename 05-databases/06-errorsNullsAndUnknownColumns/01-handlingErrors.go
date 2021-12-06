package _020506_errorsNullsAndUnknownColumns_

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	dbi "go-advanced/05-databases/00-shared"
)

func DemoHandlingErrors() {
	fmt.Println("\n020506 Databases: Errors, Nulls, and Unknown Columns — 01 Handling Errors")

	db, err := sql.Open("mysql", dbi.DSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	var name, city, dept string
	id := 999999
	err = db.QueryRow("SELECT name, city, dept FROM employees WHERE id = ?", id).Scan(&name, &city, &dept)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no rows found for id", id)
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("name: %s, city: %s, dept: %s\n", name, city, dept)
	}

	rs, err := dbi.GetEmployees(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("List of all employees:")
	for _, emp := range rs {
		fmt.Printf("%+v, ", emp)
	}
}
