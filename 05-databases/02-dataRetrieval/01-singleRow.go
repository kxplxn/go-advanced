package _020502_dataRetrieval_

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	dbi "go-advanced/05-databases/00-shared"
	"log"
)

func SingleRowDemo() {
	fmt.Println("\n020502 Databases: Data Retrieval — 01 Single Row")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbi.User, dbi.Password, dbi.Address, dbi.Dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// validate the DSN
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("connection success")
	}

	var name, city, dept string
	id := 1000
	err = db.QueryRow("SELECT `name`, `city`, `dept` FROM `employees` WHERE `id` = ?", id).Scan(&name, &city, &dept)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name: %s, city: %s, dept: %s\n", name, city, dept)
}
