package main

import "testing"

func TestPart1(t *testing.T) {
	expected := 33
	data := []string{
		"aaaaaa", //  1 = a(1)
		"aBcBac", // 32 = a(1) + B(28)  + c(3)
		"abcdef", //  0
		"ABCabc", //  0
	}
	if res := part1(data); res != expected {
		t.Log("Input:", data)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 97
	data := []string{
		"oiLSmj", //
		"SeYhjP", // S(45)
		"LfSGCe", //
		"UkWMQm", //
		"TLlIaW", // W(49)
		"DWCGqK", //
		"NYOcPO", //
		"zLdFcQ", // c(3)
		"lcxmTi", //
	}
	if res := part2(data); res != expected {
		t.Log("Input:", data)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}
