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

func testPart(t *testing.T, f func(map[string]*Valve) int, expected int) {
	valves := getValves(testData)
	if res := f(valves); res != expected {
		for _, line := range testData {
			t.Log(line)
		}
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestPart1(t *testing.T) {
	testPart(t, part1, 1651)
}

func TestPart2(t *testing.T) {
	testPart(t, part2, 1707)
}
