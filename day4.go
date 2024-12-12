package main

import "strings"

func Day4Pt1() int {
	result := 0
	grid := ReadInputRunes()
	side_length := len(grid)

	for _, row := range grid {
		result += Count(row)
	}

	for column_i := range grid[0] {
		column := []rune{}

		for row_i := range grid {
			column = append(column, grid[row_i][column_i])
		}

		result += Count(column)
	}

	for row_i, row := range grid {
		diagonal_row := []rune{}
		diagonal_column := []rune{}
		rev_diagonal_row := []rune{}
		rev_diagonal_column := []rune{}

		for column_i := range row {
			diagonal_i := row_i + column_i
			if diagonal_i < side_length {
				diagonal_row = append(diagonal_row, grid[diagonal_i][column_i])
				diagonal_column = append(diagonal_column, grid[column_i][diagonal_i])
			}

			rev_diagonal_i := side_length - row_i + side_length - column_i - 2
			if rev_diagonal_i < side_length {
				rev_diagonal_column = append(rev_diagonal_column, grid[column_i][rev_diagonal_i])

				rev_diagonal_row = append(rev_diagonal_row, grid[side_length-column_i-1][side_length-rev_diagonal_i-1])
			}
		}

		if row_i != 0 {
			result += Count(diagonal_row)
		}
		if row_i != side_length-1 {
			result += Count(rev_diagonal_row)
		}
		result += Count(diagonal_column)
		result += Count(rev_diagonal_column)

	}

	return result
}

func Count(characters []rune) int {
	return strings.Count(string(characters), "XMAS") + strings.Count(string(characters), "SAMX")
}

func Day4Pt2() int {
	result := 0
	grid := ReadInputRunes()

	for y, row := range grid {
		for x, char := range row {
			if char == 'A' {
				if Check(grid, x, y) {
					result++
				}
			}
		}
	}

	return result
}

func Check(grid [][]rune, x int, y int) bool {
	if x == 0 || x == len(grid)-1 || y == 0 || y == len(grid)-1 {
		return false
	}

	if grid[y-1][x-1] == 'M' && grid[y-1][x+1] == 'M' && grid[y+1][x-1] == 'S' && grid[y+1][x+1] == 'S' {
		return true
	}
	if grid[y-1][x-1] == 'S' && grid[y-1][x+1] == 'S' && grid[y+1][x-1] == 'M' && grid[y+1][x+1] == 'M' {
		return true
	}
	if grid[y-1][x-1] == 'S' && grid[y-1][x+1] == 'M' && grid[y+1][x-1] == 'S' && grid[y+1][x+1] == 'M' {
		return true
	}
	if grid[y-1][x-1] == 'M' && grid[y-1][x+1] == 'S' && grid[y+1][x-1] == 'M' && grid[y+1][x+1] == 'S' {
		return true
	}

	return false
}
