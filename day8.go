package main

import (
	"math"
)

type AntennaPos struct {
	row    int
	column int
}

func Day8Pt1() int {
	result := 0
	grid := ReadInputRunes()

	antennae := map[rune][]AntennaPos{}

	for row_i, row := range grid {
		for column_i, char := range row {
			if char == '.' {
				continue
			}

			antennae[char] = append(antennae[char], AntennaPos{
				row:    row_i,
				column: column_i,
			})
		}
	}

	for _, antenna := range antennae {
		for _, pair := range GetPairs(antenna) {
			antiNode1, antiNode2 := CalcAntiNodes(pair)

			if IsValid(antiNode1, grid) {
				result++
				grid[antiNode1.row][antiNode1.column] = '#'
			}

			if IsValid(antiNode2, grid) {
				grid[antiNode2.row][antiNode2.column] = '#'
				result++
			}
		}
	}

	return result
}

type PairAntennaes struct {
	antenna1 AntennaPos
	antenna2 AntennaPos
}

func GetPairs(antennaes []AntennaPos) []PairAntennaes {
	if len(antennaes) == 2 {
		return []PairAntennaes{{antenna1: antennaes[0], antenna2: antennaes[1]}}
	}

	pairs := []PairAntennaes{}

	remaining := antennaes[1:]

	for _, antenna := range remaining {
		pairs = append(pairs, PairAntennaes{antenna1: antenna, antenna2: antennaes[0]})
	}

	return append(pairs, GetPairs(remaining)...)
}

func CalcAntiNodes(pair PairAntennaes) (AntennaPos, AntennaPos) {
	pos1 := AntennaPos{
		row:    pair.antenna1.row,
		column: pair.antenna1.column,
	}

	pos2 := AntennaPos{
		row:    pair.antenna2.row,
		column: pair.antenna2.column,
	}

	diffCol := math.Abs(float64(pair.antenna1.column - pair.antenna2.column))

	if pair.antenna1.column > pair.antenna2.column {
		pos1.column += int(diffCol)
		pos2.column -= int(diffCol)
	} else {
		pos1.column -= int(diffCol)
		pos2.column += int(diffCol)
	}

	diffRow := math.Abs(float64(pair.antenna1.row - pair.antenna2.row))

	if pair.antenna1.row > pair.antenna2.row {
		pos1.row += int(diffRow)
		pos2.row -= int(diffRow)
	} else {
		pos1.row -= int(diffRow)
		pos2.row += int(diffRow)
	}

	return pos1, pos2
}

func IsValid(pos AntennaPos, grid [][]rune) bool {
	return pos.row >= 0 &&
		pos.row < len(grid) &&
		pos.column >= 0 &&
		pos.column < len(grid[pos.row]) &&
		(grid[pos.row][pos.column] != '#')
}

func Day8Pt2() int {
	result := 0
	grid := ReadInputRunes()

	antennae := map[rune][]AntennaPos{}

	for row_i, row := range grid {
		for column_i, char := range row {
			if char == '.' {
				continue
			}

			antennae[char] = append(antennae[char], AntennaPos{
				row:    row_i,
				column: column_i,
			})
		}
	}

	for _, antenna := range antennae {
		result += len(antenna)

		for _, pair := range GetPairs(antenna) {
			for _, antiNode := range CalcAntiNodes2(pair, grid) {
				if grid[antiNode.row][antiNode.column] != '.' {
					continue
				}

				result++
				grid[antiNode.row][antiNode.column] = '#'
			}
		}
	}

	return result
}

func CalcAntiNodes2(pair PairAntennaes, grid [][]rune) []AntennaPos {
	posList := []AntennaPos{}
	upper := pair.antenna1
	lower := pair.antenna2

	if pair.antenna1.row > pair.antenna2.row {
		upper = pair.antenna2
		lower = pair.antenna1
	}

	nodes := GetLowerNodes(upper, lower, grid)
	nodes = append(nodes, GetMiddleNodes(upper, lower, grid)...)
	nodes = append(nodes, GetHigherNodes(upper, lower, grid)...)

	for _, v := range nodes {
		if !IsValid(v, grid) {
			continue
		}

		posList = append(posList, v)
	}

	return posList
}

func GetHigherNodes(upper AntennaPos, lower AntennaPos, grid [][]rune) []AntennaPos {
	posList := []AntennaPos{}
	rowDist := lower.row - upper.row
	columnDist := lower.column - upper.column
	lastColumn := upper.column

	for i := upper.row - rowDist; i >= 0; i -= rowDist {
		lastColumn -= columnDist
		posList = append(posList, AntennaPos{
			row:    i,
			column: lastColumn,
		})
	}

	return posList
}

func GetLowerNodes(upper AntennaPos, lower AntennaPos, grid [][]rune) []AntennaPos {
	posList := []AntennaPos{}
	rowDist := lower.row - upper.row
	columnDist := upper.column - lower.column
	lastColumn := upper.column

	for i := upper.row + rowDist; i < len(grid); i += rowDist {
		lastColumn -= columnDist
		posList = append(posList, AntennaPos{
			row:    i,
			column: lastColumn,
		})
	}

	return posList
}

func GetMiddleNodes(upper AntennaPos, lower AntennaPos, grid [][]rune) []AntennaPos {
	posList := []AntennaPos{}
	rowDist := lower.row - upper.row
	columnDist := upper.column - lower.column
	lastColumn := upper.column

	for i := upper.row - rowDist; i > lower.row; i -= rowDist {
		lastColumn -= columnDist
		posList = append(posList, AntennaPos{
			row:    i,
			column: lastColumn,
		})
	}

	return posList
}

// 1096 too low
// 1316 too high
