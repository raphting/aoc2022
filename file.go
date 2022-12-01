package main

import (
	"fmt"
	"os"
)

func readDayN(n int) string {
	input, err := os.ReadFile(fmt.Sprintf("input/input_day%v.txt", n))
	if err != nil {
		panic(err)
	}

	return string(input)
}

func readTestDayN(n int) string {
	input, err := os.ReadFile(fmt.Sprintf("testinput/input_day%v.txt", n))
	if err != nil {
		panic(err)
	}

	return string(input)
}
