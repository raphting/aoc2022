package main

import (
	"strings"
)

func day3_1(input string) int {
	rows := strings.Split(input, "\n")
	score := 0
	for _, l := range rows {
		p1 := l[:(len(l) / 2)]
		p2 := l[len(l)/2:]

		// cache
		m := make(map[string]bool)
		for _, p := range p1 {
			m[string(p)] = true
		}

		// check
		for _, p := range p2 {
			if _, ok := m[string(p)]; ok {
				score += day3CharToScore(string(p))
				break
			}
		}
	}

	return score
}

func day3_2(input string) int {
	rows := strings.Split(input, "\n")
	rows = rows[:len(rows)-1] // remove last blank line
	score := 0
	for i := 0; i < len(rows); i += 3 {
		// cache
		m := make(map[string][]bool)
		for group := 0; group < 3; group++ {
			for _, p := range rows[i+group] {
				if _, ok := m[string(p)]; ok {
					m[string(p)][group] = true
					continue
				}
				m[string(p)] = make([]bool, 3)
				m[string(p)][group] = true
			}
		}

		// check
		for k, v := range m {
			if v[0] && v[1] && v[2] {
				score += day3CharToScore(k)
			}
		}
	}

	return score
}

func day3CharToScore(c string) int {
	if c == "" {
		return 0
	}
	n := []byte(c)[0]
	if n >= 97 && n <= 122 { // lower case char
		return int(n) - 96
	}

	if n >= 65 && n <= 90 { // upper case char
		return int(n) - 38
	}
	return 0
}
