package _020503_preparedStatements_

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	dbi "go-advanced/05-databases/00-shared"
)

func Demo() {
	fmt.Println("\n020503 Databases: Prepared Statements")

	db, err := sql.Open("mysql", dbi.DSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("INSERT INTO employees(name) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec("andreas")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec("jan")
	if err != nil {
		log.Fatal(err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("id: %d, affected: %d\n", lastId, rowCnt)
}
