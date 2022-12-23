package main

import (
	"os"
	"testing"
)

var testData = getTestData()

func getTestData() []byte {
	data, _ := os.ReadFile("input_test")
	return data[:len(data)-1]
}

func TestPart1(t *testing.T) {
	expected := 110
	if res := part1(newField(testData)); res != expected {
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 20
	if res := part2(newField(testData)); res != expected {
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}
