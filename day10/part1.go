package day10

import (
	"advent-of-code-2023/utils"
	"fmt"
	"strings"
)

type maze struct {
	pipes       [][]pipe
	startingLoc location
	rowCount    int
	colCount    int
}

type location struct {
	y int
	x int
}

type pipe struct {
	rune   rune
	inLoop bool
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
	m.pipes[currentLoc.y][currentLoc.x].inLoop = true
	// first order of business is finding out which way to go from start
	if currentLoc.x != 0 && strings.ContainsRune("-LF", m.pipes[currentLoc.y][currentLoc.x-1].rune) {
		currentLoc.Add(&location{0, -1})
	} else if currentLoc.y != 0 && strings.ContainsRune("|LJ", m.pipes[currentLoc.y-1][currentLoc.x].rune) {
		currentLoc.Add(&location{-1, 0})
	} else if currentLoc.x != m.rowCount-1 && strings.ContainsRune("-7J", m.pipes[currentLoc.y][currentLoc.x+1].rune) {
		currentLoc.Add(&location{0, 1})
	} else {
		currentLoc.Add(&location{1, 0})
	}
	m.pipes[currentLoc.y][currentLoc.x].inLoop = true
	steps++
	// then continue finding other neighbor until encountering starting pos again
	for m.pipes[currentLoc.y][currentLoc.x].rune != 'S' {
		nextLoc := currentLoc.nextNeighbor(m.pipes[currentLoc.y][currentLoc.x].rune, &prevLoc)
		prevLoc = currentLoc
		currentLoc = nextLoc
		m.pipes[currentLoc.y][currentLoc.x].inLoop = true
		steps++
	}
	return steps
}

func parseFile(filename string) maze {
	var m maze
	m.pipes = make([][]pipe, 140)
	for i := range m.pipes {
		m.pipes[i] = make([]pipe, 140)
	}
	for l := range utils.GetInputLines(filename) {
		line := l.Text
		if line == "" {
			continue
		}
		for x, r := range l.Text {
			m.pipes[l.Index][x].rune = r
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
