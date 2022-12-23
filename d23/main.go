package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

const paddingSize = 5
const (
	North = iota
	South
	West
	East
	NDirs
)

type Field struct {
	grid   [][]byte
	facing int
	round  int
}

func newField(data []byte) *Field {
	dupe := make([]byte, len(data))
	copy(dupe, data)
	return &Field{
		grid:   bytes.Split(dupe, []byte("\n")),
		facing: North,
		round:  1,
	}
}

func (f *Field) acceptProposals(elves map[[2]int]bool, proposals map[[2]int][][2]int) bool {
	for elf := range elves {
		f.grid[elf[0]][elf[1]] = '.'
	}
	hasMovement := false
	for pos, elvesProposing := range proposals {
		if len(elvesProposing) > 1 {
			continue
		}
		f.grid[pos[0]][pos[1]] = '#'
		delete(elves, elvesProposing[0])
		hasMovement = true
	}
	for elf := range elves {
		f.grid[elf[0]][elf[1]] = '#'
	}
	return hasMovement
}

func (f *Field) diffuseElves(rounds int) int {
	type Proposal struct {
		pos [2]int
		ok  bool
	}
	findElves := func(grid [][]byte) map[[2]int]bool {
		elves := map[[2]int]bool{}
		for y := range grid {
			for x := range grid[y] {
				if grid[y][x] == '#' {
					elves[[2]int{y, x}] = true
				}
			}
		}
		return elves
	}
	for rounds == -1 || f.round <= rounds {
		f.padGrid()
		elves := findElves(f.grid)
		proposals := map[[2]int][][2]int{}
		for elf := range elves {
			y, x := elf[0], elf[1]
			options := [4]Proposal{
				{[2]int{y - 1, x}, f.grid[y-1][x] == '.' && f.grid[y-1][x-1] == '.' && f.grid[y-1][x+1] == '.'},
				{[2]int{y + 1, x}, f.grid[y+1][x] == '.' && f.grid[y+1][x-1] == '.' && f.grid[y+1][x+1] == '.'},
				{[2]int{y, x - 1}, f.grid[y][x-1] == '.' && f.grid[y+1][x-1] == '.' && f.grid[y-1][x-1] == '.'},
				{[2]int{y, x + 1}, f.grid[y][x+1] == '.' && f.grid[y+1][x+1] == '.' && f.grid[y-1][x+1] == '.'},
			}
			diffused := true
			for _, opt := range options {
				if !opt.ok {
					diffused = false
					break
				}
			}
			if diffused {
				continue
			}
			for i := 0; i < NDirs; i++ {
				if p := options[(f.facing+i)%NDirs]; p.ok {
					proposals[p.pos] = append(proposals[p.pos], elf)
					break
				}
			}
		}
		if !f.acceptProposals(elves, proposals) {
			break
		}
		f.facing = (f.facing + 1) % NDirs
		f.round++
	}
	f.fitGridToElves()
	return f.freeSpace()
}

func (f *Field) fitGridToElves() {
	minY, maxY, minX, maxX := f.getElfBounds()
	f.grid = f.grid[minY : 1+maxY]
	for y := range f.grid {
		f.grid[y] = f.grid[y][minX : 1+maxX]
	}
}

func (f *Field) freeSpace() int {
	empty := 0
	for y := range f.grid {
		for x := range f.grid[y] {
			if f.grid[y][x] == '.' {
				empty++
			}
		}
	}
	return empty
}

func (f *Field) getElfBounds() (minY, maxY, minX, maxX int) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	minY, maxY = len(f.grid)-1, 0
	minX, maxX = len(f.grid[0])-1, 0
	for y := range f.grid {
		for x := range f.grid[y] {
			if f.grid[y][x] == '#' {
				minY, maxY = min(minY, y), max(maxY, y)
				minX, maxX = min(minX, x), max(maxX, x)
			}
		}
	}
	return
}

func (f *Field) padGrid() {
	minY, maxY, minX, maxX := f.getElfBounds()
	padTop := minY == 0
	padBot := maxY == len(f.grid)-1
	padLft := minX == 0
	padRht := maxX == len(f.grid[0])-1
	if padLft || padRht {
		padding := bytes.Repeat([]byte("."), paddingSize)
		for y := range f.grid {
			if padLft && padRht {
				f.grid[y] = append(append(padding, f.grid[y]...), padding...)
			} else if padLft {
				f.grid[y] = append(padding, f.grid[y]...)
			} else {
				f.grid[y] = append(f.grid[y], padding...)
			}
		}
	}
	getVertPadding := func(width int) [][]byte {
		padding := make([][]byte, paddingSize)
		for i := range padding {
			padding[i] = bytes.Repeat([]byte("."), width)
		}
		return padding
	}
	if padTop {
		f.grid = append(getVertPadding(len(f.grid[0])), f.grid...)
	}
	if padBot {
		f.grid = append(f.grid, getVertPadding(len(f.grid[0]))...)
	}
}

func (f *Field) print() {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintln("End of round", f.round))
	for _, row := range f.grid {
		sb.WriteString(string(row))
		sb.WriteByte('\n')
	}
	fmt.Println(sb.String())
}

func part1(f *Field) int {
	return f.diffuseElves(10)
}

func part2(f *Field) int {
	f.diffuseElves(-1)
	return f.round
}

func main() {
	data, _ := os.ReadFile("input")
	data = data[:len(data)-1]
	fmt.Println("Part 1:", part1(newField(data))) // Expected: 3766
	fmt.Println("Part 2:", part2(newField(data))) // Expected: 954
}
