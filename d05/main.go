package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//     [C]             [L]         [T]
//     [V] [R] [M]     [T]         [B]
//     [F] [G] [H] [Q] [Q]         [H]
//     [W] [L] [P] [V] [M] [V]     [F]
//     [P] [C] [W] [S] [Z] [B] [S] [P]
// [G] [R] [M] [B] [F] [J] [S] [Z] [D]
// [J] [L] [P] [F] [C] [H] [F] [J] [C]
// [Z] [Q] [F] [L] [G] [W] [H] [F] [M]
//  1   2   3   4   5   6   7   8   9
//
// I'd rather hardcode this by hand than go through the trouble of writing
//	something to parse it
//
// Input file generated from original with:
//	tail -n +11 input_original | cut -d' ' -f2 -f4 -f6 > input
var stacks = [][]byte{
	{'Z', 'J', 'G'},
	{'Q', 'L', 'R', 'P', 'W', 'F', 'V', 'C'},
	{'F', 'P', 'M', 'C', 'L', 'G', 'R'},
	{'L', 'F', 'B', 'W', 'P', 'H', 'M'},
	{'G', 'C', 'F', 'S', 'V', 'Q'},
	{'W', 'H', 'J', 'Z', 'M', 'Q', 'T', 'L'},
	{'H', 'F', 'S', 'B', 'V'},
	{'F', 'J', 'Z', 'S'},
	{'M', 'C', 'D', 'P', 'F', 'H', 'B', 'T'},
}

func dupe(data [][]byte) [][]byte {
	res := make([][]byte, len(data))
	for i := range res {
		res[i] = make([]byte, len(data[i]))
		copy(res[i], data[i])
	}
	return res
}

func getTop(data [][]byte) string {
	var sb strings.Builder
	for _, stack := range data {
		if len(stack) == 0 {
			sb.WriteByte(' ')
		} else {
			sb.WriteByte(stack[len(stack)-1])
		}
	}
	return sb.String()
}

func getMovements(line string) (int, int, int) {
	data := strings.Split(line, " ")
	n, _ := strconv.Atoi(data[0])
	src, _ := strconv.Atoi(data[1])
	dst, _ := strconv.Atoi(data[2])
	return n, src - 1, dst - 1
}

func part1(crates [][]byte, lines []string) string {
	for _, line := range lines {
		if line == "" {
			continue
		}
		n, src, dst := getMovements(line)
		for n > 0 {
			srcLen := len(crates[src])
			crates[dst] = append(crates[dst], crates[src][srcLen-1])
			crates[src] = crates[src][:srcLen-1]
			n--
		}
	}
	return getTop(crates)
}

func part2(crates [][]byte, lines []string) string {
	for _, line := range lines {
		if line == "" {
			continue
		}
		n, src, dst := getMovements(line)
		srcLen := len(crates[src])
		crates[dst] = append(crates[dst], crates[src][srcLen-n:]...)
		crates[src] = crates[src][:srcLen-n]
	}
	return getTop(crates)
}

func main() {
	data, _ := os.ReadFile("input")
	lines := strings.Split(string(data), "\n")
	fmt.Println("Part 1:", part1(dupe(stacks), lines)) // Expected: "WSFTMRHPP"
	fmt.Println("Part 2:", part2(dupe(stacks), lines)) // Expected: "GSLCMFBRP"
}
