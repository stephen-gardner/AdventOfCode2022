package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type (
	Sensor struct {
		y  int
		x  int
		by int
		bx int
		bd int
	}
	Sensors []*Sensor
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func getSensors(lines []string) Sensors {
	getVal := func(s string) int {
		if uint(s[len(s)-1])-'0' > 9 {
			s = s[:len(s)-1]
		}
		val, _ := strconv.Atoi(strings.Split(s, "=")[1])
		return val
	}
	sensors := make([]*Sensor, len(lines))
	for i := range sensors {
		data := strings.Split(lines[i], " ")
		s := &Sensor{
			x:  getVal(data[2]),
			y:  getVal(data[3]),
			bx: getVal(data[8]),
			by: getVal(data[9]),
		}
		s.bd = abs(s.y-s.by) + abs(s.x-s.bx)
		sensors[i] = s
	}
	return sensors
}

func (sensors Sensors) getBounds() (minY, maxY, minX, maxX int) {
	minY, maxY = sensors[0].y, sensors[0].y
	minX, maxX = sensors[0].x, sensors[0].x
	for _, s := range sensors {
		minY = min(minY, min(s.y, s.y-s.bd))
		maxY = max(maxY, max(s.y, s.y+s.bd))
		minX = min(minX, min(s.x, s.x-s.bd))
		maxX = max(maxX, max(s.x, s.x+s.bd))
	}
	return
}

func (sensors Sensors) inRange(y, x int) bool {
	for _, s := range sensors {
		if abs(s.y-y)+abs(s.x-x) <= s.bd {
			return true
		}
	}
	return false
}

// Points that cannot possibly have a beacon are all within range of a beacon,
//	except for the point with the beacon itself, as all sensors point to its
//	closest beacon
func part1(lines []string, targetY int) int {
	sensors := getSensors(lines)
	// Some sensors may point to the same beacon
	beaconsInRow := map[int]bool{}
	for _, s := range sensors {
		if s.by == targetY {
			beaconsInRow[s.bx] = true
		}
	}
	count := 0
	_, _, minX, maxX := sensors.getBounds()
	for x := minX; x <= maxX; x++ {
		if sensors.inRange(targetY, x) {
			count++
		}
	}
	return count - len(beaconsInRow)
}

// Brute forcing every possible location takes an eternity, but walking the
//	perimeter of each sensor (1 unit outside of its detected beacon) is fast
func part2(lines []string) int {
	sensors := getSensors(lines)
	// Checks every point from src to dst
	searchLine := func(srcY, srcX, dstY, dstX, dy, dx int) (int, int, bool) {
		// Find a point outside of all sensors' range that is also surrounded
		//	by points that are within the range of one
		for y, x := srcY, srcX; y <= dstY && x <= dstX; y, x = y+dy, x+dx {
			if !sensors.inRange(y, x) &&
				sensors.inRange(y+1, x) &&
				sensors.inRange(y-1, x) &&
				sensors.inRange(y, x+1) &&
				sensors.inRange(y, x-1) {
				return y, x, true
			}
		}
		return 0, 0, false
	}
	// Each point contains delta for moving to the point that follows it
	points := [4][4]int{
		{0, 0, -1, 1},
		{0, 0, 1, 1},
		{0, 0, 1, -1},
		{0, 0, -1, -1},
	}
	for _, s1 := range sensors {
		points[0][0], points[0][1] = s1.y, s1.x-s1.bd-1 // Left
		points[1][0], points[1][1] = s1.y, s1.x+s1.bd+1 // Right
		points[2][0], points[2][1] = s1.y-s1.bd-1, s1.x // Up
		points[3][0], points[3][1] = s1.y+s1.bd+1, s1.x // Down
		for i := range points {
			src := points[i]
			dst := points[(i+1)%4]
			if y, x, found := searchLine(src[0], src[1], dst[0], dst[1], src[2], src[3]); found {
				return (x * 4000000) + y
			}
		}
	}
	return -1
}

func main() {
	data, _ := os.ReadFile("input")
	lines := strings.Split(string(data[:len(data)-1]), "\n")
	fmt.Println("Part 1:", part1(lines, 2000000)) // Expected: 4861076
	fmt.Println("Part 2:", part2(lines))          // Expected: 10649103160102
}
