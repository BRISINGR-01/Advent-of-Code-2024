package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func Day3Pt1() int {
	result := 0

	for _, match := range Match(`mul\((\d{1,3}),(\d{1,3})\)`) {
		result += Multiply(match)
	}

	return result
}

func Match(pattern string) [][]string {
	return regexp.MustCompile(pattern).FindAllStringSubmatch(ReadInput(), -1)
}

func Multiply(match []string) int {
	num1, _ := strconv.Atoi(match[1])
	num2, _ := strconv.Atoi(match[2])

	return num1 * num2
}

func Day3Pt2() int {
	result := 0
	isEnabled := true

	pattern := fmt.Sprintf("%s|(%s)|(%s)", `mul\((\d{1,3}),(\d{1,3})\)`, `do\(\)`, `don\'t\(\)`)
	for _, match := range Match(pattern) {
		cmd := match[0]

		if cmd == "do()" {
			isEnabled = true
			continue
		}

		if cmd == "don't()" {
			isEnabled = false
			continue
		}

		if isEnabled {
			result += Multiply(match)
		}
	}

	return result
}
