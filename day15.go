package main

func Day15Pt1() int {
	robot, grid, movements := ExtractData()

	for _, movement := range movements {
		nextPos := NextPos(robot, movement)

		if GetRune(nextPos, grid) == '#' {
			continue
		}

		if grid[nextPos.y][nextPos.x] != 'O' {
			grid[nextPos.y][nextPos.x] = '@'
			grid[robot.y][robot.x] = '.'
			robot = nextPos
			continue
		}

		nextOPos := NextPos(nextPos, movement)

		if grid[nextOPos.y][nextOPos.x] == '#' {
			continue
		}

		if !CanPushBox(robot, movement, grid) {
			continue
		}

		PushBox(robot, movement, grid)
		grid[nextPos.y][nextPos.x] = '@'
		grid[robot.y][robot.x] = '.'
		robot = nextPos
	}

	result := 0

	for y, row := range grid {
		for x, item := range row {
			if item == 'O' {
				result += 100*y + x
			}
		}
	}

	return result
}

func ExtractData() (Pos, [][]rune, []rune) {
	movements := []rune{}
	robot := Pos{}
	grid := [][]rune{}

	isMap := true
	for row_i, row := range ReadInputRunes() {
		if len(row) == 0 {
			isMap = false
			continue
		}

		if !isMap {
			movements = append(movements, row...)
			continue
		}

		grid = append(grid, []rune{})
		for col_i, item := range row {
			if item == '@' {
				robot = Pos{col_i, row_i}
			}

			grid[row_i] = append(grid[row_i], item)
		}
	}

	return robot, grid, movements
}

func NextPos(robot Pos, movement rune) Pos {
	switch movement {
	case '^':
		return Pos{robot.x, robot.y - 1}
	case 'v':
		return Pos{robot.x, robot.y + 1}
	case '>':
		return Pos{robot.x + 1, robot.y}
	case '<':
		return Pos{robot.x - 1, robot.y}
	default:
		panic("Invalid movement " + string(movement))
	}
}

func CanPushBox(robot Pos, movement rune, grid [][]rune) bool {
	nextPos := NextPos(robot, movement)
	for {
		if GetRune(nextPos, grid) == '.' {
			return true
		}
		if GetRune(nextPos, grid) == '#' {
			return false
		}
		nextPos = NextPos(nextPos, movement)
	}
}

func PushBox(robot Pos, movement rune, grid [][]rune) Pos {
	nextPos := NextPos(robot, movement)
	grid[nextPos.y][nextPos.x] = '.'

	for nextPos = NextPos(nextPos, movement); GetRune(nextPos, grid) == 'O'; nextPos = NextPos(nextPos, movement) {
	}

	grid[nextPos.y][nextPos.x] = 'O'

	return nextPos
}

func Day15Pt2() int {
	robot, oldGrid, movements := ExtractData()
	robot.x *= 2
	grid := make([][]rune, len(oldGrid))

	for y, row := range oldGrid {
		for _, item := range row {
			switch item {
			case '@':
				grid[y] = append(grid[y], item, '.')
			case 'O':
				grid[y] = append(grid[y], '[', ']')
			default:
				grid[y] = append(grid[y], item, item)
			}
		}
	}

	for _, movement := range movements {
		nextPos := NextPos(robot, movement)

		if GetRune(nextPos, grid) == '#' {
			continue
		}

		if GetRune(nextPos, grid) == '.' {
			grid[robot.y][robot.x] = '.'
			robot = nextPos
			grid[robot.y][robot.x] = '@'
			// println(string(movement))
			continue
		}

		if !CanPushBox2(nextPos, movement, grid) {
			continue
		}

		PushBox2(nextPos, movement, grid)
		grid[robot.y][robot.x] = '.'
		robot = nextPos
		grid[robot.y][robot.x] = '@'
		// println(string(movement))
	}

	result := 0

	for y, row := range grid {
		for x, item := range row {
			if item == '[' {
				result += 100*y + x
			}
		}
	}

	return result
}

func CanPushBox2(boxPart Pos, movement rune, grid [][]rune) bool {
	if GetRune(boxPart, grid) == '#' {
		return false
	}
	if GetRune(boxPart, grid) == '.' {
		if movement == '<' || movement == '>' {
			return true
		}

		return true
	}

	otherBox := Pos{boxPart.x + 1, boxPart.y}
	isRight := GetRune(boxPart, grid) == ']'
	if isRight {
		otherBox = Pos{boxPart.x - 1, boxPart.y}
	}

	switch movement {
	case 'v', '^':
		return CanPushBox2(NextPos(boxPart, movement), movement, grid) && CanPushBox2(NextPos(otherBox, movement), movement, grid)
	case '>':
		if isRight {
			return CanPushBox2(NextPos(boxPart, movement), movement, grid)
		} else {
			return CanPushBox2(NextPos(otherBox, movement), movement, grid)
		}
	case '<':
		if isRight {
			return CanPushBox2(NextPos(otherBox, movement), movement, grid)
		} else {
			return CanPushBox2(NextPos(boxPart, movement), movement, grid)
		}
	}

	panic("Invalid movement " + string(movement))
}

func PushBox2(boxL Pos, movement rune, grid [][]rune) {
	if GetRune(boxL, grid) == ']' {
		PushBox2(Pos{boxL.x - 1, boxL.y}, movement, grid)
		return
	}

	if GetRune(boxL, grid) != '[' {
		return
	}

	grid[boxL.y][boxL.x] = '.'
	grid[boxL.y][boxL.x+1] = '.'

	switch movement {
	case 'v', '^':
		PushBox2(NextPos(boxL, movement), movement, grid)
		PushBox2(NextPos(Pos{boxL.x + 1, boxL.y}, movement), movement, grid)
	case '>':
		PushBox2(NextPos(Pos{boxL.x + 1, boxL.y}, movement), movement, grid)
	case '<':
		PushBox2(NextPos(boxL, movement), movement, grid)
	default:
		panic("Invalid movement " + string(movement))
	}

	next := NextPos(boxL, movement)
	grid[next.y][next.x] = '['
	grid[next.y][next.x+1] = ']'
}

func GetRune(pos Pos, grid [][]rune) rune {
	return grid[pos.y][pos.x]
}
