package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func Day1Pt1() int {

	left := []int{}
	right := []int{}

	for _, row := range ReadInput() {
		numbers := strings.Split(row, "   ")
		leftInt, _ := strconv.Atoi(numbers[0])
		rightInt, _ := strconv.Atoi(numbers[1])

		left = append(left, leftInt)
		right = append(right, rightInt)
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})
	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	result := 0

	for i := 0; i < len(left); i++ {
		result += int(math.Abs(float64(left[i] - right[i])))
	}

	return result
}

func Day1Pt2() int {
	left := []int{}
	frequency := map[int]int{}

	for _, row := range ReadInput() {
		numbers := strings.Split(row, "   ")
		leftInt, _ := strconv.Atoi(numbers[0])
		rightInt, _ := strconv.Atoi(numbers[1])

		left = append(left, leftInt)
		frequency[rightInt] += 1
	}

	result := 0

	for _, v := range left {
		result += v * frequency[v]
	}

	return result
}
