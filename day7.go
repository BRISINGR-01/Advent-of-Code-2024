package main

import (
	"regexp"
	"strconv"
)

func Day7Pt1() int {
	result := 0

	for _, row := range ReadInputNumbers() {
		sum := row[0]

		if CanBeCombined(sum, row[1:]) {
			result += sum
		}
	}

	return result
}

func CanBeCombined(sum int, numbers []int) bool {
	if len(numbers) == 1 {
		return sum == numbers[0]
	}

	lastIndex := len(numbers) - 1
	lastNum := numbers[lastIndex]

	if CanBeCombined(sum-lastNum, numbers[:lastIndex]) {
		return true
	} else if sum%lastNum == 0 && CanBeCombined(sum/lastNum, numbers[:lastIndex]) {
		return true
	}

	return false
}

func Day7Pt2() int {
	result := 0

	for _, row := range ReadInputNumbers() {
		sum := row[0]

		if CanBeCombined2(sum, row[1:]) {
			result += sum
		}
	}

	return result
}

func CanBeCombined2(sum int, numbers []int) bool {
	if len(numbers) == 1 {
		return sum == numbers[0]
	}

	lastIndex := len(numbers) - 1
	lastNum := numbers[lastIndex]

	if CanBeCombined2(sum-lastNum, numbers[:lastIndex]) {
		return true
	} else if sum%lastNum == 0 && CanBeCombined2(sum/lastNum, numbers[:lastIndex]) {
		return true
	}

	sumStr := strconv.Itoa(sum)
	lastNumStr := strconv.Itoa(lastNum)

	endsWithNrRefex := regexp.MustCompile(lastNumStr + "$")

	if !endsWithNrRefex.MatchString(sumStr) {
		return false
	}

	sumStr = endsWithNrRefex.ReplaceAllString(sumStr, "")

	sum, _ = strconv.Atoi(sumStr)

	return CanBeCombined2(sum, numbers[:lastIndex])
}
