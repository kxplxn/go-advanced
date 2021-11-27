package _020107_modules

import (
	"fmt"
	m "github.com/rccsdevops/calc/math"
)

var env string

func init() {
	env = "demo"
}

func Demo() {
	fmt.Println("\n0201007 Packages Deep Dive: Modules")
	fmt.Println(env, "environment is loaded.")
	fmt.Println("m.Add(1, 1) is", m.Add(1, 1))
}
