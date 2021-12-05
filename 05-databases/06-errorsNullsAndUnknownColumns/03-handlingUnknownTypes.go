package _020506_errorsNullsAndUnknownColumns_

import (
	"database/sql"
	"fmt"
	"log"

	dbi "go-advanced/05-databases/00-shared"
)

func DemoHandlingUnknownTypes() {
	fmt.Println("\n\n020506 Databases: Errors, Nulls, and Unknown Columns — 03 Handling Unknowns")

	db, err := sql.Open("mysql", dbi.DSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM `employees`")
	if err != nil {
		log.Fatal(nil)
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	getColumns(cols, rows)
}

func getColumns(cols []string, rs *sql.Rows) {
	var tMap map[string]string
	fmt.Println(tMap)
	typesMap := make(map[string]string)
	for _, col := range cols {
		typesMap[col] = ""
	}

	colTypes, err := rs.ColumnTypes()
	if err != nil {
		log.Fatal(err)
	}

	typeSlice := make([]string, len(cols))
	for i, s := range colTypes {
		typeSlice[i] = s.DatabaseTypeName()
	}

	for i, c := range cols {
		typesMap[c] = typeSlice[i]
		fmt.Printf("column named %v is of type %v\n", c, typesMap[c])
	}
}
