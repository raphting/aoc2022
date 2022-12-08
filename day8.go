package main

import (
	"strconv"
	"strings"
)

type baum struct {
	height int64
	seen   bool
}

func day8_1(input string) string {
	rows := strings.Split(input, "\n")
	forest := make([][]baum, len(rows))

	for i, line := range rows {
		trees := strings.Split(line, "")
		treeHeights := make([]baum, len(trees))
		for x, t := range trees {
			treeHeight, err := strconv.ParseInt(t, 10, 64)
			if err != nil {
				panic(err)
			}
			treeHeights[x] = baum{
				height: treeHeight,
				seen:   false,
			}
		}
		forest[i] = treeHeights
	}

	numRows := len(forest)
	result := 0
	for cnt, row := range forest {
		if cnt == 0 || cnt == len(forest)-1 {
			continue
		}
		row, num := lookIntoForest(row)
		result += num
		// reverse
		for i, j := 0, len(row)-1; i < j; i, j = i+1, j-1 {
			row[i], row[j] = row[j], row[i]
		}
		row, num = lookIntoForest(row)
		result += num
		// reverse
		for i, j := 0, len(row)-1; i < j; i, j = i+1, j-1 {
			row[i], row[j] = row[j], row[i]
		}
		forest[cnt] = row
	}

	// convert rows to columns
	columnForest := make([][]baum, 0)
	for _, row := range forest {
		for cnt, e := range row {
			if len(columnForest) <= cnt {
				columnForest = append(columnForest, []baum{})
			}
			columnForest[cnt] = append(columnForest[cnt], e)
		}
	}

	numColumns := len(columnForest)
	for cnt, row := range columnForest {
		if cnt == 0 || cnt == len(columnForest)-1 {
			continue
		}
		row, num := lookIntoForest(row)
		result += num
		// reverse
		for i, j := 0, len(row)-1; i < j; i, j = i+1, j-1 {
			row[i], row[j] = row[j], row[i]
		}
		_, num = lookIntoForest(row)
		result += num
		// reverse
		for i, j := 0, len(row)-1; i < j; i, j = i+1, j-1 {
			row[i], row[j] = row[j], row[i]
		}
		columnForest[cnt] = row
	}

	return strconv.Itoa(result + ((numRows * 2) + ((numColumns - 2) * 2)))
}

func day8_2(input string) string {
	rows := strings.Split(input, "\n")
	forest := make([][]baum, len(rows))
	for i, line := range rows {
		trees := strings.Split(line, "")
		treeHeights := make([]baum, len(trees))
		for x, t := range trees {
			treeHeight, err := strconv.ParseInt(t, 10, 64)
			if err != nil {
				panic(err)
			}
			treeHeights[x] = baum{
				height: treeHeight,
				seen:   false,
			}
		}
		forest[i] = treeHeights
	}

	type scenicTree struct {
		x, y                  int
		up, down, left, right int
	}

	// make list of all trees to consider
	numRows := len(rows)
	numColumns := len(rows[0])

	allScenicTrees := make([]scenicTree, 0)
	for y := 1; y < numRows-1; y++ {
		for x := 1; x < numColumns-1; x++ {
			allScenicTrees = append(allScenicTrees, scenicTree{
				x:     x,
				y:     y,
				up:    0,
				down:  0,
				left:  0,
				right: 0,
			})
		}
	}

	// analyze all scenic trees
	for t := range allScenicTrees {
		posX := allScenicTrees[t].x
		posY := allScenicTrees[t].y
		treeHeight := forest[posY][posX].height
		// go up
		up := 0
		for y := posY - 1; y >= 0; y-- {
			if forest[y][posX].height >= treeHeight {
				up++
				break
			}
			up++
		}
		allScenicTrees[t].up = up

		// go down
		down := 0
		for y := posY + 1; y < numColumns; y++ {
			if forest[y][posX].height >= treeHeight {
				down++
				break
			}
			down++
		}
		allScenicTrees[t].down = down

		// go left
		left := 0
		for x := posX - 1; x >= 0; x-- {
			if forest[posY][x].height >= treeHeight {
				left++
				break
			}
			left++
		}
		allScenicTrees[t].left = left

		// go right
		right := 0
		for x := posX + 1; x < numRows; x++ {
			if forest[posY][x].height >= treeHeight {
				right++
				break
			}
			right++
		}
		allScenicTrees[t].right = right
	}

	// analyze best scenic view
	bestView := 0
	for _, t := range allScenicTrees {
		scenicScore := t.up * t.down * t.left * t.right
		if scenicScore > bestView {
			bestView = scenicScore
		}
	}

	return strconv.Itoa(bestView)
}

func lookIntoForest(row []baum) ([]baum, int) {
	out := 0
	max := row[0].height
	for i := 1; i < len(row)-1; i++ {
		if row[i].height > max && !row[i].seen {
			max = row[i].height
			row[i].seen = true
			out++
		}
		if row[i].height > max {
			max = row[i].height
		}
	}
	return row, out
}
