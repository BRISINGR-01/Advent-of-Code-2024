package main

import (
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Wire struct {
	name string
	val  int
}

type Gate string

const (
	AND Gate = "AND"
	OR  Gate = "OR"
	XOR Gate = "XOR"
)

type Connection struct {
	in1  *Wire
	in2  *Wire
	gate Gate
	out  *Wire
}

func Day24Pt1() int {
	wires, connections := getWires()

	processConnections(connections)

	return getNum(wires, 'z')
}

func getWires() ([]*Wire, []Connection) {
	wires := []*Wire{}
	connections := []Connection{}

	isReadingWires := true
	for _, line := range ReadInputLines() {
		if line == "" {
			isReadingWires = false
			continue
		}

		if isReadingWires {
			matches := regexp.MustCompile(`(...): (1|0)`).FindStringSubmatch(line)

			val, _ := strconv.Atoi(matches[2])
			wires = append(wires, &Wire{name: matches[1], val: val})
		} else {
			matches := regexp.MustCompile(`(...) (\w{2,3}) (...) -> (...)`).FindStringSubmatch(line)
			connections = append(connections, Connection{in1: getWire(matches[1], &wires), in2: getWire(matches[3], &wires), out: getWire(matches[4], &wires), gate: Gate(matches[2])})
		}
	}

	return wires, connections
}

func getWire(name string, wires *[]*Wire) *Wire {
	i := slices.IndexFunc(*wires, func(wire *Wire) bool {
		return wire.name == name
	})

	if i == -1 {
		wire := &Wire{name: name, val: -1}
		*wires = append(*wires, wire)
		return wire
	}

	return (*wires)[i]
}

func processConnections(connections []Connection) {
	for i := 0; i < len(connections); i++ {
		connection := connections[i]

		if connection.in1.val == -1 || connection.in2.val == -1 {
			connections = append(connections, connection)
			continue
		}

		switch connection.gate {
		case AND:
			connection.out.val = connection.in1.val & connection.in2.val
		case OR:
			connection.out.val = connection.in1.val | connection.in2.val
		case XOR:
			connection.out.val = connection.in1.val ^ connection.in2.val
		}
	}
}

func getNum(wires []*Wire, wType rune) int {
	subWires := extract(wires, wType)

	slices.SortFunc(subWires, func(a *Wire, b *Wire) int {
		return strings.Compare(b.name, a.name)
	})

	binaryRes := ""
	for _, wire := range subWires {
		binaryRes += strconv.Itoa(wire.val)
	}

	res, _ := strconv.ParseInt(binaryRes, 2, 64)
	return int(res)
}

func extract(wires []*Wire, wType rune) []*Wire {
	subWires := []*Wire{}

	for _, wire := range wires {
		if wire.name[0] == byte(wType) {
			subWires = append(subWires, wire)
		}
	}

	return subWires
}

func Day24Pt2() string {
	wires, connections := getWires()

	for _, swaps := range getOptions(len(connections)) {
		swappedConnections := slices.Clone(connections)
		swappedWires := []string{}
		for _, combo := range swaps {
			swap(&swappedConnections[combo[0]], &swappedConnections[combo[1]])

			swappedWires = append(swappedWires, swappedConnections[combo[0]].out.name)
			swappedWires = append(swappedWires, swappedConnections[combo[1]].out.name)
		}

		processConnections(swappedConnections)
		if checkWires(wires) {
			slices.Sort(swappedWires)
			return strings.Join(swappedWires, ",")
		}
	}

	panic("Not found")
}

func getOptions(n int) [][2][2]int {
	if n == 3 {
		return [][2][2]int{
			{[2]int{0, 1}, [2]int{0, 2}},
			{[2]int{0, 1}, [2]int{1, 2}},
			{[2]int{0, 2}, [2]int{1, 2}},

			{[2]int{0, 1}, [2]int{0, 2}},
			{[2]int{0, 1}, [2]int{0, 3}},
			{[2]int{0, 1}, [2]int{1, 3}},
			{[2]int{0, 1}, [2]int{1, 2}},
			{[2]int{0, 1}, [2]int{2, 3}},

			{[2]int{0, 2}, [2]int{0, 3}},
			{[2]int{0, 2}, [2]int{1, 3}},
			{[2]int{0, 2}, [2]int{2, 3}},
		}
	}
	options := getOptions(n - 1)

	prevOptions := len(options)
	// val := n - 1
	for i := 0; i < prevOptions; i++ {

	}

	return options
}

// func calcPossibilities(n int, count int) [][2][2]int {
// 	// possibilities := [][2][2]int{}
// 	for i := 0; i < n-count; i++ {

// 		calcPossibilities(i+1, count-1)
// 	}
// }

func checkWires(wires []*Wire) bool {
	x := getNum(wires, 'x')
	y := getNum(wires, 'y')
	z := getNum(wires, 'z')

	return x&y == z
}

func swap(c1 *Connection, c2 *Connection) {
	temp := c1.out
	c1.out = c2.out
	c2.out = temp
}
