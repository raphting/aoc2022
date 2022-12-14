package main

import (
	"strconv"
	"strings"
)

func day14_1(input string) string {
	walls := make(map[int64]map[int64]bool)
	var lowestY int64

	rows := strings.Split(input, "\n")
	for _, r := range rows {
		directions := strings.Split(r, " -> ")

		for len(directions) > 1 {
			xy1 := strings.Split(directions[0], ",")
			xy2 := strings.Split(directions[1], ",")

			x1, _ := strconv.ParseInt(xy1[0], 10, 64)
			y1, _ := strconv.ParseInt(xy1[1], 10, 64)
			x2, _ := strconv.ParseInt(xy2[0], 10, 64)
			y2, _ := strconv.ParseInt(xy2[1], 10, 64)

			// Find out lowest wall to know when sand falls off the edge
			if y1 > lowestY {
				lowestY = y1
			}
			if y2 > lowestY {
				lowestY = y2
			}

			// move vertically
			if x1 == x2 {
				if y1 < y2 {
					for cnt := y1; cnt <= y2; cnt++ {
						if _, ok := walls[cnt]; !ok {
							walls[cnt] = make(map[int64]bool)
						}
						walls[cnt][x1] = true
					}
				}
				if y1 > y2 {
					for cnt := y1; cnt >= y2; cnt-- {
						if _, ok := walls[cnt]; !ok {
							walls[cnt] = make(map[int64]bool)
						}
						walls[cnt][x1] = true
					}
				}
				if y1 == y2 {
					if _, ok := walls[y1]; !ok {
						walls[y1] = make(map[int64]bool)
					}
					walls[y1][x1] = true
				}
			}

			// move horizontally
			if y1 == y2 {
				if x1 < x2 {
					for cnt := x1; cnt <= x2; cnt++ {
						if _, ok := walls[y1]; !ok {
							walls[y1] = make(map[int64]bool)
						}
						walls[y1][cnt] = true
					}
				}
				if x1 > x2 {
					for cnt := x1; cnt >= x2; cnt-- {
						if _, ok := walls[y1]; !ok {
							walls[y1] = make(map[int64]bool)
						}
						walls[y1][cnt] = true
					}
				}
				if x1 == x2 {
					if _, ok := walls[y1]; !ok {
						walls[y1] = make(map[int64]bool)
					}
					walls[y1][x1] = true
				}
			}

			directions = directions[1:]
		}
	}

	restingSand := make(map[int64]map[int64]bool)

	startX := int64(500)
	startY := int64(0)
	for {
		newX, newY := getNextSandPosition(startX, startY, walls, restingSand)

		if newY > lowestY {
			break
		}

		if newX == -1 && newY == -1 {
			if _, ok := restingSand[startY]; !ok {
				restingSand[startY] = make(map[int64]bool)
			}
			restingSand[startY][startX] = true
			startX = 500
			startY = 0
			continue
		}
		startX = newX
		startY = newY
	}

	// count sand elements
	sum := 0
	for _, y := range restingSand {
		sum += len(y)
	}

	return strconv.Itoa(sum)
}

func getNextSandPosition(currentX, currentY int64, walls, restingSand map[int64]map[int64]bool) (int64, int64) {
	// try direct down
	newX := currentX
	newY := currentY + 1
	if conflictsWith(newX, newY, walls) || conflictsWith(newX, newY, restingSand) {
		// try left down
		newX = currentX - 1
		if conflictsWith(newX, newY, walls) || conflictsWith(newX, newY, restingSand) {
			// try right down
			newX = currentX + 1
			if conflictsWith(newX, newY, walls) || conflictsWith(newX, newY, restingSand) {
				// this comes to rest
				return -1, -1
			} else {
				return newX, newY
			}
		} else {
			return newX, newY
		}
	} else {
		return newX, newY
	}
}

func conflictsWith(x, y int64, obstacles map[int64]map[int64]bool) bool {
	if _, ok := obstacles[y]; ok {
		if _, ok2 := obstacles[y][x]; ok2 {
			return true
		}
	}
	return false
}
