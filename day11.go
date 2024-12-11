package main

import (
	"strconv"
)

func Day11Pt1() int {
	result := 0
	steps := 25 // 75 for Pt2
	cache := []map[int]int{}

	for i := 0; i <= steps; i++ {
		cache = append(cache, map[int]int{})
	}

	for _, stone := range ReadInputNumbers()[0] {
		result += Process(stone, steps, cache)
	}

	return result
}

func Process(stone int, step int, cache []map[int]int) int {
	if cache[step][stone] != 0 {
		return cache[step][stone]
	}

	result := 1

	for i := 1; i <= step; i++ {
		stone1, stone2 := NextStep(stone)
		stone = stone1
		if stone2 != -1 {
			cache[step-i][stone2] = Process(stone2, step-i, cache)
			result += cache[step-i][stone2]
		}
	}

	return result
}

func NextStep(stone int) (int, int) {
	if stone == 0 {
		return 1, -1
	}

	str := strconv.Itoa(stone)

	if len(str)%2 == 0 {
		middle := len(str) / 2
		stone1, _ := strconv.Atoi(str[:middle])
		stone2, _ := strconv.Atoi(str[middle:])

		return stone1, stone2
	}

	return stone * 2024, -1
}
