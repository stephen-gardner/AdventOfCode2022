package main

import (
	"os"
	"strings"
	"testing"
)

func getTestData() []string {
	data, _ := os.ReadFile("input_test")
	return strings.Split(string(data[:len(data)-1]), "\n")
}

func TestGetMonkeys(t *testing.T) {
	testData := getTestData()
	res := getMonkeys(testData).String()
	expected := strings.Join(testData, "\n")
	if res != expected {
		t.Log("  Result:", "\n"+res)
		t.Log("Expected:", "\n"+expected)
		t.FailNow()
	}
}

func testPart(t *testing.T, f func([]string) int, expected int) {
	testData := getTestData()
	if res := f(testData); res != expected {
		t.Log("Input:", "\n"+strings.Join(testData, "\n"))
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestPart1(t *testing.T) {
	testPart(t, part1, 10605)
}

func TestPart2(t *testing.T) {
	testPart(t, part2, 2713310158)
}
