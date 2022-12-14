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

func testCave(t *testing.T, cave *Cave, expected []string, expectedDropX int) {
	fail := false ||
		cave.minY != 0 ||
		cave.minX != 0 ||
		cave.maxY != len(expected)-1 ||
		cave.maxX != len(expected[0])-1 ||
		cave.dropX != expectedDropX
	if !fail {
		for i := range expected {
			if string(cave.grid[i]) != expected[i] {
				fail = true
				break
			}
		}
	}
	if fail {
		t.Log("Result:")
		t.Log("\tdropX:", cave.dropX)
		t.Logf("\tminY: %d\tmaxY: %d", cave.minY, cave.maxY)
		t.Logf("\tminX: %d\tmaxX: %d", cave.minX, cave.maxX)
		for _, row := range cave.grid {
			t.Logf("\t%s", string(row))
		}
		t.Log("Expected:")
		t.Log("\tdropX:", expectedDropX)
		t.Logf("\tminY: %d\tmaxY: %d", 0, len(expected)-1)
		t.Logf("\tminX: %d\tmaxX: %d", 0, len(expected[0])-1)
		for _, row := range expected {
			t.Logf("\t%s", row)
		}
		t.FailNow()
	}
}

func TestGetCave(t *testing.T) {
	expected := []string{
		"..........",
		"..........",
		"..........",
		"..........",
		"....#...##",
		"....#...#.",
		"..###...#.",
		"........#.",
		"........#.",
		"#########.",
		"..........",
		"##########",
	}
	testCave(t, getCave(testData, false), expected[:len(expected)-2], 6)
	testCave(t, getCave(testData, true), expected, 6)
}

func testPart(t *testing.T, f func(*Cave) int, data []string, withFloor bool, expected int) {
	cave := getCave(data, withFloor)
	if res := f(cave); res != expected {
		for _, line := range data {
			t.Log(line)
		}
		for _, row := range cave.grid {
			t.Log(string(row))
		}
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestPart1(t *testing.T) {
	testPart(t, part1, testData, false, 24)
}

func TestPart2(t *testing.T) {
	testPart(t, part2, testData, true, 93)
}
