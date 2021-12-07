// Copyright 2021 Joe. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package math implements simple math operations Add, and Sub,
// the statistical value Avg, as well as the constant Pi.
package math

// Pi implements the constant Pi rounded to 5 decimals.
const Pi float64 = 3.14159

// Add takes 2 int arguments, adds them and returns the resulting int value.
func Add(a, b int) int {
	return a + b
}

// Sub takes 2 int arguments, substracts them and returns the resulting int value.
func Sub(a, b int) int {
	return a - b
}
