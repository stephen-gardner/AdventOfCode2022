package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type (
	Cave struct {
		grid  [][]byte
		minY  int
		maxY  int
		minX  int
		maxX  int
		dropX int
	}

	Rock struct {
		start []int
		end   []int
	}
)

const (
	initialSandDropX = 500
	widenSideWidth   = 10
)

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

// Adjust boundaries and rock locations so that minX = 0
func (cave *Cave) setupBoundaries(rocks []Rock) {
	cave.minX = math.MaxInt
	for _, r := range rocks {
		cave.maxY = max(max(cave.maxY, r.start[0]), r.end[0])
		cave.minX = min(min(cave.minX, r.start[1]), r.end[1])
		cave.maxX = max(max(cave.maxX, r.start[1]), r.end[1])
	}
	for _, r := range rocks {
		r.start[1] -= cave.minX
		r.end[1] -= cave.minX
	}
	cave.dropX = initialSandDropX - cave.minX
	cave.maxX -= cave.minX
	cave.minX = 0
}

func (cave *Cave) setupRocks(rocks []Rock) {
	for _, r := range rocks {
		startY, endY := r.start[0], r.end[0]
		startX, endX := r.start[1], r.end[1]
		if startY > endY {
			startY, endY = endY, startY
		}
		if startX > endX {
			startX, endX = endX, startX
		}
		for startY < endY || startX < endX {
			cave.grid[startY][startX] = '#'
			if startY < endY {
				startY++
			}
			if startX < endX {
				startX++
			}
		}
		cave.grid[endY][endX] = '#'
	}
}

func (cave *Cave) setupGrid(rocks []Rock, hasFloor bool) {
	cave.setupBoundaries(rocks)
	if hasFloor {
		cave.maxY += 2
	}
	cave.grid = make([][]byte, 1+cave.maxY)
	for y := range cave.grid {
		cave.grid[y] = make([]byte, 1+cave.maxX)
		for x := range cave.grid[y] {
			if y == cave.maxY && hasFloor {
				cave.grid[y][x] = '#'
			} else {
				cave.grid[y][x] = '.'
			}
		}
	}
	cave.setupRocks(rocks)
}

func (cave *Cave) widen(x *int) {
	cave.maxX += 2 * widenSideWidth
	cave.dropX += widenSideWidth
	*x += widenSideWidth
	for y := range cave.grid {
		expanded := make([]byte, 1+cave.maxX)
		for i := 0; i < widenSideWidth; i++ {
			expanded[i] = '.'
			expanded[cave.maxX-i] = '.'
		}
		copy(expanded[widenSideWidth:len(expanded)-widenSideWidth], cave.grid[y])
		cave.grid[y] = expanded
	}
	for i := 0; i < widenSideWidth; i++ {
		cave.grid[cave.maxY][i] = '#'
		cave.grid[cave.maxY][cave.maxX-i] = '#'
	}
}

func getPos(data string) []int {
	coords := strings.Split(data, ",")
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])
	return []int{y, x}
}

func getRocks(lines []string) []Rock {
	rocks := []Rock{}
	for _, line := range lines {
		path := strings.Split(line, " -> ")
		for i := 0; i+1 < len(path); i++ {
			r := Rock{
				start: getPos(path[i]),
				end:   getPos(path[i+1]),
			}
			rocks = append(rocks, r)
		}
	}
	return rocks
}

func getCave(lines []string, hasFloor bool) *Cave {
	cave := &Cave{}
	rocks := getRocks(lines)
	cave.setupGrid(rocks, hasFloor)
	return cave
}

func part1(cave *Cave) int {
	units := 0
	finished := false
	for !finished {
		dropped := false
		y, x := 1, cave.dropX
		for !finished && !dropped {
			switch {
			case y > cave.maxY:
				finished = true
			case cave.grid[y][x] == '.':
				y++
			case x-1 < cave.minX:
				finished = true
			case cave.grid[y][x-1] == '.':
				x--
				y++
			case x+1 > cave.maxX:
				finished = true
			case cave.grid[y][x+1] == '.':
				x++
				y++
			default:
				units++
				cave.grid[y-1][x] = 'o'
				dropped = true
			}
		}
	}
	return units
}

func part2(cave *Cave) int {
	units := 0
	finished := false
	for !finished {
		dropped := false
		y, x := 1, cave.dropX
		for !finished && !dropped {
			switch {
			case cave.grid[y][x] == '.':
				y++
			case x-1 < cave.minX || x+1 > cave.maxX:
				cave.widen(&x)
			case cave.grid[y][x-1] == '.':
				x--
				y++
			case cave.grid[y][x+1] == '.':
				x++
				y++
			default:
				units++
				cave.grid[y-1][x] = 'o'
				dropped = true
				if y == 1 && x == cave.dropX {
					finished = true
				}
			}
		}
	}
	return units
}

func main() {
	data, _ := os.ReadFile("input")
	lines := strings.Split(string(data[:len(data)-1]), "\n")
	fmt.Println("Part 1:", part1(getCave(lines, false))) // Expected: 719
	fmt.Println("Part 2:", part2(getCave(lines, true)))  // Expected: 23390
}
