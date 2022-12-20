package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const decryptionKey = 811589153

type Node struct {
	id    int
	value int
	prev  *Node
	next  *Node
}

func getNums(lines []string) *Node {
	var head *Node
	curr := &head
	prev := head
	for i := range lines {
		num := &Node{
			prev: prev,
			id:   i,
		}
		num.value, _ = strconv.Atoi(lines[i])
		*curr = num
		prev = *curr
		head.prev = *curr
		curr = &(*curr).next
	}
	*curr = head
	return head
}

func mix(head *Node, size, numRounds int) *Node {
	getNode := func(curr *Node, id int) *Node {
		for curr.id != id {
			curr = curr.next
		}
		return curr
	}
	for round := 0; round < numRounds; round++ {
		for i := 0; i < size; i++ {
			curr := getNode(head, i)
			n := curr.value % (size - 1)
			for n != 0 {
				if n > 0 {
					curr.id, curr.next.id = curr.next.id, curr.id
					curr.value, curr.next.value = curr.next.value, curr.value
					curr = curr.next
					n--
				} else {
					curr.id, curr.prev.id = curr.prev.id, curr.id
					curr.value, curr.prev.value = curr.prev.value, curr.value
					curr = curr.prev
					n++
				}
			}
		}
	}
	return head
}

func getGroveCoordinates(head *Node, size int) int {
	getNumAt := func(curr *Node, size, i int) int {
		for i %= size; i > 0; i-- {
			curr = curr.next
		}
		return curr.value
	}
	for head.value != 0 {
		head = head.next
	}
	gc0 := getNumAt(head, size, 1000)
	gc1 := getNumAt(head, size, 2000)
	gc2 := getNumAt(head, size, 3000)
	return gc0 + gc1 + gc2
}

func listPrint(head *Node, size int) {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = head.value
		head = head.next
	}
	fmt.Println(arr)
}

func part1(lines []string) int {
	head := getNums(lines)
	size := len(lines)
	head = mix(head, size, 1)
	return getGroveCoordinates(head, size)
}

func part2(lines []string) int {
	head := getNums(lines)
	size := len(lines)
	for i := 0; i < size; i++ {
		head.value *= decryptionKey
		head = head.next
	}
	head = mix(head, size, 10)
	return getGroveCoordinates(head, size)
}

func main() {
	data, _ := os.ReadFile("input")
	lines := strings.Split(string(data[:len(data)-1]), "\n")
	fmt.Println("Part 1:", part1(lines)) // Expected: 5498
	fmt.Println("Part 2:", part2(lines)) // Expected: 3390007892081
}
