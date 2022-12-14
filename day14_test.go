package main

import "testing"

func TestDay14_1(t *testing.T) {
	t.Parallel()
	input := readTestDayN(14)
	expected := "24"
	actual := day14_1(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
