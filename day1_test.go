package main

import "testing"

func TestDay1_1(t *testing.T) {
	t.Parallel()
	input := readTestDayN(1)
	expected := 24000
	actual := day1_1(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestDay1_2(t *testing.T) {
	t.Parallel()
	input := readTestDayN(1)
	expected := 45000
	actual := day1_2(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
