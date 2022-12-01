package main

import (
	"sort"
	"strconv"
	"strings"
)

func day1_1(input string) int {
	calories := make([]int, 0)
	for _, elf := range strings.Split(input, "\n\n") {
		calory := 0
		for _, e := range strings.Split(elf, "\n") {
			if e == "" {
				continue
			}
			n, err := strconv.ParseInt(e, 10, 64)
			if err != nil {
				panic(err)
			}
			calory += int(n)
		}
		calories = append(calories, calory)
	}
	sort.Ints(calories)
	return calories[len(calories)-1]
}

func day1_2(input string) int {
	calories := make([]int, 0)
	for _, elf := range strings.Split(input, "\n\n") {
		calory := 0
		for _, e := range strings.Split(elf, "\n") {
			if e == "" {
				continue
			}
			n, err := strconv.ParseInt(e, 10, 64)
			if err != nil {
				panic(err)
			}
			calory += int(n)
		}
		calories = append(calories, calory)
	}
	sort.Ints(calories)
	return calories[len(calories)-1] + calories[len(calories)-2] + calories[len(calories)-3]
}
