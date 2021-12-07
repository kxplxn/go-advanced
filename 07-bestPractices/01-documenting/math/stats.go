package math

// Avg takes a slice of float64 values, computes the average, returns the result
func Avg(vals []float64) float64 {
	total := 0.0
	for _, val := range vals {
		total += val
	}
	return total / float64(len(vals))
}
