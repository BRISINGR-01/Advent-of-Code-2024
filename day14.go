package main

import (
	"regexp"
	"slices"
	"strconv"
)

type Robot struct {
	x  int
	y  int
	vx int
	vy int
}

func Day14Pt1() int {
	COLS := 101
	ROWS := 103

	robots := GetRobots()

	for i := 0; i < 100; i++ {
		for i := range robots {
			MoveRobot(&robots[i], COLS, ROWS)
		}
	}

	quadrants := []int{0, 0, 0, 0}

	middleCol := (COLS - 1) / 2
	middleRow := (ROWS - 1) / 2

	for _, robot := range robots {
		if robot.x < middleCol {
			if robot.y < middleRow {
				quadrants[0]++
			} else if robot.y > middleRow {
				quadrants[1]++
			}
		} else if robot.x > middleCol {
			if robot.y < middleRow {
				quadrants[2]++
			} else if robot.y > middleRow {
				quadrants[3]++
			}
		}
	}

	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func MoveRobot(r *Robot, cols int, rows int) {
	r.x += r.vx
	if r.x < 0 {
		r.x += cols
	} else if r.x >= cols {
		r.x -= cols
	}

	r.y += r.vy
	if r.y < 0 {
		r.y += rows
	} else if r.y >= rows {
		r.y -= rows
	}
}

func GetRobots() []Robot {
	robots := []Robot{}

	for _, line := range ReadInputLines() {
		match := regexp.MustCompile(`p=(\d+),(\d+) v\=(\-?\d+),(\-?\d+)`).FindStringSubmatch(line)

		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		vx, _ := strconv.Atoi(match[3])
		vy, _ := strconv.Atoi(match[4])

		robots = append(robots, Robot{
			x:  x,
			y:  y,
			vx: vx,
			vy: vy,
		})
	}

	return robots
}

func PrintField(robots []Robot, cols int, rows int) {
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			count := 0
			for _, robot := range robots {
				if robot.x == x && robot.y == y {
					count++
				}
			}

			if count > 0 {
				// print(strconv.Itoa(count))
				print("#")
			} else {
				print(".")
			}
		}
		println()
	}
}

func Day14Pt2() int {

	COLS := 101
	ROWS := 103

	robots := GetRobots()

	step := 0
	for step < 40000 {
		step++
		rows := map[int][]int{}
		for i := range robots {
			MoveRobot(&robots[i], COLS, ROWS)
		}

		for i := range robots {
			rows[robots[i].y] = append(rows[robots[i].y], robots[i].x)
		}

		for _, row := range rows {
			slices.Sort(row)
			consecutive := 0

			for row_i, y := range row {
				if row_i == 0 || y == row[row_i-1] {
					continue
				}

				if y == row[row_i-1]+1 {
					consecutive++
					if consecutive > 8 {
						PrintField(robots, COLS, ROWS)
						return step
					}
				} else {
					consecutive = 0
				}
			}
		}
	}

	panic("not found")
}
