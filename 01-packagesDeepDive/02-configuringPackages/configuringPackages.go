package _020102_configuringPackages

import (
	"fmt"
	"go-advanced/01-packagesDeepDive/02-configuringPackages/math"
)

func Demo() {
	fmt.Println("\n020102 Packages Deep Dive: Configuring Packages")
	vals := []float64{1, 10, 100, 1000}
	fmt.Println("math.Add(10, 20):", math.Add(10, 20))
	fmt.Println("stats.Avg(vals):", math.Avg(vals))
}
