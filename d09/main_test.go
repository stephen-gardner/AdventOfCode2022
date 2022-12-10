package main

import (
	"strings"
	"testing"
)

func test(t *testing.T, data []string, knots, expected int) {
	if res := simulateRope(data, knots); res != expected {
		t.Log("Input:", "\n\t"+strings.Join(data, "\n\t"))
		t.Log("Knots:", knots)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestPart1(t *testing.T) {
	expected := 13
	knots := 2
	testData := []string{
		"R 4",
		"U 4",
		"L 3",
		"D 1",
		"R 4",
		"D 1",
		"L 5",
		"R 2",
	}
	test(t, testData, knots, expected)
}

func TestPart2(t *testing.T) {
	expected := 36
	knots := 10
	testData := []string{
		"R 5",
		"U 8",
		"L 8",
		"D 3",
		"R 17",
		"D 10",
		"L 25",
		"U 20",
	}
	test(t, testData, knots, expected)
}
