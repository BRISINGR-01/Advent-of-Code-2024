package main

import (
	"math"
	"sort"
)

func Day1Pt1() int {
	left := []int{}
	right := []int{}

	for _, row := range ReadInputNumbers() {
		left = append(left, row[0])
		right = append(right, row[1])
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})
	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	result := 0

	for i, leftVal := range left {
		result += int(math.Abs(float64(leftVal - right[i])))
	}

	return result
}

func Day1Pt2() int {
	left := []int{}
	frequency := map[int]int{}

	for _, row := range ReadInputNumbers() {
		left = append(left, row[0])
		frequency[row[1]] += 1
	}

	result := 0

	for _, v := range left {
		result += v * frequency[v]
	}

	return result
}
