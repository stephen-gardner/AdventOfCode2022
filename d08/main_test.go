package main

import (
	"strings"
	"testing"
)

var testData = []string{
	"30373",
	"25512",
	"65332",
	"33549",
	"35390",
}

func test(t *testing.T, f func([][]byte) int, input []string, expected int) {
	grid := getGrid(input)
	if res := f(grid); res != expected {
		t.Log("Input:", "\n\t"+strings.Join(input, "\n\t"))
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestPart1(t *testing.T) {
	test(t, part1, testData, 21)
}

func TestPart2(t *testing.T) {
	test(t, part2, testData, 8)
}
