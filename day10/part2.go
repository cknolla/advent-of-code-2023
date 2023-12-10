package day10

import (
	"fmt"
	"strings"
)

func (m *maze) getInsideCount() int {
	var insideCount int
	for _, row := range m.pipes {
		inside := false
		for _, pipe := range row {
			if !pipe.inLoop && inside {
				insideCount++
			} else if pipe.inLoop && strings.ContainsRune("|LJ", pipe.rune) {
				inside = !inside
			}
		}
	}
	return insideCount
}

func Part2() (string, error) {
	maze := parseFile("input.txt")
	maze.traverse()
	insideCount := maze.getInsideCount()
	return fmt.Sprintf("%d", insideCount), nil
}
