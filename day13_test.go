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

func TestDay13_2(t *testing.T) {
	t.Parallel()
	input := readTestDayN(13)
	expected := "140"
	actual := day13_2(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
