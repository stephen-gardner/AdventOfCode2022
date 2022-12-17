package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type (
	Valve struct {
		connected map[*Valve]int
		id        int
		flow      int
	}
	State struct {
		node *Valve
		mins int
	}
	Volcano struct {
		me       State
		helper   State
		released int
		opened   int
	}
)

func findDistance(valves map[string]*Valve, conns map[*Valve][]string, src, dst *Valve) int {
	type Location struct {
		node *Valve
		mins int
	}
	queue := []Location{{
		node: src,
		mins: 0,
	}}
	visited := map[*Valve]bool{}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.node == dst {
			return curr.mins
		}
		if visited[curr.node] {
			continue
		}
		visited[curr.node] = true
		for _, label := range conns[curr.node] {
			queue = append(queue, Location{valves[label], 1 + curr.mins})
		}
	}
	return -1
}

func getValves(lines []string) map[string]*Valve {
	valves := make(map[string]*Valve, len(lines))
	conns := make(map[*Valve][]string, len(lines))
	// conns is used solely for calculating distances between valves
	for i := range lines {
		data := strings.Split(lines[i], " ")
		curr := &Valve{
			connected: make(map[*Valve]int),
			id:        1 << i,
		}
		valves[data[1]] = curr
		dataFlow := strings.Split(data[4], "=")
		curr.flow, _ = strconv.Atoi(dataFlow[1][:len(dataFlow[1])-1])
		for _, connected := range data[9:] {
			conns[curr] = append(conns[curr], connected[:2])
		}
	}
	for src := range conns {
		for _, dst := range valves {
			// No need to visit valves with no flow
			if dst.flow > 0 {
				src.connected[dst] = findDistance(valves, conns, src, dst)
			}
		}
	}
	return valves
}

func part1(valves map[string]*Valve) int {
	best := 0
	stack := []Volcano{{
		me: State{valves["AA"], 30},
	}}
	visited := map[[2]int]bool{}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		key := [2]int{curr.released, curr.opened}
		if visited[key] {
			continue
		}
		visited[key] = true
		prevStackSize := len(stack)
		state := curr.me
		for next, dist := range state.node.connected {
			if key[1]&next.id == 0 {
				timeLeft := state.mins - dist - 1
				if timeLeft >= 0 {
					curr.me = State{node: next, mins: timeLeft}
					curr.released = key[0] + (timeLeft * next.flow)
					curr.opened = key[1] | next.id
					stack = append(stack, curr)
				}
			}
		}
		if len(stack) == prevStackSize && curr.released > best {
			best = curr.released
		}
	}
	return best
}

func part2(valves map[string]*Valve) int {
	best := 0
	stack := []Volcano{{
		me:     State{valves["AA"], 26},
		helper: State{valves["AA"], 26},
	}}
	visited := map[[2]int]bool{}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		key := [2]int{curr.released, curr.opened}
		if visited[key] {
			continue
		}
		visited[key] = true
		prevStackSize := len(stack)
		// Proceed with whichever State has the most time remaining
		target := &curr.me
		if curr.helper.mins > curr.me.mins {
			target = &curr.helper
		}
		state := *target
		for next, dist := range state.node.connected {
			if key[1]&next.id == 0 {
				timeLeft := state.mins - dist - 1
				if timeLeft >= 0 {
					*target = State{node: next, mins: timeLeft}
					curr.released = key[0] + (timeLeft * next.flow)
					curr.opened = key[1] | next.id
					stack = append(stack, curr)
				}
			}
		}
		if len(stack) == prevStackSize && curr.released > best {
			best = curr.released
		}
	}
	return best
}

func main() {
	data, _ := os.ReadFile("input")
	lines := strings.Split(string(data[:len(data)-1]), "\n")
	valves := getValves(lines)
	fmt.Println("Part 1:", part1(valves)) // Expected: 1647
	fmt.Println("Part 2:", part2(valves)) // Expected: 2169
}
