package main

import (
	"strconv"
	"strings"
)

func day4_1(input string) string {
	rows := strings.Split(input, "\n")
	type interval struct {
		from1 int64
		to1   int64
		from2 int64
		to2   int64
	}

	intervals := make([]interval, len(rows))

	for cnt, l := range rows {
		pair := strings.Split(l, ",")
		if len(pair) != 2 {
			continue
		}

		interval1 := strings.Split(pair[0], "-")
		interval2 := strings.Split(pair[1], "-")
		from1, _ := strconv.ParseInt(interval1[0], 10, 64)
		to1, _ := strconv.ParseInt(interval1[1], 10, 64)
		from2, _ := strconv.ParseInt(interval2[0], 10, 64)
		to2, _ := strconv.ParseInt(interval2[1], 10, 64)
		intervals[cnt] = interval{from1: from1, to1: to1, from2: from2, to2: to2}
	}

	result := 0

	for _, i := range intervals {
		if i.from1 <= i.from2 && i.to1 >= i.to2 {
			result += 1
			continue
		}
		if i.from2 <= i.from1 && i.to2 >= i.to1 {
			result += 1
		}
	}

	return strconv.Itoa(result)
}

func day4_2(input string) string {
	rows := strings.Split(input, "\n")
	type interval struct {
		from1 int64
		to1   int64
		from2 int64
		to2   int64
	}

	intervals := make([]interval, len(rows))

	for cnt, l := range rows {
		pair := strings.Split(l, ",")
		if len(pair) != 2 {
			continue
		}

		interval1 := strings.Split(pair[0], "-")
		interval2 := strings.Split(pair[1], "-")
		from1, _ := strconv.ParseInt(interval1[0], 10, 64)
		to1, _ := strconv.ParseInt(interval1[1], 10, 64)
		from2, _ := strconv.ParseInt(interval2[0], 10, 64)
		to2, _ := strconv.ParseInt(interval2[1], 10, 64)
		intervals[cnt] = interval{from1: from1, to1: to1, from2: from2, to2: to2}
	}

	result := 0

	for _, i := range intervals {
		if i.from1 >= i.from2 && i.from1 <= i.to2 {
			result += 1
			continue
		}
		if i.to1 >= i.from2 && i.to1 <= i.to2 {
			result += 1
			continue
		}
		if i.from2 >= i.from1 && i.from2 <= i.to1 {
			result += 1
			continue
		}
		if i.to2 >= i.from1 && i.to2 <= i.to1 {
			result += 1
			continue
		}
	}

	return strconv.Itoa(result)
}
