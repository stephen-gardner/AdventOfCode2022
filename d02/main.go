package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	rock = iota
	paper
	scissors
)

const (
	lose = 0
	draw = 3
	win  = 6
)

var pointsOutcome = [][]int{
	rock: {
		rock:     draw,
		paper:    win,
		scissors: lose,
	},
	paper: {
		rock:     lose,
		paper:    draw,
		scissors: win,
	},
	scissors: {
		rock:     win,
		paper:    lose,
		scissors: draw,
	},
}

var pointsHand = []int{
	rock:     1,
	paper:    2,
	scissors: 3,
}

func part1(lines []string) int {
	score := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		opponent, me := line[0], line[2]
		score += pointsHand[me]
		score += pointsOutcome[opponent][me]
	}
	return score
}

func part2(lines []string) int {
	// For self, the game must be rigged:
	// 	Rock     = Play to lose
	// 	Paper    = Play to draw
	// 	Scissors = Play to win
	playbook := [][]byte{
		rock: {
			rock:     scissors,
			paper:    rock,
			scissors: paper,
		},
		paper: {
			rock:     rock,
			paper:    paper,
			scissors: scissors,
		},
		scissors: {
			rock:     paper,
			paper:    scissors,
			scissors: rock,
		},
	}
	score := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		opponent, me := line[0], line[2]
		me = playbook[opponent][me]
		score += pointsHand[me]
		score += pointsOutcome[opponent][me]
	}
	return score
}

// Make all symbols consistent
func decode(data []byte) {
	for i := range data {
		switch data[i] {
		case 'A', 'X':
			data[i] = rock
		case 'B', 'Y':
			data[i] = paper
		case 'C', 'Z':
			data[i] = scissors
		}
	}
}

func main() {
	data, _ := os.ReadFile("input")
	decode(data)
	lines := strings.Split(string(data), "\n")
	fmt.Println("Part 1:", part1(lines)) // Expected: 15691
	fmt.Println("Part 2:", part2(lines)) // Expected: 12989
}
