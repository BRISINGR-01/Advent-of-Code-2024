package main

import (
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Ctx struct {
	A      int
	B      int
	C      int
	IP     int
	output []string
}

func Day17Pt1() string {
	input := ReadInputLines()
	ctx := &Ctx{}

	matches := regexp.MustCompile(`Register A: (\d+)`).FindStringSubmatch(input[0])
	a, _ := strconv.Atoi(matches[1])
	ctx.A = a
	matches = regexp.MustCompile(`Register B: (\d+)`).FindStringSubmatch(input[1])
	b, _ := strconv.Atoi(matches[1])
	ctx.B = b
	matches = regexp.MustCompile(`Register C: (\d+)`).FindStringSubmatch(input[2])
	c, _ := strconv.Atoi(matches[1])
	ctx.C = c

	memory := []int{}
	for _, v := range regexp.MustCompile(`(\d)`).FindAllStringSubmatch(input[len(input)-1], -1) {
		opcode, _ := strconv.ParseInt(v[1], 10, 8)
		memory = append(memory, int(opcode))
	}

	execF(memory, ctx)

	return strings.Join(ctx.output, ",")
}

func getCombo(operand int, ctx *Ctx) int {
	switch operand {
	case 0, 1, 2, 3:
		return int(operand)
	case 4:
		return ctx.A
	case 5:
		return ctx.B
	case 6:
		return ctx.C
	case 7:
		panic("Invalid operand")
	}

	panic("Invalid operand")
}

func execF(memory []int, ctx *Ctx) {
	opcode := memory[ctx.IP]
	literal := memory[ctx.IP+1]

	var f func(val int, ctx *Ctx)

	switch opcode {
	case 0:
		f = adv
	case 1:
		f = bxl
	case 2:
		f = bst
	case 3:
		f = jnz
	case 4:
		f = bxc
	case 5:
		f = out
	case 6:
		f = bdv
	case 7:
		f = cdv
	}

	f(literal, ctx)

	ctx.IP += 2

	if ctx.IP < len(memory) {
		execF(memory, ctx)
	}
}

func adv(val int, ctx *Ctx) {
	ctx.A /= int(math.Pow(2, float64(getCombo(val, ctx))))
}

func bxl(val int, ctx *Ctx) {
	ctx.B ^= val
}

func bst(val int, ctx *Ctx) {
	ctx.B = getCombo(val, ctx) % 8
}

func jnz(val int, ctx *Ctx) {
	if ctx.A == 0 {
		return
	}

	ctx.IP = val - 2
}

func bxc(_ int, ctx *Ctx) {
	ctx.B ^= ctx.C
}

func out(val int, ctx *Ctx) {
	ctx.output = append(ctx.output, strconv.Itoa(getCombo(val, ctx)%8))
}

func bdv(val int, ctx *Ctx) {
	ctx.B = ctx.A / int(math.Pow(2, float64(getCombo(val, ctx))))
}

func cdv(val int, ctx *Ctx) {
	ctx.C = ctx.A / int(math.Pow(2, float64(getCombo(val, ctx))))
}

func Day17Pt2() int {
	input := ReadInputLines()
	ctx := &Ctx{}

	matches := regexp.MustCompile(`Register A: (\d+)`).FindStringSubmatch(input[0])
	a, _ := strconv.Atoi(matches[1])
	ctx.A = a
	matches = regexp.MustCompile(`Register B: (\d+)`).FindStringSubmatch(input[1])
	b, _ := strconv.Atoi(matches[1])
	ctx.B = b
	matches = regexp.MustCompile(`Register C: (\d+)`).FindStringSubmatch(input[2])
	c, _ := strconv.Atoi(matches[1])
	ctx.C = c

	memory := []int{}
	for _, v := range regexp.MustCompile(`(\d)`).FindAllStringSubmatch(input[len(input)-1], -1) {
		opcode, _ := strconv.ParseInt(v[1], 10, 8)
		memory = append(memory, int(opcode))
	}

	mem := slices.Clone(memory)
	slices.Reverse(mem)
	// A/=8
	// >> A%8

	// 4 - (3*8+4)*8
	// 3 - 3*8
	// 0 - 0

	// B = A % 8
	// B^= 1
	// C = A / B**2
	// A /= 8
	// B ^= 4
	// B ^= C
	// >> B%8

	// B2= (A1%8)^1  x=( (B2^4) ^ (A1/(B2*B2)) )%8 A2=A1/8
	// x ^ 5 = A1^( A1/(A1^1**2) )

	for i := 0; i < 9; i++ {
		// ctx.A = i
		// ctx.B = 0
		// ctx.C = 0
		// ctx.IP = 0
		// A1 := ctx.A
		// B2 := (A1 % 8) ^ 1
		// A2 := A1 / 8

		// x := ((((A1 % 8) ^ 1) ^ 4) ^ (A1 / (((A1 % 8) ^ 1) * ((A1 % 8) ^ 1)))) % 8

		// x := A1

		// if x < 2 {
		// 	x = 0
		// }

		// x = (x ^ 5) % 8
		// x = (A1 % 8) ^ (5 ^ (A1 / int(math.Pow(float64((A1%8)^1), 2))))

		// x ^ 5 = A1^( A1/(A1^1**2) )

	}

	ctx.A = 0
	slices.Reverse(memory)
	for _, bit := range memory {
		bit4 := bit ^ 4
		bitSq := int(math.Pow(float64(bit4), 2))
		if bitSq == 0 {
			bitSq ^= 1
		}

		C := ctx.A / bitSq
		ctx.A *= 8
		ctx.A += (C ^ bit ^ 5) % 8

		println(ctx.A, bit, ctx.A%8)
	}
	slices.Reverse(memory)

	res := ctx.A

	// ctx.A = ((0^5)*8+3^5)*8 + 5 ^ 5

	execF(memory, ctx)
	println(strings.Join(ctx.output, ","))

	return res
}
