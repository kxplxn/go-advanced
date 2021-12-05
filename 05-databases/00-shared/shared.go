package _020500_shared_

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
)

const (
	User     = "root"
	Password = "P@ssw0rd!"
	Address  = "127.0.0.1"
	Dbname   = "company"
)

var DSN = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", User, Password, Address, Dbname)

type Employee struct {
	Id   int
	Name string
	City string
	Dept string
}

func GetEmployees(db *sql.DB) (es []Employee, err error) {
	rs, err := db.Query("SELECT * FROM `employees`")
	fmt.Println("type of rs:", reflect.TypeOf(rs))
	if err != nil {
		return es, err
	}
	defer func(rs *sql.Rows) {
		if err := rs.Close(); err != nil {
			log.Printf("error occured while closing row set")
		}
	}(rs)

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
