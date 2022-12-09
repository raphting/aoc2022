package main

import (
	"math"
	"strconv"
	"strings"
)

func day9_2(input string) string {
	type instruction struct {
		steps     int
		direction string
	}

	rows := strings.Split(input, "\n")
	instructions := make([]instruction, len(rows))

	for cnt, r := range rows {
		fields := strings.Split(r, " ")
		steps, _ := strconv.ParseInt(fields[1], 10, 64)
		instructions[cnt] = instruction{
			steps:     int(steps),
			direction: fields[0],
		}
	}

	type position struct {
		x, y int
	}

	ropePositions := make([]position, 10)
	tailPositions := make([]position, 0)
	// consider start position
	tailPositions = append(tailPositions, position{
		x: 0,
		y: 0,
	})

	head := position{
		x: 0,
		y: 0,
	}
	tail := position{
		x: 0,
		y: 0,
	}

	for _, inst := range instructions {
		for i := 0; i < inst.steps; i++ {
			head = ropePositions[0]
			switch inst.direction {
			case "R":
				head.x++
			case "U":
				head.y--
			case "L":
				head.x--
			case "D":
				head.y++
			}

			for cnt, r := range ropePositions {
				if cnt == len(ropePositions)-1 {
					break
				}
				if cnt == 0 {
					ropePositions[0] = head
					tail = ropePositions[1]
				} else {
					head = r
					tail = ropePositions[cnt+1]
				}

				xDistance := int(math.Abs(float64(tail.x - head.x)))
				yDistance := int(math.Abs(float64(tail.y - head.y)))

				if xDistance <= 1 && yDistance <= 1 {
					// tail wont move, distance too small
					continue
				}

				if head.y == tail.y {
					if head.x == tail.x+2 {
						tail.x++
					}

					if head.x == tail.x-2 {
						tail.x--
					}
					ropePositions[cnt+1] = tail

					if cnt == len(ropePositions)-2 {
						tailPositions = append(tailPositions, tail)
					}

					continue
				}

				if head.x == tail.x {
					if head.y == tail.y+2 {
						tail.y++
					}
					if head.y == tail.y-2 {
						tail.y--
					}
					ropePositions[cnt+1] = tail
					if cnt == len(ropePositions)-2 {
						tailPositions = append(tailPositions, tail)
					}
					continue
				}

				// diagonal move
				// 4 possible moves
				if tail.x < head.x && tail.y < head.y {
					tail.x++
					tail.y++
					ropePositions[cnt+1] = tail
					if cnt == len(ropePositions)-2 {
						tailPositions = append(tailPositions, tail)
					}
					continue
				}

				if tail.x < head.x && tail.y > head.y {
					tail.x++
					tail.y--
					ropePositions[cnt+1] = tail
					if cnt == len(ropePositions)-2 {
						tailPositions = append(tailPositions, tail)
					}
					continue
				}

				if tail.x > head.x && tail.y > head.y {
					tail.x--
					tail.y--
					ropePositions[cnt+1] = tail
					if cnt == len(ropePositions)-2 {
						tailPositions = append(tailPositions, tail)
					}
					continue
				}

				if tail.x > head.x && tail.y < head.y {
					tail.x--
					tail.y++
					ropePositions[cnt+1] = tail
					if cnt == len(ropePositions)-2 {
						tailPositions = append(tailPositions, tail)
					}
					continue
				}
			}
		}
	}

	uniqueTailPositions := make([]position, 0, len(tailPositions))
	for _, tp := range tailPositions {
		found := false
		for _, utp := range uniqueTailPositions {
			if tp.x == utp.x && tp.y == utp.y {
				found = true
				break
			}
		}
		if !found {
			uniqueTailPositions = append(uniqueTailPositions, tp)
		}
	}

	return strconv.Itoa(len(uniqueTailPositions))
}

func day9_1(input string) string {
	type instruction struct {
		steps     int
		direction string
	}

	rows := strings.Split(input, "\n")
	instructions := make([]instruction, len(rows))

	for cnt, r := range rows {
		fields := strings.Split(r, " ")
		steps, _ := strconv.ParseInt(fields[1], 10, 64)
		instructions[cnt] = instruction{
			steps:     int(steps),
			direction: fields[0],
		}
	}

	type position struct {
		x, y int
	}

	head := position{
		x: 0,
		y: 0,
	}
	tail := position{
		x: 0,
		y: 0,
	}

	tailPositions := make([]position, 0)
	// consider start position
	tailPositions = append(tailPositions, position{
		x: 0,
		y: 0,
	})
	for _, inst := range instructions {
		for i := 0; i < inst.steps; i++ {
			switch inst.direction {
			case "R":
				head.x++
			case "U":
				head.y--
			case "L":
				head.x--
			case "D":
				head.y++
			}

			xDistance := int(math.Abs(float64(tail.x - head.x)))
			yDistance := int(math.Abs(float64(tail.y - head.y)))

			if xDistance <= 1 && yDistance <= 1 {
				// tail wont move, distance too small
				continue
			}

			if head.y == tail.y {
				if head.x == tail.x+2 {
					tail.x++
				}

				if head.x == tail.x-2 {
					tail.x--
				}
				tailPositions = append(tailPositions, tail)
				continue
			}

			if head.x == tail.x {
				if head.y == tail.y+2 {
					tail.y++
				}
				if head.y == tail.y-2 {
					tail.y--
				}
				tailPositions = append(tailPositions, tail)
				continue
			}

			// diagonal move
			// 4 possible moves
			if tail.x < head.x && tail.y < head.y {
				tail.x++
				tail.y++
				tailPositions = append(tailPositions, tail)
				continue
			}

			if tail.x < head.x && tail.y > head.y {
				tail.x++
				tail.y--
				tailPositions = append(tailPositions, tail)
				continue
			}

			if tail.x > head.x && tail.y > head.y {
				tail.x--
				tail.y--
				tailPositions = append(tailPositions, tail)
				continue
			}

			if tail.x > head.x && tail.y < head.y {
				tail.x--
				tail.y++
				tailPositions = append(tailPositions, tail)
				continue
			}
		}
	}

	uniqueTailPositions := make([]position, 0, len(tailPositions))
	for _, tp := range tailPositions {
		found := false
		for _, utp := range uniqueTailPositions {
			if tp.x == utp.x && tp.y == utp.y {
				found = true
				break
			}
		}
		if !found {
			uniqueTailPositions = append(uniqueTailPositions, tp)
		}
	}
	return strconv.Itoa(len(uniqueTailPositions))
}
