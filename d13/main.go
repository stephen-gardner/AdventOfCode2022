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
	for i, end := 0, 1; i < len(raw); end = i + 1 {
		var ele interface{}
		if raw[i] == '[' {
			for open := 1; open > 0; end++ {
				switch raw[end] {
				case '[':
					open++
				case ']':
					open--
				}
			}
			ele = getPacket(raw[i:end])
		} else {
			for end < len(raw) && raw[end] != ',' {
				end++
			}
			ele, _ = strconv.Atoi(raw[i:end])
		}
		packet = append(packet, ele)
		i = end + 1
	}
	return packet
}

func getAllPackets(lines []string) [][]interface{} {
	packets := [][]interface{}{}
	for _, raw := range lines {
		if raw != "" {
			packets = append(packets, getPacket(raw))
		}
	}
	return packets
}

func isOrdered(left, right []interface{}) int {
	i := 0
	for {
		if i == len(left) {
			return len(right) - i
		}
		if i == len(right) {
			return -1
		}
		lValue, intLeft := left[i].(int)
		rValue, intRight := right[i].(int)
		if intLeft && intRight {
			if res := rValue - lValue; res != 0 {
				return res
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
	for i := 1; i < len(packets); i += 2 {
		if isOrdered(packets[i-1], packets[i]) >= 0 {
			sum += 1 + (i / 2)
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
			key *= 1 + i
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
