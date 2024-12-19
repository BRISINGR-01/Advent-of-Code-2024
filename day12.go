package main

type Plant struct {
	is_counted bool
	pl_type    rune
	neighbours []*Plant
}

func Day12Pt1() int {
	result := 0
	grid := [][]*Plant{}

	for y, row := range ReadInputRunes() {
		grid = append(grid, make([]*Plant, len(row)))
		for x, val := range row {
			grid[y][x] = &Plant{pl_type: val, is_counted: false}
		}
	}

	for y, row := range grid {
		for x, plant := range row {
			CollectNeighbours(plant, y, x, grid)
		}
	}

	for _, row := range grid {
		for _, plant := range row {
			area, perimeter := ClalcFence(plant)

			result += area * perimeter
		}
	}

	return result
}

func CollectNeighbours(plant *Plant, y, x int, grid [][]*Plant) {
	if y > 0 && grid[y-1][x].pl_type == plant.pl_type {
		plant.neighbours = append(plant.neighbours, grid[y-1][x])
	}

	if y < len(grid)-1 && grid[y+1][x].pl_type == plant.pl_type {
		plant.neighbours = append(plant.neighbours, grid[y+1][x])
	}

	if x > 0 && grid[y][x-1].pl_type == plant.pl_type {
		plant.neighbours = append(plant.neighbours, grid[y][x-1])
	}

	if x < len(grid[y])-1 && grid[y][x+1].pl_type == plant.pl_type {
		plant.neighbours = append(plant.neighbours, grid[y][x+1])
	}
}

func ClalcFence(plant *Plant) (int, int) {
	if plant.is_counted {
		return 0, 0
	}

	plant.is_counted = true
	area := 1
	perimeter := 4

	for _, neighbour := range plant.neighbours {
		neighbour_area, neighbour_perimeter := ClalcFence(neighbour)
		area += neighbour_area
		perimeter += neighbour_perimeter - 1
	}

	return area, perimeter
}

type Plant2 struct {
	is_counted bool
	pl_type    rune
	neighbours []*Plant2
	sides      []Side
}

type Side struct {
	start       int
	end         int
	axis        int
	liesOnXAxis bool
}

func Day12Pt2() int {
	result := 0
	grid := [][]*Plant2{}

	for y, row := range ReadInputRunes() {
		grid = append(grid, make([]*Plant2, len(row)))
		for x, val := range row {
			grid[y][x] = &Plant2{pl_type: val, is_counted: false}
		}
	}

	for y, row := range grid {
		for x, plant := range row {
			CollectNeighbours2(plant, y, x, grid)
		}
	}

	for _, row := range grid {
		for _, plant := range row {
			area, sides := ClalcFence2(plant)
			prev := 0

			for len(sides) != prev {
				prev = len(sides)
				sides = CombineSides(sides)
			}

			if area != 0 {
				println(string(plant.pl_type), area, len(sides))
			}

			result += area * len(sides)
		}
	}

	return result
}

func CollectNeighbours2(plant *Plant2, y, x int, grid [][]*Plant2) {
	if y > 0 && grid[y-1][x].pl_type == plant.pl_type {
		plant.neighbours = append(plant.neighbours, grid[y-1][x])
	} else {
		plant.sides = append(plant.sides, Side{start: x, end: x + 1, axis: y + 1, liesOnXAxis: true})
	}

	if y < len(grid)-1 && grid[y+1][x].pl_type == plant.pl_type {
		plant.neighbours = append(plant.neighbours, grid[y+1][x])
	} else {
		plant.sides = append(plant.sides, Side{start: x, end: x + 1, axis: y, liesOnXAxis: true})
	}

	if x > 0 && grid[y][x-1].pl_type == plant.pl_type {
		plant.neighbours = append(plant.neighbours, grid[y][x-1])
	} else {
		plant.sides = append(plant.sides, Side{start: y, end: y + 1, axis: x, liesOnXAxis: false})
	}

	if x < len(grid[y])-1 && grid[y][x+1].pl_type == plant.pl_type {
		plant.neighbours = append(plant.neighbours, grid[y][x+1])
	} else {
		plant.sides = append(plant.sides, Side{start: y, end: y + 1, axis: x + 1, liesOnXAxis: false})
	}
}

func ClalcFence2(plant *Plant2) (int, []Side) {
	if plant.is_counted {
		return 0, []Side{}
	}

	plant.is_counted = true
	area := 1

	for _, neighbour := range plant.neighbours {
		neighbour_area, neighbour_sides := ClalcFence2(neighbour)
		area += neighbour_area
		plant.sides = append(plant.sides, neighbour_sides...)
	}

	return area, plant.sides
}

func CombineSides(sides []Side) []Side {
	result := []Side{}

	for _, side := range sides {
		is_found := false
		for i, res_side := range result {
			if res_side.liesOnXAxis != side.liesOnXAxis || res_side.axis != side.axis {
				continue
			}

			if res_side.start == side.end {
				is_found = true
				res_side.start = side.start
				result[i] = res_side
				break
			}

			if res_side.end == side.start {
				is_found = true
				res_side.end = side.end
				result[i] = res_side
				break
			}
		}

		if !is_found {
			result = append(result, side)
		}
	}

	return result
}
