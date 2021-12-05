package _020502_dataRetrieval_

import (
	"database/sql"
	"fmt"
	dbi "go-advanced/05-databases/00-shared"
	"log"
)

func MultipleRowsDemo() {
	fmt.Println("\n020502 Databases: Data Retrieval — 02 Multiple Rows")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbi.User, dbi.Password, dbi.Address, dbi.Dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	es, err := dbi.GetEmployees(db)
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range es {
		fmt.Printf("%+v\n", e)
	}

	fmt.Println("We have", len(es), "employees in our DB!")
}
