package _020504_dataManipulation_

import (
	"database/sql"
	"fmt"
	"log"

	dbi "go-advanced/05-databases/00-shared"
)

func Demo() {
	fmt.Println("\n020504 Databases: Data Manipulation")

	db, err := sql.Open("mysql", dbi.DSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	demoInsert(db)
	demoUpdate(db)
	demoDelete(db)
}

func demoInsert(db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO employees(name) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec("kim")
	if err != nil {
		log.Fatal(err)
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted %d rows successfuly with ID: %d\n", rows, lastID)
}

func demoUpdate(db *sql.DB) {
	name, id := "winston", 1010
	stmt, err := db.Prepare("UPDATE employees SET name = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(name, id)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Updated %d row successfuly: set name = %s where id = %d", rows, name, id)
}

func demoDelete(db *sql.DB) {
	id := 1000
	stmt, err := db.Prepare("DELETE FROM employees WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Deleted %d rows successfuly: where id = %d", rows, id)
}
