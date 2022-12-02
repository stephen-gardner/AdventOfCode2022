package main

import (
	"strings"
	"testing"
)

func makePrintable(data []byte) {
	for i, c := range data {
		switch c {
		case rock:
			data[i] = 'R'
		case paper:
			data[i] = 'P'
		case scissors:
			data[i] = 'S'
		}
	}
}

func TestDecode(t *testing.T) {
	data := "YCYABCXYZ"
	expected := "PSPRPSRPS"
	dupe := make([]byte, len(data))
	for i := range dupe {
		dupe[i] = data[i]
	}
	decode(dupe)
	makePrintable(dupe)
	if string(dupe) != expected {
		t.Log("   Input:", data)
		t.Log("  Result:", string(dupe))
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func test(t *testing.T, f func([]string) int) {
	input := []byte{
		rock, ' ', rock, '\n',
		rock, ' ', paper, '\n',
		rock, ' ', scissors, '\n',
		paper, ' ', rock, '\n',
		paper, ' ', paper, '\n',
		paper, ' ', scissors, '\n',
		scissors, ' ', rock, '\n',
		scissors, ' ', paper, '\n',
		scissors, ' ', scissors, '\n',
	}
	expected := 3 * (pointsHand[rock] + pointsHand[paper] + pointsHand[scissors])
	expected += 3 * (win + draw + lose)
	lines := strings.Split(string(input), "\n")
	if res := f(lines); res != expected {
		for _, line := range lines {
			data := []byte(line)
			makePrintable(data)
			t.Log(string(data))
		}
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestPart1(t *testing.T) {
	test(t, part1)
}

func TestPart2(t *testing.T) {
	test(t, part2)
}
