package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day5_1(input string) int {
	rows := strings.Split(input, "\n")

	// Find block positions
	blockPositions := make([]int, 0)
	for _, r := range rows {
		if strings.HasPrefix(r, " 1") {
			for p, n := range r {
				if string(n) != " " {
					blockPositions = append(blockPositions, p)
				}
			}
			break
		}
	}

	type move struct {
		amount, from, to int
	}

	section1 := true
	cargo := make([][]string, 0)
	moves := make([]move, 0)
	for _, r := range rows {
		if r == "" {
			continue
		}
		if strings.HasPrefix(r, " 1") {
			section1 = false
			continue
		}

		// input is split into 2 sections, divided by a row of numbers
		if section1 {
			c := make([]string, 0)
			for _, p := range blockPositions {
				if len(r) >= p {
					c = append(c, string(r[p]))
				}
			}
			cargo = append(cargo, c)
		} else {
			if strings.HasPrefix(r, "move ") {
				fields := strings.Split(r, " ")
				amount, _ := strconv.ParseInt(fields[1], 10, 64)
				from, _ := strconv.ParseInt(fields[3], 10, 64)
				to, _ := strconv.ParseInt(fields[5], 10, 64)
				moves = append(moves, move{
					amount: int(amount),
					from:   int(from),
					to:     int(to),
				})
			}
		}
	}

	// reorganize cargo to a column based layout
	cargoColumn := make([][]string, 0)
	// bottom element comes first
	for i := len(cargo) - 1; i >= 0; i-- {
		row := cargo[i]
		for cnt, e := range row {
			if e == " " {
				continue
			}
			if len(cargoColumn) <= cnt {
				cargoColumn = append(cargoColumn, []string{})
			}
			cargoColumn[cnt] = append(cargoColumn[cnt], e)
		}
	}

	// work on the moves
	for _, m := range moves {
		take := cargoColumn[m.from-1][len(cargoColumn[m.from-1])-m.amount:]
		cargoColumn[m.from-1] = cargoColumn[m.from-1][:len(cargoColumn[m.from-1])-m.amount]

		// reverse take
		for i, j := 0, len(take)-1; i < j; i, j = i+1, j-1 {
			take[i], take[j] = take[j], take[i]
		}

		cargoColumn[m.to-1] = append(cargoColumn[m.to-1], take...)
	}

	for cnt := range cargoColumn {
		fmt.Print(cargoColumn[cnt][len(cargoColumn[cnt])-1])
	}
	fmt.Println()
	return 0
}

func day5_2(input string) int {
	rows := strings.Split(input, "\n")

	// Find block positions
	blockPositions := make([]int, 0)
	for _, r := range rows {
		if strings.HasPrefix(r, " 1") {
			for p, n := range r {
				if string(n) != " " {
					blockPositions = append(blockPositions, p)
				}
			}
			break
		}
	}

	type move struct {
		amount, from, to int
	}

	section1 := true
	cargo := make([][]string, 0)
	moves := make([]move, 0)
	for _, r := range rows {
		if r == "" {
			continue
		}
		if strings.HasPrefix(r, " 1") {
			section1 = false
			continue
		}

		// input is split into 2 sections, divided by a row of numbers
		if section1 {
			c := make([]string, 0)
			for _, p := range blockPositions {
				if len(r) >= p {
					c = append(c, string(r[p]))
				}
			}
			cargo = append(cargo, c)
		} else {
			if strings.HasPrefix(r, "move ") {
				fields := strings.Split(r, " ")
				amount, _ := strconv.ParseInt(fields[1], 10, 64)
				from, _ := strconv.ParseInt(fields[3], 10, 64)
				to, _ := strconv.ParseInt(fields[5], 10, 64)
				moves = append(moves, move{
					amount: int(amount),
					from:   int(from),
					to:     int(to),
				})
			}
		}
	}

	// reorganize cargo to a column based layout
	cargoColumn := make([][]string, 0)
	// bottom element comes first
	for i := len(cargo) - 1; i >= 0; i-- {
		row := cargo[i]
		for cnt, e := range row {
			if e == " " {
				continue
			}
			if len(cargoColumn) <= cnt {
				cargoColumn = append(cargoColumn, []string{})
			}
			cargoColumn[cnt] = append(cargoColumn[cnt], e)
		}
	}

	// work on the moves
	for _, m := range moves {
		take := cargoColumn[m.from-1][len(cargoColumn[m.from-1])-m.amount:]
		cargoColumn[m.from-1] = cargoColumn[m.from-1][:len(cargoColumn[m.from-1])-m.amount]
		cargoColumn[m.to-1] = append(cargoColumn[m.to-1], take...)
	}

	for cnt := range cargoColumn {
		fmt.Print(cargoColumn[cnt][len(cargoColumn[cnt])-1])
	}
	fmt.Println()
	return 0
}
