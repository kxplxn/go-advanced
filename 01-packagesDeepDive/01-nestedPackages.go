package _01_packagesDeepDive_

import (
	"fmt"
	"go-advanced/01-packagesDeepDive/math"
	"go-advanced/01-packagesDeepDive/math/stats"
)

func NestedPackages() {
	fmt.Println("\nRunning `packagesDeepDive.NestedPackages`...")
	vals := []float64{1, 10, 100, 1000}
	fmt.Println("math.Add(10, 20):", math.Add(10, 20))
	fmt.Println("stats.Avg(vals):", stats.Avg(vals))
}
