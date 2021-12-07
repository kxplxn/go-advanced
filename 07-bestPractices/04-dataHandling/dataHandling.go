package _020704_dataHandling_

import (
	"fmt"
	"go-advanced/07-bestPractices/04-dataHandling/data"
	"log"
)

const (
	dbUser = "root"
	dbPass = "P@ssw0rd!"
	dbAddr = "127.0.0.1"
	dbName = "company"
)

func Demo() {
	fmt.Println("\n020704 Best Practices: Data Handling")

	err := data.InitDB(fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPass, dbAddr, dbName))
	if err != nil {
		log.Fatal(err)
	}

	emps, err := data.GetEmployees()
	if err != nil {
		log.Fatal(err)
	}
	for _, row := range emps {
		fmt.Printf("%+v\n", row)
	}
	fmt.Printf("We have %v people in our DB.", len(emps))
}
