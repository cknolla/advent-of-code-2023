package day11

import "fmt"

func Part2() (string, error) {
	universe := parseFile("input.txt", 1000000)
	universe.expandColumns()
	answer := universe.getShortestDistances()
	return fmt.Sprintf("%d", answer), nil
}
