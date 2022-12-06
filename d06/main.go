package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	startOfPacketMarkerSize  = 4
	startOfMessageMarkerSize = 14
)

func findMarker(data string, size int) int {
	index := map[byte]int{}
	start := 0
	for i, c := range data {
		if idx, exists := index[byte(c)]; exists {
			for start <= idx {
				delete(index, data[start])
				start++
			}
		}
		index[byte(c)] = i
		if len(index) == size {
			break
		}
	}
	return start
}

func part1(data string) int {
	return findMarker(data, startOfPacketMarkerSize) + startOfPacketMarkerSize
}

func part2(data string) int {
	return findMarker(data, startOfMessageMarkerSize) + startOfMessageMarkerSize
}

func main() {
	data, _ := os.ReadFile("input")
	lines := strings.Split(string(data), "\n")
	fmt.Println("Part 1:", part1(lines[0])) // Expected: 1623
	fmt.Println("Part 2:", part2(lines[0])) // Expected: 3774
}
