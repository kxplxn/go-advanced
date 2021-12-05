package _020500_shared_

import "fmt"

const (
	User     = "root"
	Password = "P@ssw0rd!"
	Address  = "127.0.0.1"
	Dbname   = "company"
)

var DSN = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", User, Password, Address, Dbname)
