package main

import (
	"fmt"
	"os"
	"strings"
)

type (
	Blizzard struct {
		y   int
		x   int
		dir byte
	}
	Grid     [][]byte
	Point    [2]int
	Position struct {
		y     int
		x     int
		mins  int
		goals int
	}
	State map[Point]bool
)

var dirs = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {0, 0}}

func getGrid(lines []string) Grid {
	grid := make([][]byte, len(lines))
	for i := range lines {
		grid[i] = []byte(lines[i])
	}
	return grid
}

func printGrid(grid [][]byte) {
	var sb strings.Builder
	for _, row := range grid {
		sb.WriteString(string(row))
		sb.WriteByte('\n')
	}
	fmt.Println(sb.String())
}

func clearBlizzards(grid [][]byte) {
	for y := range grid {
		for x := range grid[y] {
			if !(grid[y][x] == '.' || grid[y][x] == '#') {
				grid[y][x] = '.'
			}
		}
	}
}

func getBlizzards(grid [][]byte) []Blizzard {
	blizzards := []Blizzard{}
	for y := range grid {
		for x := range grid[y] {
			dir := grid[y][x]
			switch dir {
			case '<', '>', '^', 'v':
				blizzards = append(blizzards, Blizzard{y, x, dir})
			}
		}
	}
	return blizzards
}

func hashBlizzards(blizzards []Blizzard) int {
	// djb2: http://www.cse.yorku.ca/~oz/hash.html
	hash := 5381
	for _, b := range blizzards {
		for _, c := range [3]int{b.y, b.x, int(b.dir)} {
			hash = (hash * 33) ^ c
		}
	}
	return hash
}

// Simulates the blizzard up to the point it cycles and returns an array of
// maps containing traversible locations per minute
func (grid Grid) getStates() []State {
	height, width := len(grid), len(grid[0])
	getState := func(grid [][]byte) State {
		state := State{}
		for y := range grid {
			for x := range grid[y] {
				if grid[y][x] == '.' {
					state[Point{y, x}] = true
				}
			}
		}
		return state
	}
	states := []State{}
	blizzards := getBlizzards(grid)
	visited := map[int]bool{}
	for {
		hash := hashBlizzards(blizzards)
		if visited[hash] {
			break
		}
		visited[hash] = true
		next := []Blizzard{}
		clearBlizzards(grid)
		for _, b := range blizzards {
			switch b.dir {
			case '^':
				b.y--
				if grid[b.y][b.x] == '#' {
					b.y = height - 2
				}
			case 'v':
				b.y++
				if grid[b.y][b.x] == '#' {
					b.y = 1
				}
			case '<':
				b.x--
				if grid[b.y][b.x] == '#' {
					b.x = width - 2
				}
			case '>':
				b.x++
				if grid[b.y][b.x] == '#' {
					b.x = 1
				}
			}
			grid[b.y][b.x] = b.dir
			next = append(next, b)
		}
		states = append(states, getState(grid))
		blizzards = next
	}
	return states
}

func (grid Grid) traverse(trips int) int {
	states := grid.getStates()
	startY, startX, endY, endX := 0, 1, len(grid)-1, len(grid[0])-2
	queue := []Position{{startY, startX, 0, 0}}
	visited := map[Position]bool{}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if visited[curr] {
			continue
		}
		visited[curr] = true
		if curr.goals >= 0 && curr.y == endY && curr.x == endX {
			if curr.goals++; curr.goals == trips {
				return curr.mins
			}
			curr.goals = -curr.goals
		}
		if curr.goals < 0 && curr.y == startY && curr.x == startX {
			curr.goals = -curr.goals
		}
		state := states[curr.mins%len(states)]
		for _, dir := range dirs {
			pos := Point{curr.y + dir[0], curr.x + dir[1]}
			if _, exists := state[pos]; exists {
				queue = append(queue, Position{
					y:     pos[0],
					x:     pos[1],
					mins:  curr.mins + 1,
					goals: curr.goals,
				})
			}
		}
	}
	return -1
}

func part1(lines []string) int {
	grid := getGrid(lines)
	return grid.traverse(1)
}

func part2(lines []string) int {
	grid := getGrid(lines)
	return grid.traverse(2)
}

func main() {
	data, _ := os.ReadFile("input")
	lines := strings.Split(string(data[:len(data)-1]), "\n")
	fmt.Println("Part 1:", part1(lines)) // Expected: 332
	fmt.Println("Part 2:", part2(lines)) // Expected: 942
}
