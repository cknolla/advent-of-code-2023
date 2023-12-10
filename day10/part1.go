package day10

import (
	"advent-of-code-2023/utils"
	"fmt"
	"strings"
)

type maze struct {
	runes       [][]rune
	startingLoc location
	rowCount    int
	colCount    int
}

type location struct {
	y int
	x int
}

func (l *location) Add(direction *location) {
	l.x += direction.x
	l.y += direction.y
}

func (l *location) nextNeighbor(r rune, prevLoc *location) location {
	directions := directionMap[r]
	if l.y+directions[0].y == prevLoc.y && l.x+directions[0].x == prevLoc.x {
		return location{l.y + directions[1].y, l.x + directions[1].x}
	} else {
		return location{l.y + directions[0].y, l.x + directions[0].x}
	}
}

var directionMap = map[rune][]location{
	'-': {{0, -1}, {0, 1}},
	'|': {{-1, 0}, {1, 0}},
	'L': {{0, 1}, {-1, 0}},
	'7': {{0, -1}, {1, 0}},
	'J': {{0, -1}, {-1, 0}},
	'F': {{0, 1}, {1, 0}},
}

// from starting traverse maze and return total steps taken
func (m *maze) traverse() (steps int) {
	currentLoc := m.startingLoc
	prevLoc := currentLoc
	// first order of business is finding out which way to go from start
	if currentLoc.x != 0 && strings.ContainsRune("-LF", m.runes[currentLoc.y][currentLoc.x-1]) {
		currentLoc.Add(&location{0, -1})
	} else if currentLoc.y != 0 && strings.ContainsRune("|LJ", m.runes[currentLoc.y-1][currentLoc.x]) {
		currentLoc.Add(&location{-1, 0})
	} else if currentLoc.x != m.rowCount-1 && strings.ContainsRune("-7J", m.runes[currentLoc.y][currentLoc.x+1]) {
		currentLoc.Add(&location{0, 1})
	} else {
		currentLoc.Add(&location{1, 0})
	}
	steps++
	// then continue finding other neighbor until encountering starting pos again
	for m.runes[currentLoc.y][currentLoc.x] != 'S' {
		nextLoc := currentLoc.nextNeighbor(m.runes[currentLoc.y][currentLoc.x], &prevLoc)
		prevLoc = currentLoc
		currentLoc = nextLoc
		steps++
	}
	return steps
}

func parseFile(filename string) maze {
	var m maze
	m.runes = make([][]rune, 140)
	for i := range m.runes {
		m.runes[i] = make([]rune, 140)
	}
	for l := range utils.GetInputLines(filename) {
		line := l.Text
		if line == "" {
			continue
		}
		for x, r := range l.Text {
			m.runes[l.Index][x] = r
			if r == 'S' {
				m.startingLoc = location{l.Index, x}
			}
			if m.rowCount == 0 {
				m.colCount++
			}
		}
		m.rowCount++
	}
	return m
}

func Part1() (string, error) {
	maze := parseFile("input.txt")
	steps := maze.traverse()
	return fmt.Sprintf("%d", steps/2), nil
}
