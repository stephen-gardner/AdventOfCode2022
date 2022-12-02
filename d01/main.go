package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(lines []string) int {
	var curr, largest int
	for _, line := range lines {
		if line == "" {
			if curr > largest {
				largest = curr
			}
			curr = 0
		} else {
			n, _ := strconv.Atoi(line)
			curr += n
		}
	}
	return largest
}

func part2(lines []string) int {
	var curr, l1, l2, l3 int
	for _, line := range lines {
		if line == "" {
			if curr > l3 {
				l3 = curr
			}
			if l2 < l3 {
				l2, l3 = l3, l2
			}
			if l1 < l2 {
				l1, l2 = l2, l1
			}
			curr = 0
		} else {
			n, _ := strconv.Atoi(line)
			curr += n
		}
	}
	return l1 + l2 + l3
}

func main() {
	data, _ := os.ReadFile("input")
	lines := strings.Split(string(data), "\n")
	fmt.Println("Part 1:", part1(lines)) // Expected: 67633
	fmt.Println("Part 2:", part2(lines)) // Expected: 199628
}
