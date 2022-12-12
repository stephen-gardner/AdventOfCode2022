package main

import (
	"os"
	"strings"
	"testing"
)

var testData = getTestData()

func getTestData() []string {
	data, _ := os.ReadFile("input_test")
	return strings.Split(string(data[:len(data)-1]), "\n")
}

func testPart(t *testing.T, f func([]string) int, expected int) {
	if res := f(testData); res != expected {
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestPart1(t *testing.T) {
	testPart(t, part1, 31)
}

func TestPart2(t *testing.T) {
	testPart(t, part2, 29)
}
