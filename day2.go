package main

import (
	"math"
)

func Day2Pt1() int {
	result := 0

	for _, row := range ReadInputNumbers() {
		if IsSafe(row) {
			result++
		}
	}

	return result
}

func IsSafe(list []int) bool {
	shouldIncrease := list[0] < list[1]

	for i := 0; i < len(list)-1; i++ {
		if !AreClose(list[i], list[i+1]) || !AreContinous(list[i], list[i+1], shouldIncrease) {
			return false
		}
	}

	return true
}

func AreClose(a int, b int) bool {
	diff := math.Abs(float64(a - b))

	return diff >= 1 && diff <= 3
}

func AreContinous(a int, b int, shouldIncrease bool) bool {
	if shouldIncrease {
		return a < b
	}

	return a > b
}

func Day2Pt2() int {
	result := 0

	for _, row := range ReadInputNumbers() {
		if IsSafe(row) {
			result++
			continue
		}

		for i := 0; i < len(row); i++ {
			if IsSafe(Without(row, i)) {
				result++
				break
			}
		}
	}

	return result
}

func Without(list []int, index int) []int {
	copied := make([]int, index+1)
	copy(copied, list)
	return append(copied[:index], list[index+1:]...)
}
