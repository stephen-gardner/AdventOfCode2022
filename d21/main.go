package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const NOP = 0

type Monkey struct {
	label string
	lhs   *Monkey
	rhs   *Monkey
	value int
	op    byte
}

func getMonkeys(lines []string) map[string]*Monkey {
	monkeys := map[string]*Monkey{}
	argsMap := map[string][2]string{}
	for _, line := range lines {
		data := strings.Split(line, ": ")
		m := &Monkey{label: data[0]}
		if n, err := strconv.Atoi(data[1]); err == nil {
			m.value = n
			m.op = NOP
		} else {
			calc := strings.Split(data[1], " ")
			m.op = calc[1][0]
			argsMap[m.label] = [2]string{calc[0], calc[2]}
		}
		monkeys[m.label] = m
	}
	for label, args := range argsMap {
		m := monkeys[label]
		m.lhs = monkeys[args[0]]
		m.rhs = monkeys[args[1]]
	}
	return monkeys
}

func (m *Monkey) calculate() int {
	if m == nil {
		return 0
	}
	m.lhs.calculate()
	m.rhs.calculate()
	switch m.op {
	case '=':
		if m.lhs.value == m.rhs.value {
			return 1
		}
		return 0
	case '+':
		m.value = m.lhs.value + m.rhs.value
	case '-':
		m.value = m.lhs.value - m.rhs.value
	case '*':
		m.value = m.lhs.value * m.rhs.value
	case '/':
		m.value = m.lhs.value / m.rhs.value
	}
	return m.value
}

func (m *Monkey) reverseCalc(res *int) {
	if m.op == NOP {
		if m.label == "humn" {
			*res = m.value
		}
		return
	}
	switch m.op {
	case '+':
		m.lhs.value, m.rhs.value = m.value-m.rhs.value, m.value-m.lhs.value
	case '-':
		m.lhs.value, m.rhs.value = m.value+m.rhs.value, m.lhs.value-m.value
	case '*':
		m.lhs.value, m.rhs.value = m.value/m.rhs.value, m.value/m.lhs.value
	case '/':
		m.lhs.value, m.rhs.value = m.value*m.rhs.value, m.lhs.value/m.value
	}
	m.lhs.reverseCalc(res)
	m.rhs.reverseCalc(res)
}

func verifyHumanInput(lines []string, n int) bool {
	monkeys := getMonkeys(lines)
	monkeys["root"].op = '='
	monkeys["humn"].op = NOP
	monkeys["humn"].value = n
	return monkeys["root"].calculate() == 1
}

func part1(lines []string) int {
	monkeys := getMonkeys(lines)
	return monkeys["root"].calculate()
}

func part2(lines []string) int {
	res := 0
	monkeys := getMonkeys(lines)
	root := monkeys["root"]
	root.calculate()
	monkeys["humn"].op = NOP
	// Some inputs may require this operation on rhs instead
	root.lhs.value = root.rhs.value
	root.lhs.reverseCalc(&res)
	return res
}

func main() {
	data, _ := os.ReadFile("input")
	lines := strings.Split(string(data[:len(data)-1]), "\n")
	fmt.Println("Part 1:", part1(lines)) // Expected: 353837700405464
	fmt.Println("Part 2:", part2(lines)) // Expected: 3678125408017
}
