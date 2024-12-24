package main

import (
	"strconv"
)

func Day22Pt1() int {
	result := 0

	for _, line := range ReadInputLines() {
		number, _ := strconv.Atoi(line)

		for i := 0; i < 2000; i++ {
			number = evolve(number)
		}

		result += number
	}

	return result
}

func evolve(number int) int {
	number = mix(number, number*64)
	number = mix(number, number/32)
	number = mix(number, number*2048)

	return number
}

func mix(a int, b int) int {
	return prune(a ^ b)
}

func prune(n int) int {
	return n % 16777216
}

type PriceData struct {
	price int
	seq   []int
}

func Day22Pt2() int {
	data := [][]PriceData{}
	allChanges := [][]int{}

	for monkeyI, line := range ReadInputLines() {
		println("monkey", monkeyI)
		data = append(data, []PriceData{})
		changes := []int{}
		number, _ := strconv.Atoi(line)

		bananas := get1s(number)
		prev := bananas

		for i := 1; i < 2000; i++ {
			number = evolve(number)
			bananas = get1s(number)
			changes = append(changes, bananas-prev)
			prev = bananas

			if len(changes) < 4 {
				continue
			}

			d := PriceData{
				price: bananas,
				seq:   changes[len(changes)-4:],
			}

			data[monkeyI] = append(data[monkeyI], d)
		}

		allChanges = append(allChanges, changes)
	}

	highest := 0
	// best := data[0][0:100]

	// for i, tradeData := range data {
	// 	for _, trade := range tradeData {
	// 		// if slices.ContainsFunc()
	// 	}
	// }

	return highest
}

func calc(data [][]PriceData, seq []int) int {
	result := 0

	for _, tradeData := range data {
		for _, trade := range tradeData {
			if areEqual(trade.seq, seq) {
				result += trade.price
				break
			}
		}
	}

	return result
}

func get1s(n int) int {
	return n % 10
}

func areEqual(a []int, b []int) bool {
	return a[0] == b[0] && a[1] == b[1] && a[2] == b[2] && a[3] == b[3]
}

// 15219 too high
// 1597 too low
