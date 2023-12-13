package day12

import (
	"fmt"
	"math"
)

func (r *row) unfold() ([]rune, []int) {
	extendedSprings := make([]rune, len(r.springs)*2+1)
	criteria := make([]int, len(r.criteria)*2)
	copy(extendedSprings, r.springs)
	originalSprings := make([]rune, len(r.springs))
	copy(originalSprings, r.springs)
	extendedSprings[len(r.springs)] = '?'
	for i := len(r.springs) + 1; i < len(extendedSprings); i++ {
		extendedSprings[i] = r.springs[i-(len(r.springs)+1)]
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < len(r.criteria); j++ {
			criteria[i*len(r.criteria)+j] = r.criteria[j]
		}
	}
	return extendedSprings, criteria
}

// ended up copying this. wasn't gonna figure it out
// https://raw.githubusercontent.com/Oupsman/AOC2023/main/d12/advent_12.go
func Part2Alt() (string, error) {
	sum := 0
	rows := parseFile("input.txt")
	for index, row := range rows {
		fmt.Printf("Starting row %d\n", index)
		extendedSprings, criteria := row.unfold()
		originalArrangements := row.testAll()
		row.springs = extendedSprings
		row.criteria = criteria
		secondArrangements := row.testAll()
		sum += originalArrangements * int(math.Pow(float64(secondArrangements)/float64(originalArrangements), float64(4)))
		fmt.Printf("Completed row %d\n", index)
	}
	return fmt.Sprintf("%d", sum), nil
}
