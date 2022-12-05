package main

import "testing"

func TestDay5_1(t *testing.T) {
	t.Parallel()
	input := readTestDayN(5)
	expected := "CMZ"
	actual := day5_1(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestDay5_2(t *testing.T) {
	t.Parallel()
	input := readTestDayN(5)
	expected := "MCD"
	actual := day5_2(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
