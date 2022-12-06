package main

import (
	"sort"
	"strconv"
	"strings"
)

func day6_1(input string) string {

	cnt := 4
	for len(input) >= 4 {
		a := input[0]
		b := input[1]
		c := input[2]
		d := input[3]

		if a != b && a != c && a != d && b != c && b != d && c != d {
			return strconv.Itoa(cnt)
		}

		cnt++
		input = input[1:]
	}

	return "0"
}

func day6_2(input string) string {
	cnt := 14
	for len(input) >= 14 {
		check := strings.Split(input[:14], "")
		sort.Strings(check)

		equal := false
		for i := 0; i < len(check)-1; i++ {
			if check[i] == check[i+1] {
				equal = true
				break
			}
		}

		if !equal {
			return strconv.Itoa(cnt)
		}

		cnt++
		input = input[1:]
	}

	return "0"
}
