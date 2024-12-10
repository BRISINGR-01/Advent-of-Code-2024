package main

import "slices"

type MPos struct {
	x     int
	y     int
	value int
}

func PosFrom(x int, y int, grid [][]int) MPos {
	return MPos{
		x:     x,
		y:     y,
		value: grid[y][x],
	}
}

func Day10Pt1() int {
	result := 0

	grid := ReadInputGridNumbers()

	for _, trailHead := range FindTrailHeads(grid) {
		result += len(FilterUnique(Get9s(trailHead, grid)))
	}

	return result
}

func FindTrailHeads(grid [][]int) []MPos {
	positions := []MPos{}

	for row_i, row := range grid {
		for column_i, value := range row {
			if value == 0 {
				positions = append(positions, PosFrom(column_i, row_i, grid))
			}
		}
	}

	return positions
}

func Get9s(pos MPos, grid [][]int) []MPos {
	if pos.value == 9 {
		return []MPos{pos}
	}

	result := []MPos{}

	if pos.y > 0 && grid[pos.y-1][pos.x] == pos.value+1 {
		result = append(result, Get9s(PosFrom(pos.x, pos.y-1, grid), grid)...)
	}

	if pos.y < len(grid)-1 && grid[pos.y+1][pos.x] == pos.value+1 {
		result = append(result, Get9s(PosFrom(pos.x, pos.y+1, grid), grid)...)
	}

	if pos.x > 0 && grid[pos.y][pos.x-1] == pos.value+1 {
		result = append(result, Get9s(PosFrom(pos.x-1, pos.y, grid), grid)...)
	}

	if pos.x < len(grid[pos.y])-1 && grid[pos.y][pos.x+1] == pos.value+1 {
		result = append(result, Get9s(PosFrom(pos.x+1, pos.y, grid), grid)...)
	}

	return result
}

func FilterUnique(positions []MPos) []MPos {
	result := []MPos{}

	for _, pos := range positions {
		if !slices.ContainsFunc(result, func(other MPos) bool {
			return pos.x == other.x && pos.y == other.y
		}) {
			result = append(result, pos)
		}
	}

	return result
}

func Day10Pt2() int {
	result := 0

	grid := ReadInputGridNumbers()

	for _, trailHead := range FindTrailHeads(grid) {
		result += len(Get9s(trailHead, grid))
	}

	return result
}
