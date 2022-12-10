package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type (
	Instr struct {
		op      byte
		arg     int
		latency int
	}
	Canvas [][]byte
)

const (
	opAddx = iota
	opNoop

	canvasHeight    = 6
	canvasWidth     = 40
	spriteWidth     = 3
	spriteWidthHalf = spriteWidth / 2
)

func getSchedule(lines []string) []Instr {
	instr := make([]Instr, len(lines))
	for i, line := range lines {
		instr := &instr[i]
		code := strings.Split(line, " ")
		switch code[0] {
		case "addx":
			instr.op = opAddx
			instr.arg, _ = strconv.Atoi(code[1])
			instr.latency = 2
		case "noop":
			instr.op = opNoop
			instr.latency = 1
		}
	}
	return instr
}

func getCanvas() Canvas {
	canvas := make([][]byte, canvasHeight)
	for i := range canvas {
		canvas[i] = make([]byte, canvasWidth)
		for j := range canvas[i] {
			canvas[i][j] = '.'
		}
	}
	return canvas
}

func (canvas Canvas) Draw(cycle, x int) {
	col := cycle % canvasWidth
	if uint(x-(col-spriteWidthHalf)) < spriteWidth {
		row := cycle / canvasWidth
		canvas[row][col] = '#'
	}
}

func (canvas Canvas) Print() {
	for _, row := range canvas {
		fmt.Println(string(row))
	}
}

func render(canvas Canvas, schedule []Instr) int {
	breakpoint := map[int]bool{
		20:  true,
		60:  true,
		100: true,
		140: true,
		180: true,
		220: true,
	}
	sum := 0
	x := 1
	for cycle := 1; len(schedule) > 0; cycle++ {
		canvas.Draw(cycle-1, x)
		if breakpoint[cycle] {
			sum += cycle * x
		}
		curr := &schedule[0]
		curr.latency--
		if curr.latency == 0 {
			schedule = schedule[1:]
			switch curr.op {
			case opAddx:
				x += curr.arg
			}
		}
	}
	return sum
}

func main() {
	data, _ := os.ReadFile("input")
	lines := strings.Split(string(data[:len(data)-1]), "\n")
	schedule := getSchedule(lines)
	canvas := getCanvas()
	fmt.Println("Part 1:", render(canvas, schedule)) // Expected: 15680
	fmt.Println("Part 2:")                           // Expected: ZFBFHGUP
	canvas.Print()
}
