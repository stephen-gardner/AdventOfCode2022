package main

import (
	"fmt"
	"os"
	"strings"
)

func getPriorities(s string) []int {
	p := make([]int, 52)
	for _, c := range s {
		if c >= 'a' {
			p[c-'a']++
		} else {
			p[26+(c-'A')]++
		}
	}
	return p
}

func part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		c1, c2 := line[:len(line)/2], line[len(line)/2:]
		p1, p2 := getPriorities(c1), getPriorities(c2)
		for i := range p1 {
			if p1[i] > 0 && p2[i] > 0 {
				sum += 1 + i
			}
		}
	}
	return sum
}

func part2(lines []string) int {
	sum := 0
	for i := 0; i+2 < len(lines); i += 3 {
		l1, l2, l3 := lines[i], lines[i+1], lines[i+2]
		p1, p2, p3 := getPriorities(l1), getPriorities(l2), getPriorities(l3)
		for i := range p1 {
			if p1[i] > 0 && p2[i] > 0 && p3[i] > 0 {
				sum += 1 + i
				continue
			}
		}
	}
	return sum
}

func main() {
	data, _ := os.ReadFile("input")
	lines := strings.Split(string(data), "\n")
	fmt.Println("Part 1:", part1(lines)) // Expected: 7716
	fmt.Println("Part 2:", part2(lines)) // Expected: 2973
}
