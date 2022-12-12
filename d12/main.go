package main

import (
	"fmt"
	"os"
	"strings"
)

type (
	Location struct {
		y, x int
	}
	Node struct {
		isEnd     bool
		connected []*Node
	}
	Position struct {
		node  *Node
		steps int
	}
)

var dirs = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func getGrid(lines []string) [][]byte {
	grid := make([][]byte, len(lines))
	for i := range grid {
		grid[i] = []byte(lines[i])
	}
	return grid
}

func getGraph(grid [][]byte) map[Location]*Node {
	nodes := map[Location]*Node{}
	for y := range grid {
		for x := range grid[0] {
			nodes[Location{y, x}] = &Node{
				isEnd: grid[y][x] == 'E',
			}
		}
	}
	for pos, currNode := range nodes {
		currHeight := grid[pos.y][pos.x]
		for _, dir := range dirs {
			dest := Location{pos.y + dir[0], pos.x + dir[1]}
			if destNode, exists := nodes[dest]; exists {
				destHeight := grid[dest.y][dest.x]
				if (destNode.isEnd && currHeight != 'z') || currHeight+1 < destHeight {
					continue
				}
				currNode.connected = append(currNode.connected, destNode)
			}
		}
	}
	return nodes
}

func getStartingPositions(lines []string, startSymbols string) []Position {
	var startingLocations []Location
	grid := getGrid(lines)
	for y := range grid {
		for x := range grid[0] {
			c := grid[y][x]
			if strings.Contains(startSymbols, string(c)) {
				grid[y][x] = 'a'
				startingLocations = append(startingLocations, Location{y, x})
			}
		}
	}
	graph := getGraph(grid)
	positions := make([]Position, len(startingLocations))
	for i := range startingLocations {
		positions[i] = Position{graph[startingLocations[i]], 0}
	}
	return positions
}

func bfs(queue []Position) int {
	visited := map[Position]bool{}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if visited[curr] {
			continue
		}
		visited[curr] = true
		if curr.node.isEnd {
			return curr.steps
		}
		for _, node := range curr.node.connected {
			queue = append(queue, Position{node, curr.steps + 1})
		}
	}
	return -1
}

func part1(lines []string) int {
	return bfs(getStartingPositions(lines, "S"))
}

func part2(lines []string) int {
	return bfs(getStartingPositions(lines, "Sa"))
}

func main() {
	data, _ := os.ReadFile("input")
	lines := strings.Split(string(data[:len(data)-1]), "\n")
	fmt.Println("Part 1:", part1(lines)) // Expected: 534
	fmt.Println("Part 2:", part2(lines)) // Expected: 525
}
