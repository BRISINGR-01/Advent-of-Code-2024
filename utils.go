package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadInput() string {
	input, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	return string(input)
}

func ReadInputLines() []string {
	res := strings.Split(ReadInput(), "\n")

	// remove empty last line
	if res[len(res)-1] == "" {
		res = res[:len(res)-1]
	}

	return res
}

func ReadInputStrings() [][]string {
	input := ReadInputLines()
	result := [][]string{}

	for _, row := range input {
		result = append(result, regexp.MustCompile(`\s+`).Split(row, -1))
	}

	return result
}

func ReadInputNumbers() [][]int {
	input := ReadInputLines()
	result := [][]int{}

	for _, row := range input {
		raw_numbers := regexp.MustCompile(`\s+`).Split(row, -1)
		numbers := []int{}

		for _, raw_number := range raw_numbers {
			num, _ := strconv.Atoi(raw_number)
			numbers = append(numbers, num)
		}

		result = append(result, numbers)
	}

	return result
}
