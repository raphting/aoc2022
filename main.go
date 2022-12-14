package main

import (
	"fmt"
)

func main() {
	type aoc struct {
		day          int
		task1, task2 func(string) string
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
		{
			day:   6,
			task1: day6_1,
			task2: day6_2,
		},
		{
			day:   7,
			task1: day7_1,
			task2: day7_2,
		},
		{
			day:   8,
			task1: day8_1,
			task2: day8_2,
		},
		{
			day:   9,
			task1: day9_1,
			task2: day9_2,
		},
		{
			day:   10,
			task1: day10_1,
			task2: day10_2,
		},
		{
			day:   11,
			task1: day11_1,
			task2: day11_2,
		},
		{
			day:   12,
			task1: day12_1,
			task2: day12_2,
		},
		{
			day:   13,
			task1: day13_1,
			task2: day13_2,
		},
		{
			day:   14,
			task1: day14_1,
			task2: day14_2,
		},
		{
			day:   15,
			task1: day15_1,
			task2: day15_2,
		},
	}

	type result struct {
		day, part, seq int
		solution       string
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
