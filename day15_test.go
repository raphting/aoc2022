package main

import "testing"

func TestDay15_1(t *testing.T) {
	t.Parallel()
	input := readTestDayN(15)
	expected := "26"
	actual := day15_1(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
