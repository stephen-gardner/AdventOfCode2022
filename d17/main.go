package main

import (
	"fmt"
	"os"
	"strings"
)

type (
	Chamber struct {
		data        [][]byte
		wind        string
		top         [chamberWidth]int
		rockIdx     int
		topY        int
		towerHeight int
		windIdx     int
	}
	CycleData struct {
		key         [2 + chamberWidth]int
		height      int
		heightDelta int
		length      int
		lengthDelta int
		complete    bool
		started     bool
	}
	Rock struct {
		data []string
		rows int
		cols int
	}
)

const (
	chamberWidth       = 7
	rockClearance      = 3
	targetBlocksPlaced = 1000000000000
)

var rocks = []Rock{
	{
		data: []string{
			"####",
		},
		rows: 1,
		cols: 4,
	},
	{
		data: []string{
			".#.",
			"###",
			".#.",
		},
		rows: 3,
		cols: 3,
	}, {
		data: []string{
			"..#",
			"..#",
			"###",
		},
		rows: 3,
		cols: 3,
	}, {
		data: []string{
			"#",
			"#",
			"#",
			"#",
		},
		rows: 4,
		cols: 1,
	}, {
		data: []string{
			"##",
			"##",
		},
		rows: 2,
		cols: 2,
	},
}

func newChamber(wind string) *Chamber {
	ch := &Chamber{wind: wind}
	ch.data = make([][]byte, rockClearance+rocks[0].rows)
	for y := range ch.data {
		ch.data[y] = make([]byte, chamberWidth)
		for x := 0; x < chamberWidth; x++ {
			ch.data[y][x] = '.'
		}
	}
	return ch
}

func (ch *Chamber) print(towerOnly bool) {
	for _, row := range ch.data {
		fmt.Println(string(row))
	}
	if !towerOnly {
		fmt.Println("Top:", ch.top)
		fmt.Println("TopY:", ch.topY)
		fmt.Println("RockIdx:", ch.rockIdx)
		fmt.Println("WindIdx:", ch.windIdx)
		fmt.Println("Tower Height:", ch.towerHeight)
	}
	fmt.Print("\n")
}

func (ch *Chamber) resize() {
	if ch.towerHeight == 0 {
		return
	}
	rowsToAdd := (rockClearance + rocks[ch.rockIdx].rows) - ch.topY
	if rowsToAdd == 0 {
		return
	}
	if rowsToAdd < 0 {
		ch.data = ch.data[-rowsToAdd:]
	} else {
		for i := 0; i < rowsToAdd; i++ {
			ch.data = append(ch.data, nil)
		}
		for y := len(ch.data) - 1; y >= rowsToAdd; y-- {
			ch.data[y] = ch.data[y-rowsToAdd]
		}
		for y := 0; y < rowsToAdd; y++ {
			ch.data[y] = make([]byte, chamberWidth)
			for x := 0; x < chamberWidth; x++ {
				ch.data[y][x] = '.'
			}
		}
	}
	ch.topY += rowsToAdd
	for x := 0; x < chamberWidth; x++ {
		ch.top[x] += rowsToAdd
	}
}

func (ch *Chamber) rockCanMove(y, x int) bool {
	rock := rocks[ch.rockIdx]
	if x < 0 || x+rock.cols > chamberWidth {
		return false
	}
	for rockY := range rock.data {
		if y+rockY == len(ch.data) {
			return false
		}
		for rockX := range rock.data[rockY] {
			if rock.data[rockY][rockX] == '#' && ch.data[y+rockY][x+rockX] != '.' {
				return false
			}
		}
	}
	return true
}

func (ch *Chamber) rockDrop() {
	y, x := 0, 2
	falling, wind := true, true
	for falling {
		if wind {
			switch ch.wind[ch.windIdx] {
			case '>':
				if ch.rockCanMove(y, x+1) {
					x++
				}
			case '<':
				if ch.rockCanMove(y, x-1) {
					x--
				}
			}
			ch.windIdx = (ch.windIdx + 1) % len(ch.wind)
		} else {
			falling = ch.rockCanMove(y+1, x)
			if falling {
				y++
			}
		}
		wind = !wind
	}
	ch.rockSettle(y, x)
	ch.updateTop()
	ch.rockIdx = (ch.rockIdx + 1) % len(rocks)
	ch.resize()
}

func (ch *Chamber) rockSettle(originY, originX int) {
	rData := rocks[ch.rockIdx].data
	for y := range rData {
		for x := range rData[y] {
			if rData[y][x] == '#' {
				ch.data[originY+y][originX+x] = '#'
			}
		}
	}
}

func (ch *Chamber) updateTop() {
	ch.topY = len(ch.data)
	for x := range ch.data[0] {
		for y := range ch.data {
			if ch.data[y][x] == '#' {
				ch.top[x] = y
				if y < ch.topY {
					ch.topY = y
				}
				break
			}
			ch.top[x] = len(ch.data)
		}
	}
	ch.towerHeight = len(ch.data) - ch.topY
}

func simulate(lines []string, times int) (*Chamber, *CycleData) {
	ch := newChamber(lines[0])
	cycle := &CycleData{}
	initialHeight := 0
	initialLength := 0
	visited := map[[2 + chamberWidth]int]int{}
	for i := 0; i < times; i++ {
		key := [2 + chamberWidth]int{ch.rockIdx, ch.windIdx}
		ch.rockDrop()
		if cycle.complete {
			continue
		}
		for i := range ch.top {
			key[2+i] = ch.top[i]
		}
		if !cycle.started {
			if _, exists := visited[key]; exists {
				cycle.key = key
				cycle.started = true
				initialHeight = ch.towerHeight
				initialLength = i + 1
			}
			visited[key] = i + 1
			continue
		}
		sameKey := true
		for i := range cycle.key {
			if cycle.key[i] != key[i] {
				sameKey = false
				break
			}
		}
		if sameKey {
			cycle.height = ch.towerHeight
			cycle.heightDelta = ch.towerHeight - initialHeight
			cycle.length = i + 1
			cycle.lengthDelta = (i + 1) - initialLength
			cycle.complete = true
		}
	}
	return ch, cycle
}

func part1(lines []string) int {
	ch, _ := simulate(lines, 2022)
	return ch.towerHeight
}

func part2(lines []string) int {
	_, cycle := simulate(lines, 4000)
	height := cycle.height
	blocksPlaced := cycle.length
	for blocksPlaced+cycle.lengthDelta <= targetBlocksPlaced {
		height += cycle.heightDelta
		blocksPlaced += cycle.lengthDelta
	}
	blocksNeeded := targetBlocksPlaced - blocksPlaced
	ch, _ := simulate(lines, cycle.length+blocksNeeded)
	return height + (ch.towerHeight - cycle.height)
}

func main() {
	data, _ := os.ReadFile("input")
	lines := strings.Split(string(data[:len(data)-1]), "\n")
	fmt.Println("Part 1:", part1(lines)) // Expected: 3059
	fmt.Println("Part 2:", part2(lines)) // Expected: 1500874635587
}
