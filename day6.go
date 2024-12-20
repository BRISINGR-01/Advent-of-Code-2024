package main

import (
	"slices"
)

type DPos struct {
	column    int
	row       int
	direction int
}
type MoveFunc = func(pos *DPos)

func (pos *DPos) MoveToNext(grid [][]rune) bool {
	moves := []MoveFunc{MoveUp, MoveRight, MoveDown, MoveLeft}
	moveFunc := moves[pos.direction]

	temp := DPos{
		column:    pos.column,
		row:       pos.row,
		direction: pos.direction,
	}

	moveFunc(&temp)

	if !IsInField(temp, grid) {
		return false
	}

	if grid[temp.row][temp.column] == '#' {
		pos.direction++

		if pos.direction >= len(moves) {
			pos.direction = 0
		}
		return true
	}

	pos.column = temp.column
	pos.row = temp.row

	return true
}

func Day6Pt1() int {
	result := 1
	grid := ReadInputRunes()
	pos := FindBeginning(grid)
	grid[pos.row][pos.column] = 'X'

	for pos.MoveToNext(grid) {
		if grid[pos.row][pos.column] == 'X' {
			continue
		}

		grid[pos.row][pos.column] = 'X'
		result++
	}

	return result
}

func FindBeginning(grid [][]rune) DPos {
	pos := DPos{-1, -1, 0}

	pos.row = slices.IndexFunc(grid, func(row []rune) bool {
		return slices.Contains(row, '^')
	})

	pos.column = slices.IndexFunc(grid[pos.row], func(char rune) bool {
		return char == '^'
	})

	return pos
}

func MoveUp(pos *DPos) {
	pos.row--
}

func MoveDown(pos *DPos) {
	pos.row++
}

func MoveLeft(pos *DPos) {
	pos.column--
}

func MoveRight(pos *DPos) {
	pos.column++
}

func IsInField(pos DPos, grid [][]rune) bool {
	return pos.row >= 0 && pos.row < len(grid) && pos.column >= 0 && pos.column < len(grid[pos.row])
}

func Day6Pt2() int {
	result := 0
	grid := ReadInputRunes()
	pos := FindBeginning(grid)
	grid[pos.row][pos.column] = 'X'
	prev := DPos{
		column:    pos.column,
		row:       pos.row,
		direction: pos.direction,
	}

	for pos.MoveToNext(grid) {
		if grid[pos.row][pos.column] == 'X' {
			continue
		}

		grid[pos.row][pos.column] = '#'

		if WillCircle(prev, grid) {
			result++
		}

		grid[pos.row][pos.column] = 'X'
		prev = pos
	}

	return result
}

func WillCircle(pos DPos, grid [][]rune) bool {
	history := []DPos{}

	prevDir := pos.direction
	for pos.MoveToNext(grid) {
		if pos.direction == prevDir {
			continue
		}

		if slices.ContainsFunc(history, func(origin DPos) bool {
			return origin.column == pos.column && origin.row == pos.row && origin.direction == pos.direction
		}) {
			return true
		}

		history = append(history, pos)
		prevDir = pos.direction
	}

	return false
}
