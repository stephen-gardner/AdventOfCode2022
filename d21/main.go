package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const NOP = 0

type Monkey struct {
	value int
	op    byte
	lhs   string
	rhs   string
}

func getMonkeys(lines []string) map[string]*Monkey {
	monkeys := map[string]*Monkey{}
	for _, line := range lines {
		data := strings.Split(line, ": ")
		m := &Monkey{}
		if n, err := strconv.Atoi(data[1]); err == nil {
			m.value = n
			m.op = NOP
		} else {
			calc := strings.Split(data[1], " ")
			m.op = calc[1][0]
			m.lhs = calc[0]
			m.rhs = calc[2]
		}
		monkeys[data[0]] = m
	}
	return monkeys
}

func calculate(monkeys map[string]*Monkey, curr string) int {
	m := monkeys[curr]
	lhs := monkeys[m.lhs]
	if lhs.op != NOP {
		calculate(monkeys, m.lhs)
	}
	rhs := monkeys[m.rhs]
	if rhs.op != NOP {
		calculate(monkeys, m.rhs)
	}
	switch m.op {
	case '=':
		if lhs.value == rhs.value {
			return 1
		} else {
			return 0
		}
	case '+':
		m.value = lhs.value + rhs.value
	case '-':
		m.value = lhs.value - rhs.value
	case '*':
		m.value = lhs.value * rhs.value
	case '/':
		m.value = lhs.value / rhs.value
	}
	return m.value
}

func reverseCalc(res *int, monkeys map[string]*Monkey, curr string) {
	m := monkeys[curr]
	if m.op == NOP {
		if curr == "humn" {
			*res = m.value
		}
		return
	}
	lhs := monkeys[m.lhs]
	rhs := monkeys[m.rhs]
	switch m.op {
	case '+':
		lhs.value, rhs.value = m.value-rhs.value, m.value-lhs.value
	case '-':
		lhs.value, rhs.value = m.value+rhs.value, lhs.value-m.value
	case '*':
		lhs.value, rhs.value = m.value/rhs.value, m.value/lhs.value
	case '/':
		lhs.value, rhs.value = m.value*rhs.value, lhs.value/m.value
	}
	reverseCalc(res, monkeys, m.lhs)
	reverseCalc(res, monkeys, m.rhs)
}

func verifyEqual(lines []string, n int) bool {
	monkeys := getMonkeys(lines)
	monkeys["root"].op = '='
	monkeys["humn"].op = NOP
	monkeys["humn"].value = n
	return calculate(monkeys, "root") == 1
}

func part1(lines []string) int {
	return calculate(getMonkeys(lines), "root")
}

func part2(lines []string) int {
	monkeys := getMonkeys(lines)
	calculate(monkeys, "root")
	root := monkeys["root"]
	// Some inputs may require this operation on rhs instead
	monkeys[root.lhs].value = monkeys[root.rhs].value
	monkeys["humn"].op = NOP
	res := 0
	reverseCalc(&res, monkeys, root.lhs)
	return res
}

func main() {
	data, _ := os.ReadFile("input")
	lines := strings.Split(string(data[:len(data)-1]), "\n")
	fmt.Println("Part 1:", part1(lines)) // Expected: 353837700405464
	fmt.Println("Part 2:", part2(lines)) // Expected: 3678125408017
}
