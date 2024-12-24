package main

import (
	"slices"
)

type Direction rune

const (
	Up    Direction = '^'
	Down  Direction = 'v'
	Left  Direction = '<'
	Right Direction = '>'
)

type Movement struct {
	direction Direction
	x         int
	y         int
}

func Day16Pt1() int {
	maze := ReadInputRunes()

	path := explore(getStartPos(maze), maze, []Movement{})

	println()
	for _, p := range path {
		print(string(p.direction))
	}
	println()
	return calcScore(path)
}

func getStartPos(maze [][]rune) Movement {
	for y, row := range maze {
		for x, char := range row {
			if char == 'S' {
				return Movement{Right, x, y}
			}
		}
	}

	panic("Could not find start position")
}

func explore(pos Movement, maze [][]rune, path []Movement) []Movement {
	path = append(path, pos)
	possibilities := [][]Movement{}

	for _, dir := range []Direction{Up, Down, Left, Right} {
		nextPos := calcNext(pos, dir)
		if isEnd(nextPos, maze) {
			return path
		}
		if !canPass(nextPos, maze) ||
			slices.ContainsFunc(path, func(p Movement) bool { return p.x == nextPos.x && p.y == nextPos.y }) {
			continue
		}

		p := explore(nextPos, maze, path)

		if len(p) == 0 {
			continue
		}

		possibilities = append(possibilities, p)
	}

	if len(possibilities) == 0 {
		return []Movement{}
	}

	min := possibilities[0]

	for _, pos := range possibilities[1:] {
		if calcScore(pos) < calcScore(min) {
			min = pos
		}
	}

	return min
}

func canPass(pos Movement, maze [][]rune) bool {
	return maze[pos.y][pos.x] == '.'
}

func isEnd(pos Movement, maze [][]rune) bool {
	return maze[pos.y][pos.x] == 'E'
}

func calcNext(pos Movement, direction Direction) Movement {
	switch direction {
	case Up:
		return Movement{direction, pos.x, pos.y - 1}
	case Down:
		return Movement{direction, pos.x, pos.y + 1}
	case Left:
		return Movement{direction, pos.x - 1, pos.y}
	case Right:
		return Movement{direction, pos.x + 1, pos.y}
	default:
		panic("Invalid direction")
	}
}

func calcScore(path []Movement) int {
	score := 0

	prev := Right
	for _, pos := range path {
		if pos.direction == prev {
			score++
			continue
		}

		prev = pos.direction
		score += 1000 + 1
	}

	return score
}
