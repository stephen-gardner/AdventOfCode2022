package main

import (
	"strings"
	"testing"
)

var testData = []string{
	"$ cd /",
	"$ ls",
	"dir a",
	"14848514 b.txt",
	"8504156 c.dat",
	"dir d",
	"$ cd a",
	"$ ls",
	"dir e",
	"29116 f",
	"2557 g",
	"62596 h.lst",
	"$ cd e",
	"$ ls",
	"584 i",
	"$ cd ..",
	"$ cd ..",
	"$ cd d",
	"$ ls",
	"4060174 j",
	"8033020 d.log",
	"5626152 d.ext",
	"7214296 k",
}

func testCD(t *testing.T, from, to, expected string) {
	res := from
	cd(&res, to)
	if res != expected {
		t.Log("    From:", from)
		t.Log("      To:", to)
		t.Log("  Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestCD(t *testing.T) {
	testCD(t, "", "/", "/")
	testCD(t, "/a", "b", "/a/b")
	testCD(t, "/a/b/c/d", "..", "/a/b/c")
	testCD(t, "/a", "..", "/")
}

func compareFiles(t *testing.T, res, expected map[string]int) {
	fail := len(res) != len(expected)
	if !fail {
		for path, expectedSize := range expected {
			size, present := res[path]
			fail = !present || size != expectedSize
			if fail {
				break
			}
		}
	}
	if fail {
		t.Log("   Input:", "\n"+strings.Join(testData, "\n"))
		t.Log("  Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestGetFiles(t *testing.T) {
	expected := map[string]int{
		"/a/e/i":   584,
		"/a/f":     29116,
		"/a/g":     2557,
		"/a/h.lst": 62596,
		"/b.txt":   14848514,
		"/c.dat":   8504156,
		"/d/d.ext": 5626152,
		"/d/d.log": 8033020,
		"/d/j":     4060174,
		"/d/k":     7214296,
	}
	res := getFiles(testData)
	compareFiles(t, res, expected)
}

func TestGetFolders(t *testing.T) {
	expected := map[string]int{
		"/":    48381165,
		"/a":   94853,
		"/a/e": 584,
		"/d":   24933642,
	}
	files := getFiles(testData)
	res := getFolders(files)
	compareFiles(t, res, expected)
}

func testPart(t *testing.T, f func([]string) int, expected int) {
	if res := f(testData); res != expected {
		t.Log("Input:", "\n"+strings.Join(testData, "\n"))
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestPart1(t *testing.T) {
	testPart(t, part1, 95437)
}

func TestPart2(t *testing.T) {
	testPart(t, part2, 24933642)
}
