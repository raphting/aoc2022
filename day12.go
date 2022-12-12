package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func day12_1(input string) string {
	rows := strings.Split(input, "\n")
	heightmap := make([][]string, len(rows))
	for cntR, r := range rows {
		vertices := strings.Split(r, "")
		heightmap[cntR] = make([]string, len(vertices))
		for cntV, v := range vertices {
			heightmap[cntR][cntV] = v
		}
	}

	start := &vertex{}
	end := &vertex{}

	// create all pointers
	vertices := make([][]*vertex, len(heightmap))
	for v1 := range vertices {
		vertices[v1] = make([]*vertex, len(heightmap[v1]))
		for v2 := range vertices[v1] {
			vertices[v1][v2] = &vertex{}
		}
	}

	for y := range heightmap {
		for x := range heightmap[y] {
			v := vertices[y][x]
			v.x = x
			v.y = y

			height := heightmap[y][x]

			if height == "S" {
				start = v
			}
			if height == "E" {
				end = v
			}

			// up
			if y > 0 {
				neighborHeight := heightmap[y-1][x]
				if canClimb(height, neighborHeight) {
					pointerToNeighbor := vertices[y-1][x]
					v.neighbors = append(v.neighbors, pointerToNeighbor)
				}
			}

			// down
			if y < len(heightmap)-1 {
				neighborHeight := heightmap[y+1][x]
				if canClimb(height, neighborHeight) {
					pointerToNeighbor := vertices[y+1][x]
					v.neighbors = append(v.neighbors, pointerToNeighbor)
				}
			}

			// left
			if x > 0 {
				neighborHeight := heightmap[y][x-1]
				if canClimb(height, neighborHeight) {
					pointerToNeighbor := vertices[y][x-1]
					v.neighbors = append(v.neighbors, pointerToNeighbor)
				}
			}

			// right
			if x < len(heightmap[y])-1 {
				neighborHeight := heightmap[y][x+1]
				if canClimb(height, neighborHeight) {
					pointerToNeighbor := vertices[y][x+1]
					v.neighbors = append(v.neighbors, pointerToNeighbor)
				}
			}
		}
	}

	// dijkstras algorithm
	q := make([]*vertex, 0)
	for _, y := range vertices {
		for _, v := range y {
			v.dist = math.MaxInt
			v.prev = nil

			if v.x == start.x && v.y == start.y {
				v.dist = 0
			}

			q = append(q, v)
		}
	}

	for len(q) > 0 {
		sort.Slice(q, func(i, j int) bool {
			return q[i].dist < q[j].dist
		})

		u := q[0] // u has min dist
		if u.x == end.x && u.y == end.y {
			break
		}
		q = q[1:] // remove u from q

		for _, v := range u.neighbors {
			// check that v is part of q
			foundInQ := false
			for _, checkQ := range q {
				if checkQ.x == v.x && checkQ.y == v.y {
					foundInQ = true
					break
				}
			}

			if foundInQ {
				alt := u.dist + 1 // edges always have the length 1
				if alt < v.dist {
					v.dist = alt
					v.prev = u
				}
			}
		}
	}

	sequence := make([]*vertex, 0)
	u := end

	for {
		sequence = append(sequence, u)
		u = u.prev
		if u == nil {
			break
		}
	}

	return fmt.Sprint(len(sequence) - 1)
}

func day12_2(input string) string {
	rows := strings.Split(input, "\n")
	heightmap := make([][]string, len(rows))
	for cntR, r := range rows {
		vertices := strings.Split(r, "")
		heightmap[cntR] = make([]string, len(vertices))
		for cntV, v := range vertices {
			heightmap[cntR][cntV] = v
		}
	}

	starts := make([]*vertex, 0)
	end := &vertex{}

	// create all pointers
	vertices := make([][]*vertex, len(heightmap))
	for v1 := range vertices {
		vertices[v1] = make([]*vertex, len(heightmap[v1]))
		for v2 := range vertices[v1] {
			vertices[v1][v2] = &vertex{}
		}
	}

	for y := range heightmap {
		for x := range heightmap[y] {
			v := vertices[y][x]
			v.x = x
			v.y = y

			height := heightmap[y][x]

			if height == "a" || height == "S" {
				starts = append(starts, v)
			}
			if height == "E" {
				end = v
			}

			// up
			if y > 0 {
				neighborHeight := heightmap[y-1][x]
				if canClimb2(height, neighborHeight) {
					pointerToNeighbor := vertices[y-1][x]
					v.neighbors = append(v.neighbors, pointerToNeighbor)
				}
			}

			// down
			if y < len(heightmap)-1 {
				neighborHeight := heightmap[y+1][x]
				if canClimb2(height, neighborHeight) {
					pointerToNeighbor := vertices[y+1][x]
					v.neighbors = append(v.neighbors, pointerToNeighbor)
				}
			}

			// left
			if x > 0 {
				neighborHeight := heightmap[y][x-1]
				if canClimb2(height, neighborHeight) {
					pointerToNeighbor := vertices[y][x-1]
					v.neighbors = append(v.neighbors, pointerToNeighbor)
				}
			}

			// right
			if x < len(heightmap[y])-1 {
				neighborHeight := heightmap[y][x+1]
				if canClimb2(height, neighborHeight) {
					pointerToNeighbor := vertices[y][x+1]
					v.neighbors = append(v.neighbors, pointerToNeighbor)
				}
			}
		}
	}

	// dijkstras algorithm
	q := make([]*vertex, 0)
	for _, y := range vertices {
		for _, v := range y {
			v.dist = math.MaxInt
			v.prev = nil

			if v.x == end.x && v.y == end.y {
				v.dist = 0
			}

			q = append(q, v)
		}
	}

	for len(q) > 0 {
		sort.Slice(q, func(i, j int) bool {
			return q[i].dist < q[j].dist
		})
		u := q[0] // u has min dist
		q = q[1:] // remove u from q

		for _, v := range u.neighbors {
			// check that v is part of q
			foundInQ := false
			for _, checkQ := range q {
				if checkQ.x == v.x && checkQ.y == v.y {
					foundInQ = true
					break
				}
			}

			if foundInQ {
				alt := u.dist + 1 // edges always have the length 1
				if alt < v.dist {
					v.dist = alt
					v.prev = u
				}
			}
		}
	}

	distances := make([]int, 0)
	for _, s := range starts {
		sequence := make([]*vertex, 0)
		u := s

		for {
			sequence = append(sequence, u)
			u = u.prev
			if u == nil {
				break
			}
		}

		//  horrible hack, but it gets the job done :)
		if len(sequence) > len(vertices) {
			distances = append(distances, len(sequence)-1)
		}
	}

	sort.Ints(distances)
	return fmt.Sprint(distances[0])
}

func canClimb(a, b string) bool {
	if a == "S" {
		a = "a"
	}

	if b == "S" {
		b = "a"
	}

	if a == "E" {
		a = "z"
	}

	if b == "E" {
		b = "z"
	}

	if int(a[0]) == int(b[0]) || int(a[0]) == int(b[0])-1 || int(a[0]) > int(b[0]) {
		return true
	}
	return false
}

func canClimb2(a, b string) bool {
	if a == "S" {
		a = "a"
	}

	if b == "S" {
		b = "a"
	}

	if a == "E" {
		a = "z"
	}

	if b == "E" {
		b = "z"
	}

	if int(a[0]) == int(b[0]) || int(a[0]) == int(b[0])+1 || int(a[0]) < int(b[0]) {
		return true
	}
	return false
}

type vertex struct {
	x, y, dist int
	neighbors  []*vertex
	prev       *vertex
}
