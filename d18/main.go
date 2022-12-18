package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Cube struct {
	x int
	y int
	z int
}

const pad = 10

var dirs = [][]int{
	{1, 0, 0},
	{-1, 0, 0},
	{0, 1, 0},
	{0, -1, 0},
	{0, 0, 1},
	{0, 0, -1},
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getCubes(lines []string) map[[3]int]*Cube {
	cubes := map[[3]int]*Cube{}
	for _, line := range lines {
		data := strings.Split(line, ",")
		c := &Cube{}
		c.x, _ = strconv.Atoi(data[0])
		c.y, _ = strconv.Atoi(data[1])
		c.z, _ = strconv.Atoi(data[2])
		cubes[[3]int{c.x, c.y, c.z}] = c
	}
	return cubes
}

// Brute force: check if there's actually a cube at every position next to
//	the current cube.
func part1(lines []string) int {
	cubes := getCubes(lines)
	faces := 6 * len(cubes)
	for _, c := range cubes {
		for _, dir := range dirs {
			key := [3]int{c.x + dir[0], c.y + dir[1], c.z + dir[2]}
			if _, exists := cubes[key]; exists {
				faces--
			}
		}
	}
	return faces
}

// BFS: by using the delta used to arrive at a position as part of the key for
//	the visited set, we can re-visit a cube to strike it from multiple
//	directions.
func part2(lines []string) int {
	cubes := getCubes(lines)
	minX, maxX := math.MaxInt, math.MinInt
	minY, maxY := math.MaxInt, math.MinInt
	minZ, maxZ := math.MaxInt, math.MinInt
	for _, c := range cubes {
		// Add some padding to the bounds so the cubes are surrounded by air
		minX, maxX = min(minX, c.x-pad), max(maxX, c.x+pad)
		minY, maxY = min(minY, c.y-pad), max(maxY, c.y+pad)
		minZ, maxZ = min(minZ, c.x-pad), max(maxZ, c.x+pad)
	}
	faces := 0
	queue := [][6]int{{minX, minY, minZ}}
	visited := map[[6]int]bool{}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if visited[curr] {
			continue
		}
		visited[curr] = true
		if curr[0] < minX || curr[0] > maxX ||
			curr[1] < minY || curr[1] > maxY ||
			curr[2] < minZ || curr[2] > maxZ {
			continue
		}
		if _, exists := cubes[[3]int{curr[0], curr[1], curr[2]}]; exists {
			faces++
			continue
		}
		for _, dir := range dirs {
			queue = append(queue, [6]int{
				curr[0] + dir[0],
				curr[1] + dir[1],
				curr[2] + dir[2],
				dir[0],
				dir[1],
				dir[2],
			})
		}
	}
	return faces
}

func main() {
	data, _ := os.ReadFile("input")
	lines := strings.Split(string(data[:len(data)-1]), "\n")
	fmt.Println("Part 1:", part1(lines)) // Expected: 4580
	fmt.Println("Part 2:", part2(lines)) // Expected: 2610
}
