package _020501_accessingDatabases_

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

	db "go-advanced/05-databases/00-shared"
)

func Demo() {
	fmt.Println("\n020501 Databases: Accessing Databases")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db.User, db.Password, db.Address, db.Dbname)
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
