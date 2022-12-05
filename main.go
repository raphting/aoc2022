package main

import (
	"fmt"
)

func main() {
	type aoc struct {
		day          int
		task1, task2 func(string) int
	}

	days := []aoc{
		{
			day:   1,
			task1: day1_1,
			task2: day1_2,
		},
		{
			day:   2,
			task1: day2_1,
			task2: day2_2,
		},
		{
			day:   3,
			task1: day3_1,
			task2: day3_2,
		},
		{
			day:   4,
			task1: day4_1,
			task2: day4_2,
		},
		{
			day:   5,
			task1: day5_1,
			task2: day5_2,
		},
	}

	type result struct {
		day, part, solution, seq int
	}
	results := make(chan result)
	numRoutines := 0
	for _, d := range days {
		go func(d aoc, n int) {
			input := readDayN(d.day)
			results <- result{day: d.day, part: 1, solution: d.task1(input), seq: n}
		}(d, numRoutines)
		numRoutines++

		if d.task2 != nil {
			go func(d aoc, n int) {
				input := readDayN(d.day)
				results <- result{day: d.day, part: 2, solution: d.task2(input), seq: n}
			}(d, numRoutines)
			numRoutines++
		}
	}

	sortedResults := make([]result, numRoutines)
	for n := 0; n < numRoutines; n++ {
		r := <-results
		sortedResults[r.seq] = r
	}

	for _, s := range sortedResults {
		fmt.Printf("Day %v, Part %v: %v\n", s.day, s.part, s.solution)
	}
}
