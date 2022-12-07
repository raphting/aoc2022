package main

import (
	"sort"
	"strconv"
	"strings"
)

type dir struct {
	name     string
	size     int64
	parent   *dir
	children []*dir
	files    map[string]int64
}

func day7_1(input string) string {
	root := &dir{
		name:     "/",
		size:     0,
		parent:   nil,
		children: make([]*dir, 0),
		files:    make(map[string]int64),
	}

	rows := strings.Split(input, "\n")

	currentDir := root
	for _, row := range rows {
		fields := strings.Split(row, " ")
		if strings.HasPrefix(row, "$ cd ") {
			if fields[2] == "/" {
				currentDir = root
				continue
			}
			if fields[2] == ".." {
				currentDir = currentDir.parent
				continue
			}

			// Enter new dir
			newDir := &dir{
				name:     fields[2],
				size:     0,
				parent:   currentDir,
				children: make([]*dir, 0),
				files:    make(map[string]int64),
			}

			currentDir.children = append(currentDir.children, newDir)
			currentDir = newDir
			continue
		}

		if strings.HasPrefix(row, "$ ls") {
			continue
		}

		if strings.HasPrefix(row, "dir ") {
			continue
		}

		// after here, it can only be a file with a given size
		size, err := strconv.ParseInt(fields[0], 10, 64)
		if err != nil {
			panic(err)
		}

		currentDir.files[fields[1]] = size
	}

	traverseDir(root)
	sizes := getDirSizes(root)

	var result int64
	for _, size := range sizes {
		if size < 100000 {
			result += size
		}
	}

	return strconv.Itoa(int(result))
}

func day7_2(input string) string {
	root := &dir{
		name:     "/",
		size:     0,
		parent:   nil,
		children: make([]*dir, 0),
		files:    make(map[string]int64),
	}

	rows := strings.Split(input, "\n")

	currentDir := root
	for _, row := range rows {
		fields := strings.Split(row, " ")
		if strings.HasPrefix(row, "$ cd ") {
			if fields[2] == "/" {
				currentDir = root
				continue
			}
			if fields[2] == ".." {
				currentDir = currentDir.parent
				continue
			}

			// Enter new dir
			newDir := &dir{
				name:     fields[2],
				size:     0,
				parent:   currentDir,
				children: make([]*dir, 0),
				files:    make(map[string]int64),
			}

			currentDir.children = append(currentDir.children, newDir)
			currentDir = newDir
			continue
		}

		if strings.HasPrefix(row, "$ ls") {
			continue
		}

		if strings.HasPrefix(row, "dir ") {
			continue
		}

		// after here, it can only be a file with a given size
		size, err := strconv.ParseInt(fields[0], 10, 64)
		if err != nil {
			panic(err)
		}

		currentDir.files[fields[1]] = size
	}

	traverseDir(root)
	sizes := getDirSizes(root)

	fileSystemSize := int64(70000000)
	usedSpace := root.size
	unusedSpace := fileSystemSize - usedSpace
	requiredSpace := 30000000 - unusedSpace

	candidates := make([]int, 0)
	for _, size := range sizes {
		if size >= requiredSpace {
			candidates = append(candidates, int(size))
		}
	}

	sort.Ints(candidates)
	return strconv.Itoa(candidates[0])
}

func getDirSizes(root *dir) []int64 {
	out := make([]int64, 0)
	for _, child := range root.children {
		if len(child.children) != 0 {
			out = append(out, getDirSizes(child)...)
		}
		out = append(out, child.size)
	}
	return out
}

// traverseDir calculates all sizes for all dirs
func traverseDir(root *dir) {
	for _, child := range root.children {
		if len(child.children) != 0 {
			traverseDir(child)
		}
		var totalSize int64
		for _, size := range child.files {
			totalSize += size
		}
		child.size = totalSize
		for _, c := range child.children {
			child.size += c.size
		}
	}

	for _, size := range root.files {
		root.size += size
	}
	for _, c := range root.children {
		root.size += c.size
	}
}
