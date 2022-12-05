package main

import "testing"

func TestDay2_1(t *testing.T) {
	t.Parallel()
	input := readTestDayN(2)
	expected := "15"
	actual := day2_1(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestDay2_2(t *testing.T) {
	t.Parallel()
	input := readTestDayN(2)
	expected := "12"
	actual := day2_2(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
