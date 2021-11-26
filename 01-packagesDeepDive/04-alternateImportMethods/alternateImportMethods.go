package _020104_alternateImportMethods

import (
	f "fmt"
	"math"
	. "os"
	_ "strings"

	m "go-advanced/01-packagesDeepDive/04-alternateImportMethods/math"
)

func Demo() {
	f.Println("\n020104 Packages Deep Dive: Alternate Import Methods")
	f.Println("m.Add(10, 20):", m.Add(10, 20))
	f.Println("math.Pi:", math.Pi)
	f.Fprintln(Stdout, "f.Fprintln(Stdout, \"Hello\"):", "Hello")
}
