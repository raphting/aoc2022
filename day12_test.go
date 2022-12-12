package main

import "testing"

func TestDay12_1(t *testing.T) {
	t.Parallel()
	input := readTestDayN(12)
	expected := "31"
	actual := day12_1(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestDay12_2(t *testing.T) {
	t.Parallel()
	input := readTestDayN(12)
	expected := "29"
	actual := day12_2(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
