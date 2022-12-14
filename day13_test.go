package main

import "testing"

func TestDay13_1(t *testing.T) {
	t.Parallel()
	input := readTestDayN(13)
	expected := "13"
	actual := day13_1(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
