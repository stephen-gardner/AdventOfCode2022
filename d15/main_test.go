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

func TestPart1(t *testing.T) {
	expected := 26
	if res := part1(testData, 10); res != expected {
		for _, line := range testData {
			t.Log(line)
		}
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 56000011
	if res := part2(testData); res != expected {
		for _, line := range testData {
			t.Log(line)
		}
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}
