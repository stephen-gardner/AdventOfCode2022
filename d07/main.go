package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	diskSize   = 70000000
	updateSize = 30000000
)

func cd(cwd *string, to string) {
	dir := *cwd
	if to == ".." {
		path := strings.Split(dir, "/")
		path = path[:len(path)-1]
		if len(path) == 1 {
			dir = "/"
		} else {
			dir = strings.Join(path, "/")
		}
	} else {
		if len(dir) > 1 {
			dir += "/"
		}
		dir += to
	}
	*cwd = dir
}

func getFiles(lines []string) map[string]int {
	dir := ""
	files := map[string]int{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		args := strings.Split(line, " ")
		if args[0] == "$" {
			if args[1] == "cd" {
				cd(&dir, args[2])
			}
		} else if args[0] != "dir" {
			fileSize, _ := strconv.Atoi(args[0])
			fileName := args[1]
			path := dir
			if len(path) > 1 {
				path += "/"
			}
			path += fileName
			files[path] = fileSize
		}
	}
	return files
}

func getFolders(files map[string]int) map[string]int {
	folders := map[string]int{}
	for f, sz := range files {
		path := strings.Split(f, "/")
		path = path[:len(path)-1]
		for len(path) > 0 {
			curr := strings.Join(path, "/")
			if curr == "" {
				curr = "/"
			}
			folders[curr] += sz
			path = path[:len(path)-1]
		}
	}
	return folders
}

func part1(lines []string) int {
	sum := 0
	files := getFiles(lines)
	folders := getFolders(files)
	for _, sz := range folders {
		if sz <= 100000 {
			sum += sz
		}
	}
	return sum
}

func part2(lines []string) int {
	files := getFiles(lines)
	folders := getFolders(files)
	diskUsed := folders["/"]
	diskRemaining := diskSize - diskUsed
	diskNeeded := updateSize - diskRemaining
	smallest := diskSize
	for _, sz := range folders {
		if sz >= diskNeeded && sz < smallest {
			smallest = sz
		}
	}
	return smallest
}

func main() {
	data, _ := os.ReadFile("input")
	lines := strings.Split(string(data), "\n")
	fmt.Println("Part 1:", part1(lines)) // Expected: 2061777
	fmt.Println("Part 2:", part2(lines)) // Expected: 4473403
}
