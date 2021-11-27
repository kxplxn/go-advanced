// Package math implements simple math operations Add and Sub, the statistical
// value Avg, as well as the constant Pi.
package math

// Pi implements the constant Pi rounded to 5 decimals.
const Pi = 3.14159

// Add takes two int arguments, adds them, and returns the resulting int value.
func Add(x, y int) int {
	return x + y
}

// Sub takes two int arguments, substracts them, and returns the resulting int
// value.
func Sub(x, y int) int {
	return x - y
}
