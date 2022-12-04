package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getRange(s string) []int {
	ends := strings.Split(s, "-")
	n1, _ := strconv.Atoi(ends[0])
	n2, _ := strconv.Atoi(ends[1])
	return []int{n1, n2}
}

func isContained(p1, p2 []int) bool {
	return p2[0] >= p1[0] && p2[1] <= p1[1]
}

func part1(lines []string) int {
	count := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		pair := strings.Split(line, ",")
		p1 := getRange(pair[0])
		p2 := getRange(pair[1])
		if isContained(p1, p2) || isContained(p2, p1) {
			count++
		}
	}
	return count
}

func hasOverlap(p1, p2 []int) bool {
	return (p2[0] >= p1[0] && p2[0] <= p1[1]) || (p2[1] >= p1[0] && p2[1] <= p1[1])
}

func part2(lines []string) int {
	count := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		pair := strings.Split(line, ",")
		p1 := getRange(pair[0])
		p2 := getRange(pair[1])
		if hasOverlap(p1, p2) || hasOverlap(p2, p1) {
			count++
		}
	}
	return count
}

func main() {
	data, _ := os.ReadFile("input")
	lines := strings.Split(string(data), "\n")
	fmt.Println("Part 1:", part1(lines)) // Expected: 450
	fmt.Println("Part 2:", part2(lines)) // Expected: 837
}
