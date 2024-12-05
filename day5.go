package main

import (
	"regexp"
	"slices"
	"strconv"
)

type Order struct {
	before int
	after  int
}

func Day5Pt1() int {
	result := 0

	orders := []Order{}
	input := ReadInputLines()

	isFillingOrders := true
	for _, row := range input {
		if row == "" {
			isFillingOrders = false
			continue
		}

		if isFillingOrders {
			orders = append(orders, ParseOrder(row))
		} else {
			updates := ParseUpdates(row)
			if IsCorrect(updates, orders) {
				result += updates[len(updates)/2]
			}
		}
	}

	return result
}

func ParseOrder(row string) Order {
	matches := regexp.MustCompile(`(\d+)\|(\d+)`).FindStringSubmatch(row)

	before, _ := strconv.Atoi(matches[1])
	after, _ := strconv.Atoi(matches[2])

	return Order{
		before: before,
		after:  after,
	}
}

func ParseUpdates(row string) []int {
	matches := regexp.MustCompile(`(\d+)`).FindAllStringSubmatch(row, -1)

	result := []int{}

	for _, match := range matches {
		num, _ := strconv.Atoi(match[1])
		result = append(result, num)
	}

	return result
}

func IsCorrect(updates []int, orders []Order) bool {
	for _, order := range orders {
		before_i := slices.Index(updates, order.before)
		after_i := slices.Index(updates, order.after)

		if before_i == -1 || after_i == -1 {
			continue
		}

		if before_i >= after_i {
			return false
		}
	}

	return true
}

func Day5Pt2() int {
	result := 0

	orders := []Order{}
	input := ReadInputLines()

	isFillingOrders := true
	for _, row := range input {
		if row == "" {
			isFillingOrders = false
			continue
		}

		if isFillingOrders {
			orders = append(orders, ParseOrder(row))
		} else {
			updates := ParseUpdates(row)
			if !IsCorrect(updates, orders) {
				updates = SortUpdates(updates, orders)

				result += updates[len(updates)/2]
			}
		}
	}

	return result
}

func SortUpdates(updates []int, orders []Order) []int {
	slices.SortFunc(updates, func(a int, b int) int {
		for _, order := range orders {
			if a == order.before && b == order.after {
				return 1
			}

			if a == order.after && b == order.before {
				return -1

			}
		}

		return 0
	})

	return updates
}
