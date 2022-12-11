package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	modCommon = -1
	opArgOld  = -1

	opAdd = '+'
	opMul = '*'
)

type (
	Monkey struct {
		items       []int
		opType      int
		opArg       int
		testMod     int
		destTrue    int
		destFalse   int
		inspections int
	}
	Monkeys []*Monkey
)

func getMonkeys(lines []string) Monkeys {
	monkeys := Monkeys{}
	for i := 0; i < len(lines); i += 7 {
		m := &Monkey{}
		m.ParseItems(lines[i+1])
		m.ParseOperation(lines[i+2])
		m.ParseTest(lines[i+3 : i+6])
		monkeys = append(monkeys, m)
	}
	return monkeys
}

func (monkeys Monkeys) GetBusinessLevel() int {
	var b1, b2 int
	for _, m := range monkeys {
		if m.inspections > b2 {
			b2 = m.inspections
		}
		if b2 > b1 {
			b1, b2 = b2, b1
		}
	}
	return b1 * b2
}

func (monkeys Monkeys) GetMod(mod int) func(int) int {
	if mod != modCommon {
		return func(worry int) int {
			return worry / mod
		}
	}
	mod = 1
	for _, m := range monkeys {
		mod *= m.testMod
	}
	return func(worry int) int {
		return worry % mod
	}
}

func (monkeys Monkeys) String() string {
	var sb strings.Builder
	for i := range monkeys {
		sb.WriteString(fmt.Sprintf("Monkey %d:\n%s", i, monkeys[i]))
		if i < len(monkeys)-1 {
			sb.WriteString("\n\n")
		}
	}
	return sb.String()
}

func (m *Monkey) ParseItems(data string) {
	items := strings.Split(data, ":")
	items = strings.Split(items[1], ",")
	for i := range items {
		itemStr := strings.Trim(items[i], " ")
		item, _ := strconv.Atoi(itemStr)
		m.items = append(m.items, item)
	}
}

func (m *Monkey) ParseOperation(data string) {
	op := strings.Split(data, " ")
	switch op[len(op)-2] {
	case "+":
		m.opType = opAdd
	case "*":
		m.opType = opMul
	}
	if op[len(op)-1] == "old" {
		m.opArg = opArgOld
	} else {
		m.opArg, _ = strconv.Atoi(op[len(op)-1])
	}
}

func (m *Monkey) ParseTest(data []string) {
	test := strings.Split(data[0], " ")
	destTrue := strings.Split(data[1], " ")
	destFalse := strings.Split(data[2], " ")
	m.testMod, _ = strconv.Atoi(test[len(test)-1])
	m.destTrue, _ = strconv.Atoi(destTrue[len(destTrue)-1])
	m.destFalse, _ = strconv.Atoi(destFalse[len(destFalse)-1])
}

func (m *Monkey) String() string {
	var sb strings.Builder
	sb.WriteString("  Starting items: ")
	for i := range m.items {
		sb.WriteString(strconv.Itoa(m.items[i]))
		if i < len(m.items)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString(fmt.Sprintf("\n  Operation: new = old %c ", m.opType))
	if m.opArg == opArgOld {
		sb.WriteString("old")
	} else {
		sb.WriteString(strconv.Itoa(m.opArg))
	}
	sb.WriteString(fmt.Sprintf("\n  Test: divisible by %d", m.testMod))
	sb.WriteString(fmt.Sprintf("\n    If true: throw to monkey %d", m.destTrue))
	sb.WriteString(fmt.Sprintf("\n    If false: throw to monkey %d", m.destFalse))
	return sb.String()
}

func goBananas(monkeys Monkeys, numRounds int, mod func(int) int) int {
	for round := 1; round <= numRounds; round++ {
		for _, m := range monkeys {
			for _, worry := range m.items {
				opArg := m.opArg
				if opArg == opArgOld {
					opArg = worry
				}
				switch m.opType {
				case opAdd:
					worry += opArg
				case opMul:
					worry *= opArg
				}
				worry = mod(worry)
				dest := m.destTrue
				if worry%m.testMod != 0 {
					dest = m.destFalse
				}
				monkeys[dest].items = append(monkeys[dest].items, worry)
				m.inspections++
			}
			m.items = nil
		}
	}
	return monkeys.GetBusinessLevel()
}

func part1(lines []string) int {
	monkeys := getMonkeys(lines)
	return goBananas(monkeys, 20, monkeys.GetMod(3))
}

func part2(lines []string) int {
	monkeys := getMonkeys(lines)
	return goBananas(monkeys, 10000, monkeys.GetMod(modCommon))
}

func main() {
	data, _ := os.ReadFile("input")
	lines := strings.Split(string(data[:len(data)-1]), "\n")
	fmt.Printf("%s\n\n", getMonkeys(lines))
	fmt.Println("Part 1:", part1(lines)) // Expected: 117640
	fmt.Println("Part 2:", part2(lines)) // Expected: 30616425600
}
