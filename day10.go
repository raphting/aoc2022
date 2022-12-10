package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day10_1(input string) string {
	rows := strings.Split(input, "\n")

	type cpuInstruction struct {
		command string
		value   int
	}

	cpuInstructions := make([]cpuInstruction, len(rows))
	for cnt, r := range rows {
		fields := strings.Split(r, " ")
		if fields[0] == "noop" {
			cpuInstructions[cnt] = cpuInstruction{
				command: "noop",
				value:   0,
			}
			continue
		}

		if fields[0] == "addx" {
			value, _ := strconv.ParseInt(fields[1], 10, 64)
			cpuInstructions[cnt] = cpuInstruction{
				command: "addx",
				value:   int(value),
			}
			continue
		}
	}

	ticks := 0
	registerX := 1
	probe := make([]int, 0)
	for _, i := range cpuInstructions {
		switch i.command {
		case "noop":
			ticks++
			if ticks == 20 || ticks == 60 || ticks == 100 || ticks == 140 || ticks == 180 || ticks == 220 {
				probe = append(probe, registerX)
			}
		case "addx":
			ticks++
			if ticks == 20 || ticks == 60 || ticks == 100 || ticks == 140 || ticks == 180 || ticks == 220 {
				probe = append(probe, registerX)
			}

			ticks++
			if ticks == 20 || ticks == 60 || ticks == 100 || ticks == 140 || ticks == 180 || ticks == 220 {
				probe = append(probe, registerX)
			}
			registerX += i.value
		}
	}

	signalStrength := 20*probe[0] + 60*probe[1] + 100*probe[2] + 140*probe[3] + 180*probe[4] + 220*probe[5]
	return strconv.Itoa(signalStrength)
}

func day10_2(input string) string {
	rows := strings.Split(input, "\n")

	type cpuInstruction struct {
		command string
		value   int
	}

	cpuInstructions := make([]cpuInstruction, len(rows))
	for cnt, r := range rows {
		fields := strings.Split(r, " ")
		if fields[0] == "noop" {
			cpuInstructions[cnt] = cpuInstruction{
				command: "noop",
				value:   0,
			}
			continue
		}

		if fields[0] == "addx" {
			value, _ := strconv.ParseInt(fields[1], 10, 64)
			cpuInstructions[cnt] = cpuInstruction{
				command: "addx",
				value:   int(value),
			}
			continue
		}
	}

	pixelDrawer := func(regX, tick int) string {
		if regX-1 == tick%40 || regX == tick%40 || regX+1 == tick%40 {
			return "#"
		}
		return "."
	}

	ticks := 0
	registerX := 1
	display := make([]string, 0)
	for _, i := range cpuInstructions {
		switch i.command {
		case "noop":
			display = append(display, pixelDrawer(registerX, ticks))
			ticks++
		case "addx":
			display = append(display, pixelDrawer(registerX, ticks))
			ticks++

			display = append(display, pixelDrawer(registerX, ticks))
			ticks++
			registerX += i.value
		}
	}

	solution := "\n"
	solution += strings.Join(display[:40], "")
	solution = fmt.Sprintln(solution)
	solution += strings.Join(display[40:80], "")
	solution = fmt.Sprintln(solution)
	solution += strings.Join(display[80:120], "")
	solution = fmt.Sprintln(solution)
	solution += strings.Join(display[120:160], "")
	solution = fmt.Sprintln(solution)
	solution += strings.Join(display[160:200], "")
	solution = fmt.Sprintln(solution)
	solution += strings.Join(display[200:], "")
	solution = fmt.Sprintln(solution)
	return solution
}
