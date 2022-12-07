package main

import "testing"

func TestDay7_1(t *testing.T) {
	t.Parallel()
	input := readTestDayN(7)
	expected := "95437"
	actual := day7_1(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestDay7_2(t *testing.T) {
	t.Parallel()
	input := readTestDayN(7)
	expected := "24933642"
	actual := day7_2(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
