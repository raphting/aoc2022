package main

import "fmt"

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
	}

	for _, d := range days {
		input := readDayN(d.day)
		fmt.Printf("Day %v\n", d.day)
		if d.task1 != nil {
			fmt.Printf("Task1: %v\n", d.task1(input))
		}
		if d.task2 != nil {
			fmt.Printf("Task2: %v\n", d.task2(input))
		}
		fmt.Println("-----------------")
	}
}
