package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

var testData = getTestData()

func getTestData() []string {
	data, _ := os.ReadFile("input_test")
	return strings.Split(string(data[:len(data)-1]), "\n")
}

func TestGetPacket(t *testing.T) {
	for _, line := range testData {
		if line == "" {
			continue
		}
		packet := getPacket(line)
		packetStr := fmt.Sprint(packet)
		packetStr = strings.ReplaceAll(packetStr, " ", ",")
		if packetStr != line {
			t.Log("  Result:", packetStr)
			t.Log("Expected:", line)
			t.FailNow()
		}
	}
}

func testPart(t *testing.T, f func([]string) int, expected int) {
	if res := f(testData); res != expected {
		for _, line := range testData {
			t.Log(line)
		}
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestPart1(t *testing.T) {
	testPart(t, part1, 13)
}

func TestPart2(t *testing.T) {
	testPart(t, part2, 140)
}
