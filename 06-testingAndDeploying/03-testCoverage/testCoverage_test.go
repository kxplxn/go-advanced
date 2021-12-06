package _020502_unitTesting_

import "testing"

func TestAdd(t *testing.T) {
	x, y, expected := 10, 20, 30
	actual := Add(x, y)
	if expected != actual {
		t.Errorf("expected: %v, actual: %v\n", expected, actual)
	}
}

// table-driven testing
func TestSubTable(t *testing.T) {
	data := []struct {
		x        int
		y        int
		expected int
	}{
		{x: 2, y: 1, expected: 1},
		{x: 2, y: 2, expected: 0},
		{x: 2, y: 3, expected: -1},
	}

	for _, val := range data {
		got := Sub(val.x, val.y)
		if val.expected != got {
			t.Errorf("failed. expected: %v, got: %v\n", val.expected, got)
		}
	}
}

func TestAddTable(t *testing.T) {
	data := []struct {
		x        int
		y        int
		expected int
	}{
		{x: 2, y: 1, expected: 3},
		{x: 2, y: 2, expected: 4},
		{x: 2, y: 3, expected: 5},
	}

	for _, val := range data {
		got := Add(val.x, val.y)
		if val.expected != got {
			t.Errorf("failed. expected: %v, got: %v\n", val.expected, got)
		}
	}
}
