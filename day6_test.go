package main

import "testing"

func TestDay6_1(t *testing.T) {
	t.Parallel()
	input := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
	expected := "7"
	actual := day6_1(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}

	input = "bvwbjplbgvbhsrlpgdmjqwftvncz"
	expected = "5"
	actual = day6_1(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}

	input = "nppdvjthqldpwncqszvftbrmjlhg"
	expected = "6"
	actual = day6_1(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}

	input = "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"
	expected = "10"
	actual = day6_1(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}

	input = "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"
	expected = "11"
	actual = day6_1(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestDay6_2(t *testing.T) {
	t.Parallel()
	input := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
	expected := "19"
	actual := day6_2(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}

	input = "bvwbjplbgvbhsrlpgdmjqwftvncz"
	expected = "23"
	actual = day6_2(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}

	input = "nppdvjthqldpwncqszvftbrmjlhg"
	expected = "23"
	actual = day6_2(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}

	input = "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"
	expected = "29"
	actual = day6_2(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}

	input = "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"
	expected = "26"
	actual = day6_2(input)
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
