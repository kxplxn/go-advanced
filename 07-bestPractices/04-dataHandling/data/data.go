package data

import (
	"database/sql"
	"log"
)

type Employee struct {
	ID   int
	Name sql.NullString
	City sql.NullString
	Dept sql.NullString
}

var db *sql.DB

func InitDB(dataSourceName string) error {
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return db.Ping()
}

func GetEmployees() (emps []Employee, err error) {
	rows, err := db.Query("SELECT * FROM  `employees`")
	if err != nil {
		log.Fatal(err)
	}

	var e Employee
	for rows.Next() {
		err = rows.Scan(&e.ID, &e.Name, &e.City, &e.Dept)
		if err != nil {
			return emps, err
		}
		emps = append(emps, e)
	}

	return emps, nil
}
