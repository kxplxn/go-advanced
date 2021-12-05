package _020502_dataRetrieval_

import (
	"database/sql"
	"fmt"
	dbi "go-advanced/05-databases/00-shared"
	"log"
	"reflect"
)

type Employee struct {
	Id   int
	Name string
	City string
	Dept string
}

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

	es, err := getEmployees(db)
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range es {
		fmt.Printf("%+v\n", e)
	}

	fmt.Println("We have", len(es), "employees in our DB!")
}

func getEmployees(db *sql.DB) (es []Employee, err error) {
	rs, err := db.Query("SELECT * FROM `employees`")
	fmt.Println("type of rs:", reflect.TypeOf(rs))
	if err != nil {
		return es, err
	}
	defer rs.Close()

	var e Employee
	for rs.Next() {
		err := rs.Scan(&e.Id, &e.Name, &e.City, &e.Dept)
		if err != nil {
			return es, err
		}
		es = append(es, e)
	}

	fmt.Println("type of es:", reflect.TypeOf(es))
	return es, nil
}
