package main

import (
	"os"
	"strings"
)

func ReadInput() []string {
	input, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	return strings.Split(string(input), "\n")
}
