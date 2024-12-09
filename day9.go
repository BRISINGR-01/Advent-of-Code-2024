package main

type MemBlock struct {
	id     int
	isFree bool
}

func Day9Pt1() int {
	result := 0

	input := ReadInputRunes()[0]
	filesystem := []MemBlock{}

	id := 0
	for i, char := range input {
		size := int(char - '0')

		memory := MemBlock{id: -1, isFree: true}

		if i%2 == 0 {
			memory = MemBlock{id: id, isFree: false}
			id++
		}

		for ii := 0; ii < size; ii++ {
			filesystem = append(filesystem, memory)
		}
	}

	lastIndex := len(filesystem) - 1
	for i := 0; i <= lastIndex; i++ {
		memory := filesystem[i]

		if !memory.isFree {
			// result += memory.id * i
			continue
		}

		if i == lastIndex {
			break
		}

		for filesystem[lastIndex].isFree && lastIndex > i {
			lastIndex--
		}

		filesystem[i] = filesystem[lastIndex]
		filesystem[lastIndex].isFree = true
	}

	for i := 0; i < len(filesystem); i++ {
		if filesystem[i].isFree {
			break
		} else {
			result += filesystem[i].id * i
		}
	}

	return result
}

func Day9Pt2() int {
	result := 0

	input := ReadInputRunes()[0]
	filesystem := []rune{}

	id := 0
	for i, char := range input {
		size := int(char - '0')

		memory := '.'

		if i%2 == 0 {
			memory = rune(id + '0')
			id++
		}

		for ii := 0; ii < size; ii++ {
			filesystem = append(filesystem, memory)
		}
	}

	for endMemI := len(filesystem) - 1; endMemI >= 0; endMemI-- {
		endMemory := filesystem[endMemI]

		if endMemory == '.' {
			continue
		}

		fileSize := 0

		for ; endMemI >= 0 && filesystem[endMemI] == endMemory; endMemI-- {
			fileSize++
		}
		endMemI++

		for startMemoryI, startMemory := range filesystem[:endMemI] {
			if startMemory != '.' {
				continue
			}

			emptySpace := 1

			for freeI := startMemoryI + 1; filesystem[freeI] == '.'; freeI++ {
				emptySpace++
			}

			if emptySpace < fileSize {
				continue
			}

			for freeI := startMemoryI; freeI < startMemoryI+fileSize; freeI++ {
				filesystem[freeI] = endMemory
			}

			for occupiedI := endMemI; occupiedI < endMemI+fileSize; occupiedI++ {
				filesystem[occupiedI] = '.'
			}

			break
		}
	}

	for i, memory := range filesystem {
		if memory == '.' {
			continue
		} else {
			result += int(memory-'0') * i
		}
	}

	return result
}

// 6353923330830 is too high
// 6353658451014
// 164063653075 wrong
// 90233883775 is too low
// 5598524210 is too low
