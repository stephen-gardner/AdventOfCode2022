package main

import "testing"

var input = []string{
	"2-4,6-8",
	"2-3,4-5",
	"5-7,7-9",
	"2-8,3-7",
	"6-6,4-6",
	"2-6,4-8",
}

func test(t *testing.T, f func([]string) int, data []string, expected int) {
	if res := f(data); res != expected {
		for _, line := range data {
			t.Log(line)
		}
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestPart1(t *testing.T) {
	test(t, part1, input, 2)
}

func TestPart2(t *testing.T) {
	test(t, part2, input, 4)
}
