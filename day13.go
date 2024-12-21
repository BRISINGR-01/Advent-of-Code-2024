package main

import (
	"regexp"
	"strconv"
)

type Button struct {
	X int
	Y int
}

type Game struct {
	A      Button
	B      Button
	prizeX int
	prizeY int
}

func ExtractButton(input string) Button {
	match := regexp.MustCompile(`Button \w: X\+(\d+), Y\+(\d+)`).FindAllStringSubmatch(input, -1)[0]

	x, _ := strconv.Atoi(match[1])
	y, _ := strconv.Atoi(match[2])

	return Button{
		X: x,
		Y: y,
	}
}

func ExtractGames() []Game {
	games := []Game{}
	input := ReadInputLines()

	for i := 0; i < len(input); i += 4 {
		prizeMatch := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`).FindAllStringSubmatch(input[i+2], -1)[0]
		prizeX, _ := strconv.Atoi(prizeMatch[1])
		prizeY, _ := strconv.Atoi(prizeMatch[2])

		games = append(games, Game{
			A:      ExtractButton(input[i]),
			B:      ExtractButton(input[i+1]),
			prizeX: prizeX,
			prizeY: prizeY,
		})
	}

	return games
}

func CountTokens(game Game) int {
	const MaxInt = int(^uint(0) >> 1)
	tokens := MaxInt

	for cA := 0; cA <= 100; cA++ {
		for cB := 0; cB <= 100; cB++ {
			if CheckScore(game, cA, cB) {
				continue
			}

			tokens = min(tokens, cA*3+cB)
		}
	}

	if tokens == MaxInt {
		return 0
	}

	return tokens
}

func CheckScore(game Game, cA int, cB int) bool {
	return cA*game.A.X+cB*game.B.X == game.prizeX && cA*game.A.Y+cB*game.B.Y == game.prizeY
}

// 35291 too high

func Day13Pt1() int {
	result := 0

	games := ExtractGames()

	for _, game := range games {
		result += CountTokens(game)
	}

	return result
}

func ExtractGames2() []Game {
	games := []Game{}
	input := ReadInputLines()

	for i := 0; i < len(input); i += 4 {
		prizeMatch := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`).FindAllStringSubmatch(input[i+2], -1)[0]
		prizeX, _ := strconv.Atoi(prizeMatch[1])
		prizeY, _ := strconv.Atoi(prizeMatch[2])

		games = append(games, Game{
			A:      ExtractButton(input[i]),
			B:      ExtractButton(input[i+1]),
			prizeX: prizeX + 10000000000000,
			prizeY: prizeY + 10000000000000,
		})
	}

	return games
}

func Day13Pt2() int {
	result := 0

	games := ExtractGames2()

	for _, game := range games {
		result += CountTokens2(game)
	}

	return result
}

func CountTokens2(game Game) int {
	const MaxInt = int(^uint(0) >> 1)
	tokens := MaxInt

	for cA := 1000000000000 / (game.A.X + game.B.X); cA <= 10000000000000/(game.A.X+game.B.X)+1000; cA++ {
		for cB := 1000000000000 / (game.A.X + game.B.X); cB <= 10000000000000/(game.A.X+game.B.X)+1000; cB++ {
			sum := cA*game.A.X + cB*game.B.X

			if sum == game.prizeX && cA*game.A.Y+cB*game.B.Y == game.prizeY {
				// tokens := min(tokens, cA*3+cB)
				println(cA, cB)
				return cA*3 + cB
			}

		}
	}

	if tokens == MaxInt {
		return 0
	}

	return tokens
}
