package main

import "testing"

func TestDay4_1(t *testing.T) {
	t.Parallel()
	input := readTestDayN(4)
	expected := "2"
	actual := day4_1(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestDay4_2(t *testing.T) {
	t.Parallel()
	input := readTestDayN(4)
	expected := "4"
	actual := day4_2(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
