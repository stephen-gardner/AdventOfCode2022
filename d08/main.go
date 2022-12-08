package main

import (
	"fmt"
	"os"
	"strings"
)

var dirs = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func getGrid(lines []string) [][]byte {
	grid := make([][]byte, len(lines))
	for y := range lines {
		grid[y] = make([]byte, len(lines[y]))
		for x := range lines[y] {
			grid[y][x] = lines[y][x] - '0'
		}
	}
	return grid
}

func isVisible(grid [][]byte, y, x, deltaY, deltaX int) bool {
	height := grid[y][x]
	for {
		y += deltaY
		x += deltaX
		if y < 0 || y == len(grid) || x < 0 || x == len(grid[y]) {
			break
		}
		if grid[y][x] >= height {
			return false
		}
	}
	return true
}

func part1(grid [][]byte) int {
	visible := 0
	for y := range grid {
		for x := range grid[y] {
			for _, dir := range dirs {
				if isVisible(grid, y, x, dir[0], dir[1]) {
					visible++
					break
				}
			}
		}
	}
	return visible
}

func scoreDir(grid [][]byte, y, x, deltaY, deltaX int) int {
	height := grid[y][x]
	visible := 0
	for {
		y += deltaY
		x += deltaX
		if y < 0 || y == len(grid) || x < 0 || x == len(grid[y]) {
			break
		}
		visible++
		if grid[y][x] >= height {
			break
		}
	}
	return visible
}

func part2(grid [][]byte) int {
	highScore := 0
	for y := range grid {
		for x := range grid[y] {
			score := 1
			for _, dir := range dirs {
				score *= scoreDir(grid, y, x, dir[0], dir[1])
			}
			if score > highScore {
				highScore = score
			}
		}
	}
	return highScore
}

func main() {
	data, _ := os.ReadFile("input")
	lines := strings.Split(string(data), "\n")
	grid := getGrid(lines[:len(lines)-1])
	fmt.Println("Part 1:", part1(grid)) // Expected: 1693
	fmt.Println("Part 2:", part2(grid)) // Expected: 422059
}
