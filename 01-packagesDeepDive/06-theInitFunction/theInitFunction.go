package _020106_initFunction

import (
	"fmt"
	_ "go-advanced/01-packagesDeepDive/06-theInitFunction/math"
	_ "go-advanced/01-packagesDeepDive/06-theInitFunction/math/stats"
)

var env string

func init() {
	fmt.Println("Secret key from package main is fetched.")
}

func init() {
	fmt.Println("Database is initialised and ready.")
}

func init() {
	// global variable env is configured
	env = "development"
}

func Demo() {
	fmt.Println("\n0201006 Packages Deep Dive: The init Function")
	fmt.Println(env, "environment is loaded.")
}
