package main

import (
	"log"
	"os/exec"
	"strconv"
)

func main() {
	result := Day14Pt2()
	println("Result: ", result)

	cmd := exec.Command("wl-copy", strconv.Itoa(result))
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
