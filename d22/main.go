package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	DirRight = iota
	DirDown
	DirLeft
	DirUp
	NDirs
)

const (
	Move = 0xFF
	Keep = 0xFF
)

type (
	Cell struct {
		up     *Cell
		down   *Cell
		left   *Cell
		right  *Cell
		y      int
		x      int
		face   int
		symbol byte
	}
	Op struct {
		code int
		arg  int
	}
	Board struct {
		grid  []string
		cells map[[2]int]*Cell
		ops   []Op
		pos   *Cell
		dir   int
	}
)

func getBoard(data []byte) *Board {
	b := &Board{dir: DirRight}
	parts := strings.Split(string(data), "\n\n")
	b.setupGrid(parts[0])
	b.setupCells()
	b.linkCells()
	b.setupsOps(parts[1])
	x := 0
	for b.grid[0][x] != '.' {
		x++
	}
	b.pos = b.cells[[2]int{0, x}]
	return b
}

func (b *Board) execute(dirChange [][]int) {
	for _, op := range b.ops {
		switch op.code {
		case Move:
			for i := 0; i < op.arg; i++ {
				prevFace := b.pos.face
				nextCell := []*Cell{b.pos.right, b.pos.down, b.pos.left, b.pos.up}[b.dir]
				if nextCell.symbol == '.' {
					b.pos = nextCell
				}
				currFace := b.pos.face
				if currFace != prevFace {
					nextDir := dirChange[prevFace-1][currFace-1]
					if nextDir != Keep {
						b.dir = nextDir
					}
				}
			}
		case DirLeft:
			b.dir = ((b.dir - 1) + NDirs) % NDirs
		case DirRight:
			b.dir = (b.dir + 1) % NDirs
		}
	}
}

func (b *Board) getEdges(startY, endY, startX, endX int) [][]*Cell {
	edges := make([][]*Cell, 4)
	for x := startX; x <= endX; x++ {
		edges[DirUp] = append(edges[DirUp], b.cells[[2]int{startY, x}])   // top
		edges[DirDown] = append(edges[DirDown], b.cells[[2]int{endY, x}]) // bottom
	}
	for y := startY; y <= endY; y++ {
		edges[DirLeft] = append(edges[DirLeft], b.cells[[2]int{y, startX}]) // left
		edges[DirRight] = append(edges[DirRight], b.cells[[2]int{y, endX}]) // right
	}
	return edges
}

func (b *Board) getPassword() int {
	return (1000 * (b.pos.y + 1)) + (4 * (b.pos.x + 1)) + b.dir
}

func (b *Board) linkCells() {
	link := func(dir **Cell, y, x int) {
		if cell, exists := b.cells[[2]int{y, x}]; exists {
			*dir = cell
		}
	}
	for _, cell := range b.cells {
		link(&cell.up, cell.y-1, cell.x)
		link(&cell.down, cell.y+1, cell.x)
		link(&cell.left, cell.y, cell.x-1)
		link(&cell.right, cell.y, cell.x+1)
	}
}

func (b *Board) print() {
	for y := range b.grid {
		for x := range b.grid[y] {
			pos := [2]int{y, x}
			if cell, exists := b.cells[pos]; exists {
				fmt.Printf("%c", '0'+cell.face)
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

func (b *Board) setupGrid(raw string) {
	grid := strings.Split(raw, "\n")
	width := 0
	for _, row := range grid {
		if len(row) > width {
			width = len(row)
		}
	}
	for y := range grid {
		if len(grid[y]) < width {
			grid[y] = grid[y] + strings.Repeat(" ", width-len(grid[y]))
		}
	}
	b.grid = grid
}

func (b *Board) setupCells() {
	cells := map[[2]int]*Cell{}
	for y := range b.grid {
		for x := range b.grid[y] {
			c := b.grid[y][x]
			if c == ' ' {
				continue
			}
			cells[[2]int{y, x}] = &Cell{
				y:      y,
				x:      x,
				face:   1,
				symbol: c,
			}
		}
	}
	b.cells = cells
}

func (b *Board) setupFace(startY, endY, startX, endX, face int) {
	for y := startY; y <= endY; y++ {
		for x := startX; x <= endX; x++ {
			b.cells[[2]int{y, x}].face = face
		}
	}
}

func (b *Board) setupsOps(raw string) {
	i := 0
	for i < len(raw) {
		switch raw[i] {
		case 'L':
			b.ops = append(b.ops, Op{code: DirLeft})
			i++
		case 'R':
			b.ops = append(b.ops, Op{code: DirRight})
			i++
		default:
			n := 0
			for i < len(raw) && uint(raw[i]-'0') < 10 {
				n = (n * 10) + int(raw[i]-'0')
				i++
			}
			b.ops = append(b.ops, Op{code: Move, arg: n})
		}
	}
}

func part1(data []byte) int {
	b := getBoard(data)
	for y := range b.grid {
		x := 0
		for b.grid[y][x] == ' ' {
			x++
		}
		endX := x
		for endX+1 < len(b.grid[y]) && b.grid[y][endX+1] != ' ' {
			endX++
		}
		lft := b.cells[[2]int{y, x}]
		rht := b.cells[[2]int{y, endX}]
		lft.left = rht
		rht.right = lft
	}
	for x := range b.grid[0] {
		y := 0
		for b.grid[y][x] == ' ' {
			y++
		}
		endY := y
		for endY+1 < len(b.grid) && b.grid[endY+1][x] != ' ' {
			endY++
		}
		top := b.cells[[2]int{y, x}]
		bot := b.cells[[2]int{endY, x}]
		top.up = bot
		bot.down = top
	}
	b.execute(nil)
	return b.getPassword()
}

// Hard-coded cube regions, edge joins, and direction changes... disgusting 😱
func part2(data []byte) int {
	b := getBoard(data)
	edges := map[int][][]*Cell{}
	faces := [][]int{
		{0, 49, 50, 99},
		{0, 49, 100, 149},
		{50, 99, 50, 99},
		{100, 149, 0, 49},
		{100, 149, 50, 99},
		{150, 199, 0, 49},
	}
	for i, face := range faces {
		startY, endY, startX, endX := face[0], face[1], face[2], face[3]
		b.setupFace(startY, endY, startX, endX, i+1)
		edges[i+1] = b.getEdges(startY, endY, startX, endX)
	}
	reverse := func(edge []*Cell) []*Cell {
		half := len(edge) / 2
		for i := 0; i < half; i++ {
			edge[i], edge[len(edge)-i-1] = edge[len(edge)-i-1], edge[i]
		}
		return edge
	}
	joinEdges := func(e1 []*Cell, e1Dir int, e2 []*Cell, e2Dir int) {
		for i := range e1 {
			c1, c2 := e1[i], e2[i]
			*[]**Cell{&c1.right, &c1.down, &c1.left, &c1.up}[e1Dir] = c2
			*[]**Cell{&c2.right, &c2.down, &c2.left, &c2.up}[e2Dir] = c1
		}
	}
	joinEdges(edges[1][DirUp], DirUp, edges[6][DirLeft], DirLeft)
	joinEdges(edges[1][DirLeft], DirLeft, reverse(edges[4][DirLeft]), DirLeft)
	joinEdges(edges[2][DirUp], DirUp, edges[6][DirDown], DirDown)
	joinEdges(edges[2][DirDown], DirDown, edges[3][DirRight], DirRight)
	joinEdges(edges[2][DirRight], DirRight, reverse(edges[5][DirRight]), DirRight)
	joinEdges(edges[3][DirLeft], DirLeft, edges[4][DirUp], DirUp)
	joinEdges(edges[5][DirDown], DirDown, edges[6][DirRight], DirRight)
	dirChange := [][]int{
		{-1, Keep, Keep, DirRight, -1, DirRight},
		{Keep, -1, DirLeft, -1, DirLeft, DirUp},
		{Keep, DirUp, -1, DirDown, Keep, -1},
		{DirRight, -1, DirRight, -1, Keep, Keep},
		{-1, DirLeft, Keep, Keep, -1, DirLeft},
		{DirDown, DirDown, -1, Keep, DirUp, -1},
	}
	b.execute(dirChange)
	return b.getPassword()
}

func main() {
	data, _ := os.ReadFile("input")
	data = data[:len(data)-1]
	fmt.Println("Part 1:", part1(data)) // Expected: 64256
	fmt.Println("Part 2:", part2(data)) // Expected: 109224
}
