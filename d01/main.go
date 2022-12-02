package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Part 1
func findLargestLoad(lines []string) int {
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

// Part 2
func findLargestLoadOfTopThree(lines []string) int {
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
	fmt.Println("Largest load:", findLargestLoad(lines))
	fmt.Println("Largest load of top three:", findLargestLoadOfTopThree(lines))
}
