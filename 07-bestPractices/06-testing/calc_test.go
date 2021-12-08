package _020706_testing__test

import (
	"testing"

	calc "go-advanced/07-bestPractices/06-testing"
)

func TestAddTable(t *testing.T) {
	cases := []struct {
		a        int
		b        int
		expected int
	}{
		{a: 1, b: 2, expected: 3},
		{a: 2, b: 2, expected: 4},
		{a: 2, b: 3, expected: 5},
	}
	for _, c := range cases {
		if got := calc.Add(c.a, c.b); got != c.expected {
			t.Errorf("ran: calc.Add(%d, %d), exp: %d, got: %d", c.a, c.b, c.expected, got)
		}
	}
}

func TestSubTable(t *testing.T) {
	cases := []struct {
		a        int
		b        int
		expected int
	}{
		{a: 2, b: 1, expected: 1},
		{a: 1, b: 1, expected: 0},
		{a: 1, b: 2, expected: -1},
	}
	for _, c := range cases {
		if got := calc.Sub(c.a, c.b); got != c.expected {
			t.Errorf("ran: calc.Sub(%d, %d), expected: %d, got: %d", c.a, c.b, c.expected, got)
		}
	}
}
