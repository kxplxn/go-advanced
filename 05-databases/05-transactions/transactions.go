package _020505_transactions_

import (
	"database/sql"
	"fmt"
	dbi "go-advanced/05-databases/00-shared"
	"log"
	"strconv"
)

func Demo() {
	fmt.Println("\n020505 Databases: Transactions")

	db, err := sql.Open("mysql", dbi.DSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// ************** tx starts
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	for i := 0; i < 10; i++ {
		name := "user" + strconv.Itoa(i)
		_, err := tx.Exec("INSERT INTO `employees`(name) VALUES (?)", name)

		// simulate error in the transaction
		//if i == 5 {
		//	_, err = tx.Exec("INSERT INTO `employees`(name) VALUES (?)", name, "fail")
		//}

		if err != nil {
			tx.Rollback()
			fmt.Println("error occured while writing", name)
			return
		}
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}
