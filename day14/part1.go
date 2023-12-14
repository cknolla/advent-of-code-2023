package day14

import (
	"advent-of-code-2023/utils"
	"fmt"
)

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

func tilt(locations [][]int, direction int) [][]int {
	switch direction {
	case NORTH:
		for rowIndex := 0; rowIndex < len(locations); rowIndex++ {
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
	case EAST:
		for rowIndex := 0; rowIndex < len(locations); rowIndex++ {
			for colIndex := len(locations[rowIndex]) - 2; colIndex >= 0; colIndex-- {
				if locations[rowIndex][colIndex] != ROUND_ROCK {
					continue
				}
				for index := colIndex; index < len(locations[rowIndex])-1; index++ {
					if locations[rowIndex][index+1] == SPACE {
						locations[rowIndex][index+1] = ROUND_ROCK
						locations[rowIndex][index] = SPACE
					} else {
						break
					}
				}
			}
		}
	case SOUTH:
		for rowIndex := len(locations) - 1; rowIndex >= 0; rowIndex-- {
			for colIndex := 0; colIndex < len(locations[rowIndex]); colIndex++ {
				if locations[rowIndex][colIndex] != ROUND_ROCK {
					continue
				}
				for index := rowIndex; index < len(locations)-1; index++ {
					if locations[index+1][colIndex] == SPACE {
						locations[index+1][colIndex] = ROUND_ROCK
						locations[index][colIndex] = SPACE
					} else {
						break
					}
				}
			}
		}
	case WEST:
		for rowIndex := 0; rowIndex < len(locations); rowIndex++ {
			for colIndex := 1; colIndex < len(locations[rowIndex]); colIndex++ {
				if locations[rowIndex][colIndex] != ROUND_ROCK {
					continue
				}
				for index := colIndex; index > 0; index-- {
					if locations[rowIndex][index-1] == SPACE {
						locations[rowIndex][index-1] = ROUND_ROCK
						locations[rowIndex][index] = SPACE
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

func printDish(locations *[][]int) {
	for _, location := range *locations {
		for _, object := range location {
			char := "."
			switch object {
			case ROUND_ROCK:
				char = "O"
			case CUBE_ROCK:
				char = "#"
			}
			fmt.Printf("%s", char)
		}
		fmt.Println("")
	}
}

func Part1(filename string) int {
	answer := 0
	locations := parseFile(filename)
	locations = tilt(locations, NORTH)
	answer = calculateLoad(&locations)
	return answer
}
