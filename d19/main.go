package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	Ore = iota
	Clay
	Obsidian
	Geode
	NumMineralTypes
)

const (
	BotOre = iota
	BotClay
	BotObsidian
	BotGeode
)

type (
	Blueprint struct {
		id          int
		prices      [NumMineralTypes][NumMineralTypes]int
		maxCosts    [NumMineralTypes]int
		geodesMined int
	}
	Blueprints []*Blueprint
	Factory    struct {
		bots     [NumMineralTypes]int
		minerals [NumMineralTypes]int
		minutes  int
		created  int
	}
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (bp *Blueprint) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf(
		"Blueprint %d:\n",
		bp.id,
	))
	sb.WriteString(fmt.Sprintf(
		"\tEach ore robot costs %d ore.\n",
		bp.prices[BotOre][Ore],
	))
	sb.WriteString(fmt.Sprintf(
		"\tEach clay robot costs %d ore.\n",
		bp.prices[BotClay][Ore],
	))
	sb.WriteString(fmt.Sprintf(
		"\tEach obsidian robot costs %d ore and %d clay.\n",
		bp.prices[BotObsidian][Ore],
		bp.prices[BotObsidian][Clay],
	))
	sb.WriteString(fmt.Sprintf(
		"\tEach geode robot costs %d ore and %d obsidian.",
		bp.prices[BotGeode][Ore],
		bp.prices[BotGeode][Obsidian],
	))
	return sb.String()
}

func (blueprints Blueprints) String() string {
	var sb strings.Builder
	for i, bp := range blueprints {
		sb.WriteString(fmt.Sprint(bp))
		if i < len(blueprints)-1 {
			sb.WriteString("\n\n")
		}
	}
	return sb.String()
}

func setPrices(data string, args ...*int) {
	arr := strings.Split(data, " ")
	for _, str := range arr {
		if n, err := strconv.Atoi(str); err == nil {
			*args[0] = n
			args = args[1:]
			if len(args) == 0 {
				break
			}
		}
	}
}

func getBlueprints(lines []string) Blueprints {
	blueprints := make([]*Blueprint, len(lines))
	for i := range lines {
		bp := &Blueprint{id: i + 1}
		data := strings.Split(lines[i], ". ")
		setPrices(data[0], &bp.prices[BotOre][Ore])
		setPrices(data[1], &bp.prices[BotClay][Ore])
		setPrices(data[2], &bp.prices[BotObsidian][Ore], &bp.prices[BotObsidian][Clay])
		setPrices(data[3], &bp.prices[BotGeode][Ore], &bp.prices[BotGeode][Obsidian])
		for _, bot := range bp.prices {
			for i := range bot {
				bp.maxCosts[i] = max(bp.maxCosts[i], bot[i])
			}
		}
		blueprints[i] = bp
	}

	return blueprints
}

func mineGeode(bp *Blueprint, minutes int) {
	potentials := map[[2]int]int{}
	getPotential := func(geodes, bots, minutes int) int {
		key := [2]int{bots, minutes}
		if _, exists := potentials[key]; !exists {
			potential := 0
			for i := minutes; i > 0; i-- {
				bots++
				potential += bots
			}
			potentials[key] = potential
		}
		return geodes + potentials[key]
	}
	f := Factory{
		bots:    [NumMineralTypes]int{Ore: 1},
		minutes: minutes,
		created: -1,
	}
	stack := []Factory{f}
	visited := map[Factory]bool{}
	for len(stack) > 0 {
		f := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if visited[f] {
			continue
		}
		visited[f] = true
		if f.minutes == 0 {
			if f.minerals[Geode] > bp.geodesMined {
				bp.geodesMined = f.minerals[Geode]
				fmt.Printf(" -> Blueprint %d: New High = %d\n", bp.id, bp.geodesMined)
			}
			continue
		}
		f.minutes--
		for i := range f.minerals {
			f.minerals[i] += f.bots[i]
		}
		if getPotential(f.minerals[Geode], f.bots[BotGeode], f.minutes) < bp.geodesMined {
			continue
		}
		if f.created != -1 {
			f.bots[f.created]++
		}
		f.created = -1
		stack = append(stack, f)
		for i := range f.bots {
			if bp.maxCosts[i] > 0 && f.bots[i] >= bp.maxCosts[i] {
				continue
			}
			affordable := true
			for j := range f.minerals {
				if f.minerals[j] < bp.prices[i][j] {
					affordable = false
					break
				}
			}
			if affordable {
				next := f
				next.created = i
				for j := range bp.prices[i] {
					next.minerals[j] -= bp.prices[i][j]
				}
				stack = append(stack, next)
			}
		}
	}
}

func part1(blueprints Blueprints) int {
	res := 0
	for _, bp := range blueprints {
		mineGeode(bp, 24)
		res += bp.id * bp.geodesMined
		fmt.Printf("Blueprint %d: Quality = %d (mined: %d)\n", bp.id, bp.id*bp.geodesMined, bp.geodesMined)
	}
	return res
}

func part2(blueprints Blueprints) int {
	res := 1
	for _, bp := range blueprints {
		mineGeode(bp, 32)
		res *= bp.geodesMined
		fmt.Printf("Blueprint %d: Mined = %d\n", bp.id, bp.geodesMined)
	}
	return res
}

func main() {
	data, _ := os.ReadFile("input")
	lines := strings.Split(string(data[:len(data)-1]), "\n")
	blueprints := getBlueprints(lines)
	fmt.Println("Part 1:", part1(blueprints))     // Expected: 1389
	fmt.Println("Part 2:", part2(blueprints[:3])) // Expected: 3003
}
