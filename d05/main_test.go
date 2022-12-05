package main

import (
	"fmt"
	"strings"
	"testing"
)

// [D]
// [N] [C]
// [Z] [M] [P]
//  1   2   3
var testStacks = [][]byte{
	{'Z', 'N'},
	{'M', 'C', 'D'},
	{'P'},
}

// move 1 from 2 to 1
// move 3 from 1 to 3
// move 2 from 2 to 1
// move 1 from 1 to 2
var testOps = []string{
	"1 2 1",
	"3 1 3",
	"2 2 1",
	"1 1 2",
}

func printStacks(t *testing.T, label string, stacks [][]byte) {
	var sb strings.Builder
	high := 0
	for i := range stacks {
		n := len(stacks[i])
		if n-1 > high {
			high = n - 1
		}
	}
	sb.WriteString("\n" + label + "\n\t")
	for high >= 0 {
		for i := range stacks {
			if high < len(stacks[i]) {
				sb.WriteByte(stacks[i][high])
			} else {
				sb.WriteByte(' ')
			}
		}
		sb.WriteString("\n\t")
		high--
	}
	t.Log(sb.String())
}

func test(t *testing.T, f func([][]byte, []string) string, expected string) {
	dupeStacks := dupe(testStacks)
	if res := f(dupeStacks, testOps); res != expected {
		printStacks(t, "Stack:", testStacks)
		var sb strings.Builder
		sb.WriteString("\nOperations:\n")
		for _, op := range testOps {
			sb.WriteString(fmt.Sprintf("\t%s\n", op))
		}
		t.Log(sb.String())
		printStacks(t, "Result:", dupeStacks)
		t.Fatalf(`"%s" != "%s"`, res, expected)
	}
}

func TestPart1(t *testing.T) {
	test(t, part1, "CMZ")
}

func TestPart2(t *testing.T) {
	test(t, part2, "MCD")
}
