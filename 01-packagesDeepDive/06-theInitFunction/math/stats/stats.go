package stats

import "fmt"

func Avg(vals []float64) float64 {
	total := 0.0
	for _, val := range vals {
		total += val
	}
	return total / float64(len(vals))
}

func init() {
	fmt.Println("Package stats is initialised.")
}
