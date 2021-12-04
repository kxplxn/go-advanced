package _020501_accessingDatabases_

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

	dbi "go-advanced/05-databases/00-shared"
)

func Demo() {
	fmt.Println("\n020501 Databases: Accessing Databases")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbi.User, dbi.Password, dbi.Address, dbi.Dbname)
	fmt.Println("dsn:", dsn)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// db settings to consider
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// validate the DSN
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("connection success")
	}
}
