package main

import (
	"fmt"
	"os"
	"strings"
)

func getInt(line string) int {
	n := 0
	for _, c := range line {
		switch c {
		case '=':
			n = (n * 5) - 2
		case '-':
			n = (n * 5) - 1
		default:
			n = (n * 5) + int(c-'0')
		}
	}
	return n
}

func getString(n int) string {
	buff := []byte{}
	for n > 0 {
		d := byte(n % 5)
		n /= 5
		switch d {
		case 0, 1, 2:
			buff = append(buff, d+'0')
		case 3:
			buff = append(buff, '=')
			n++
		case 4:
			buff = append(buff, '-')
			n++
		}
	}
	for i := 0; i < len(buff)/2; i++ {
		buff[i], buff[len(buff)-i-1] = buff[len(buff)-i-1], buff[i]
	}
	return string(buff)
}

func part1(lines []string) string {
	sum := 0
	for _, line := range lines {
		sum += getInt(line)
	}
	return getString(sum)
}

func main() {
	data, _ := os.ReadFile("input")
	lines := strings.Split(string(data[:len(data)-1]), "\n")
	fmt.Println("Part 1:", part1(lines)) // Expected: "2=001=-2=--0212-22-2"
}
