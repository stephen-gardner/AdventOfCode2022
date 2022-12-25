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
	expected := "2=-1=0"
	if res := part1(testData); res != expected {
		t.Fatalf("\"%s\" != \"%s\" (expected)", res, expected)
	}
}
