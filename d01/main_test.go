package main

import "testing"

func TestFindLargestLoad(t *testing.T) {
	expected := 7
	lines := []string{
		"1", "2", "",
		"3", "4", "", // Largest
	}
	if res := findLargestLoad(lines); res != expected {
		t.Log("Input:", lines)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestFindLargestLoadOfTopThree(t *testing.T) {
	expected := 39
	lines := []string{
		"1", "2", "",
		"3", "4", "", // Third
		"4", "5", "6", "", // Second
		"7", "",
		"8", "9", "", // First
		"1", "",
		"1", "1", "1", "",
	}
	if res := findLargestLoadOfTopThree(lines); res != expected {
		t.Log("Input:", lines)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}
