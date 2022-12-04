package main

import (
	"bytes"
	"fmt"
	"os"
)

func readDayN(n int) string {
	return readFileAndTrim(fmt.Sprintf("input/input_day%v.txt", n))
}

func readTestDayN(n int) string {
	return readFileAndTrim(fmt.Sprintf("testinput/input_day%v.txt", n))
}

func readFileAndTrim(path string) string {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	// detect empty last blank line and remove it
	if bytes.HasSuffix(input, []byte("\n")) {
		input = input[:len(input)-1]
	}

	return string(input)
}
