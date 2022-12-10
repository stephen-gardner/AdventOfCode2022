package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func dir(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func getUpdatedPosition(ay, ax, by, bx int) (int, int) {
	if abs(ay-by) > 1 || abs(ax-bx) > 1 {
		by += dir(ay, by)
		bx += dir(ax, bx)
	}
	return by, bx
}

func move(tailPositions map[[2]int]bool, y, x []int, dy, dx, n int) {
	for n > 0 {
		y[0] += dy
		x[0] += dx
		for i := 1; i < len(y); i++ {
			y[i], x[i] = getUpdatedPosition(y[i-1], x[i-1], y[i], x[i])
		}
		tailPos := [2]int{y[len(y)-1], x[len(x)-1]}
		tailPositions[tailPos] = true
		n--
	}
}

func simulateRope(lines []string, knots int) int {
	tailPositions := map[[2]int]bool{{0, 0}: true}
	y := make([]int, knots)
	x := make([]int, knots)
	for _, line := range lines {
		dir := line[0]
		n, _ := strconv.Atoi(line[2:])
		switch dir {
		case 'U':
			move(tailPositions, y, x, -1, 0, n)
		case 'D':
			move(tailPositions, y, x, 1, 0, n)
		case 'L':
			move(tailPositions, y, x, 0, -1, n)
		case 'R':
			move(tailPositions, y, x, 0, 1, n)
		}
	}
	return len(tailPositions)
}

func main() {
	data, _ := os.ReadFile("input")
	lines := strings.Split(string(data[:len(data)-1]), "\n")
	fmt.Println("Part 1:", simulateRope(lines, 2))  // Expected: 5960
	fmt.Println("Part 2:", simulateRope(lines, 10)) // Expected: 2327
}
