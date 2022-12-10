package main

import "testing"

func TestDay10_1(t *testing.T) {
	t.Parallel()
	input := readTestDayN(10)
	expected := "13140"
	actual := day10_1(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestDay10_2(t *testing.T) {
	t.Parallel()
	input := readTestDayN(10)
	expected := "\n##..##..##..##..##..##..##..##..##..##..\n###...###...###...###...###...###...###.\n####....####....####....####....####....\n#####.....#####.....#####.....#####.....\n######......######......######......####\n#######.......#######.......#######.....\n"
	actual := day10_2(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
