package main

import (
	"regexp"
	"slices"
	"strings"
)

type Computer struct {
	name        string
	connections []*Computer
}

func Day23Pt1() int {
	computers := getComputers()
	connections := []string{}

	for _, computer := range computers {
		for i := len(computer.connections) - 1; i > 0; i-- {
			connection := computer.connections[i]

			for j := i - 1; j >= 0; j-- {
				otherConnection := computer.connections[j]
				if !areConnected(connection, otherConnection) {
					continue
				}

				connection := []string{computer.name, connection.name, otherConnection.name}
				slices.Sort(connection)

				connections = append(connections, "-"+connection[0]+"-"+connection[1]+"-"+connection[2])
			}
		}
	}

	slices.Sort(connections)
	connections = slices.Compact(connections)
	result := 0

	for _, truple := range connections {
		if strings.Contains(truple, "-t") {
			result++
		}
	}

	return result
}

func getComputers() []*Computer {
	computers := []*Computer{}

	for _, line := range ReadInputLines() {
		matches := regexp.MustCompile(`(\w\w)\-(\w\w)`).FindStringSubmatch(line)

		comp1 := find(matches[1], computers)
		if comp1 == nil {
			comp1 = &Computer{
				name: matches[1],
			}
			computers = append(computers, comp1)
		}

		comp2 := find(matches[2], computers)
		if comp2 == nil {
			comp2 = &Computer{
				name: matches[2],
			}
			computers = append(computers, comp2)
		}

		comp1.connections = append(comp1.connections, comp2)
		comp2.connections = append(comp2.connections, comp1)
	}

	return computers
}

func areConnected(a *Computer, b *Computer) bool {
	for _, connection := range a.connections {
		if b.name == connection.name {
			return true
		}
	}

	return false
}

func find(name string, computers []*Computer) *Computer {
	i := slices.IndexFunc(computers, func(comp *Computer) bool { return comp.name == name })

	if i == -1 {
		return nil
	}

	return computers[i]
}

func Day23Pt2() string {
	computers := getComputers()
	longest := []string{}

	for i, computer := range computers {
		for _, comp := range computer.connections {
			party := []*Computer{computer, comp}
			for j := i - 1; j >= 0; j-- {
				if canAdd(computers[j], party) {
					party = append(party, computers[j])
				}
			}

			if len(longest) < len(party) {
				longest = []string{}
				for _, c := range party {
					longest = append(longest, c.name)
				}
			}
		}
	}

	slices.Sort(longest)

	return strings.Join(longest, ",")
}

func canAdd(comp *Computer, party []*Computer) bool {
	for _, c := range party {
		if comp.name == c.name || !areConnected(comp, c) {
			return false
		}
	}
	return true
}
