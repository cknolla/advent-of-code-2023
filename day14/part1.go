package day14

import "advent-of-code-2023/utils"

const (
	SPACE = iota
	ROUND_ROCK
	CUBE_ROCK
)

const (
	NORTH = iota
	EAST
	SOUTH
	WEST
)

//type location struct {
//	occupant int
//
//}

func tilt(locations [][]int, direction int) [][]int {
	for rowIndex := 0; rowIndex < len(locations); rowIndex++ {
		switch direction {
		case NORTH:
			for colIndex := 0; colIndex < len(locations[rowIndex]); colIndex++ {
				if locations[rowIndex][colIndex] != ROUND_ROCK {
					continue
				}
				for index := rowIndex; index > 0; index-- {
					if locations[index-1][colIndex] == SPACE {
						locations[index-1][colIndex] = ROUND_ROCK
						locations[index][colIndex] = SPACE
					} else {
						break
					}
				}
			}
		}
	}
	return locations
}

func calculateLoad(locations *[][]int) (totalLoad int) {
	for rowIndex, location := range *locations {
		for _, object := range location {
			if object == ROUND_ROCK {
				totalLoad += len(*locations) - rowIndex
			}
		}
	}
	return totalLoad
}

func parseFile(filename string) [][]int {
	var locations [][]int
	for l := range utils.GetInputLines(filename) {
		line := l.Text
		if line == "" {
			continue
		}
		row := make([]int, len(line))
		for index, char := range line {
			switch char {
			case '.':
				row[index] = SPACE
			case 'O':
				row[index] = ROUND_ROCK
			default:
				row[index] = CUBE_ROCK
			}
		}
		locations = append(locations, row)
	}
	return locations
}

func Part1(filename string) int {
	answer := 0
	locations := parseFile(filename)
	locations = tilt(locations, NORTH)
	answer = calculateLoad(&locations)
	return answer
}
