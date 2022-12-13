package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	dividerRaw1 = "[[2]]"
	dividerRaw2 = "[[6]]"
)

func getPacket(raw string) []interface{} {
	raw = raw[1 : len(raw)-1]
	packet := []interface{}{}
	i := 0
	for i < len(raw) {
		if raw[i] == '[' {
			end := i + 1
			open := 1
			for open > 0 {
				switch raw[end] {
				case '[':
					open++
				case ']':
					open--
				}
				end++
			}
			subList := getPacket(raw[i:end])
			packet = append(packet, subList)
			i = end + 1
		} else {
			end := i + 1
			for end < len(raw) && raw[end] != ',' {
				end++
			}
			n, _ := strconv.Atoi(raw[i:end])
			packet = append(packet, n)
			i = end + 1
		}
	}
	return packet
}

func getAllPackets(lines []string) [][]interface{} {
	packets := [][]interface{}{}
	for i := 0; i < len(lines); i += 3 {
		packets = append(packets, getPacket(lines[i]))
		packets = append(packets, getPacket(lines[i+1]))
	}
	return packets
}

func isOrdered(left, right []interface{}) int {
	i := 0
	for {
		if i == len(left) {
			if i == len(right) {
				return 0
			}
			return 1
		}
		if i == len(right) {
			return -1
		}
		lValue, intLeft := left[i].(int)
		rValue, intRight := right[i].(int)
		if intLeft && intRight {
			if lValue < rValue {
				return 1
			}
			if lValue > rValue {
				return -1
			}
		} else {
			var subLeft, subRight []interface{}
			if intLeft {
				subLeft = left[i : i+1]
			} else {
				subLeft = left[i].([]interface{})
			}
			if intRight {
				subRight = right[i : i+1]
			} else {
				subRight = right[i].([]interface{})
			}
			if res := isOrdered(subLeft, subRight); res != 0 {
				return res
			}
		}
		i++
	}
}

func part1(lines []string) int {
	sum := 0
	packets := getAllPackets(lines)
	for i := 0; i+1 < len(packets); i += 2 {
		left := packets[i]
		right := packets[i+1]
		pairIndex := ((i + 1) / 2) + 1
		if isOrdered(left, right) >= 0 {
			sum += pairIndex
		}
	}
	return sum
}

func part2(lines []string) int {
	packets := getAllPackets(lines)
	packets = append(packets, getPacket(dividerRaw1))
	packets = append(packets, getPacket(dividerRaw2))
	sort.Slice(packets, func(i, j int) bool {
		return isOrdered(packets[i], packets[j]) > 0
	})
	key := 1
	for i, p := range packets {
		pStr := fmt.Sprint(p)
		if pStr == dividerRaw1 || pStr == dividerRaw2 {
			key *= i + 1
		}
	}
	return key
}

func main() {
	data, _ := os.ReadFile("input")
	lines := strings.Split(string(data[:len(data)-1]), "\n")
	fmt.Println("Part 1:", part1(lines)) // Expected: 5557
	fmt.Println("Part 2:", part2(lines)) // Expected: 22425
}
