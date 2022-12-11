package main

import "testing"

func TestDay11_1(t *testing.T) {
	t.Parallel()
	input := readTestDayN(11)
	expected := "10605"
	actual := day11_1(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestDay11_2(t *testing.T) {
	t.Parallel()
	input := readTestDayN(11)
	expected := "2713310158"
	actual := day11_2(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
