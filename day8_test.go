package main

import "testing"

func TestDay8_1(t *testing.T) {
	t.Parallel()
	input := readTestDayN(8)
	expected := "21"
	actual := day8_1(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestDay8_2(t *testing.T) {
	t.Parallel()
	input := readTestDayN(8)
	expected := "8"
	actual := day8_2(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
