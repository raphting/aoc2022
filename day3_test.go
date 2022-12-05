package main

import "testing"

func TestDay3_1(t *testing.T) {
	t.Parallel()
	input := readTestDayN(3)
	expected := "157"
	actual := day3_1(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestDay3_2(t *testing.T) {
	t.Parallel()
	input := readTestDayN(3)
	expected := "70"
	actual := day3_2(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
