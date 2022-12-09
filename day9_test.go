package main

import "testing"

func TestDay9_1(t *testing.T) {
	t.Parallel()
	input := readTestDayN(9)
	expected := "13"
	actual := day9_1(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestDay9_2(t *testing.T) {
	t.Parallel()
	input := readTestDayN(9)
	expected := "1"
	actual := day9_2(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}

	input = readTestDayN(92)
	expected = "36"
	actual = day9_2(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
