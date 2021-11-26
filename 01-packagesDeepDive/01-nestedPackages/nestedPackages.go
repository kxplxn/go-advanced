package _020101_nestedPackages

import (
	"fmt"
	"go-advanced/01-packagesDeepDive/01-nestedPackages/math"
	"go-advanced/01-packagesDeepDive/01-nestedPackages/math/stats"
)

func Demo() {
	fmt.Println("\n020101 Packages Deep Dive: Nested Packages")
	vals := []float64{1, 10, 100, 1000}
	fmt.Println("math.Add(10, 20):", math.Add(10, 20))
	fmt.Println("stats.Avg(vals):", stats.Avg(vals))
}
