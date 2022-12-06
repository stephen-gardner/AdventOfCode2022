package main

import (
	"testing"
)

var testData = []string{
	"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
	"bvwbjplbgvbhsrlpgdmjqwftvncz",
	"nppdvjthqldpwncqszvftbrmjlhg",
	"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
	"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
}

func test(t *testing.T, f func(string) int, data []string, expected []int) {
	for i := range data {
		if res := f(testData[i]); res != expected[i] {
			t.Log("Input:", testData[i])
			t.Fatalf("%d != %d (expected)", res, expected)
		}
	}
}

func TestPart1(t *testing.T) {
	test(t, part1, testData, []int{7, 5, 6, 10, 11})
}

func TestPart2(t *testing.T) {
	test(t, part2, testData, []int{19, 23, 23, 29, 26})
}
