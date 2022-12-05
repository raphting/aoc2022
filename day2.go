package main

import (
	"strconv"
	"strings"
)

func day2_1(input string) string {
	rows := strings.Split(input, "\n")
	score := 0
	for _, r := range rows {
		game := strings.Split(r, " ")
		if len(game) != 2 {
			continue
		}
		score += day2CalcPoints(game[0], game[1])
	}
	return strconv.Itoa(score)
}

func day2CalcPoints(opponent, me string) int {
	if opponent == "A" {
		if me == "X" {
			return 3 + 1

		}
		if me == "Y" {
			return 6 + 2

		}
		if me == "Z" {
			return 0 + 3

		}
	}

	if opponent == "B" {
		if me == "X" {
			return 0 + 1

		}
		if me == "Y" {
			return 3 + 2

		}
		if me == "Z" {
			return 6 + 3

		}
	}

	if opponent == "C" {
		if me == "X" {
			return 6 + 1

		}
		if me == "Y" {
			return 0 + 2

		}
		if me == "Z" {
			return 3 + 3

		}
	}
	return 0
}

func day2_2(input string) string {
	rows := strings.Split(input, "\n")
	score := 0
	for _, r := range rows {
		game := strings.Split(r, " ")
		if len(game) != 2 {
			continue
		}

		// need to loose
		if game[1] == "X" {
			if game[0] == "A" { // rock
				score += day2CalcPoints("A", "Z") // scissor
				continue
			}
			if game[0] == "B" { // paper
				score += day2CalcPoints("B", "X") // rock
				continue
			}
			if game[0] == "C" { // scissor
				score += day2CalcPoints("C", "Y") // paper
				continue
			}
		}

		// need to draw
		if game[1] == "Y" {
			if game[0] == "A" { // rock
				score += day2CalcPoints("A", "X") // rock
				continue
			}
			if game[0] == "B" { // paper
				score += day2CalcPoints("B", "Y") // paper
				continue
			}
			if game[0] == "C" { // scissor
				score += day2CalcPoints("C", "Z") // scissor
				continue
			}
		}

		// need to win
		if game[1] == "Z" {
			if game[0] == "A" { // rock
				score += day2CalcPoints("A", "Y") // paper
				continue
			}
			if game[0] == "B" { // paper
				score += day2CalcPoints("B", "Z") // scissor
				continue
			}
			if game[0] == "C" { // scissor
				score += day2CalcPoints("C", "X") // rock
				continue
			}
		}
	}

	return strconv.Itoa(score)
}
