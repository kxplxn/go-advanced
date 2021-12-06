package _020502_unitTesting_

import "testing"

func TestAdd(t *testing.T) {
	x, y, expected := 10, 20, 30
	actual := Add(x, y)
	if expected != actual {
		t.Errorf("expected: %v, actual: %v\n", expected, actual)
	}
}

func TestSub(t *testing.T) {
	x, y, expected := 20, 5, 15
	actual := Sub(x, y)
	if expected != actual {
		t.Errorf("expected: %v, actual: %v\n", expected, actual)
	}
}
