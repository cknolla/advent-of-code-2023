package day11

import (
	"advent-of-code-2023/utils"
	"fmt"
	"math"
	"slices"
)

type universe struct {
	galaxies       []location
	emptyColumns   []int
	expansionRatio int
}

type location struct {
	y int
	x int
}

func (u *universe) expandColumns() {
	for galaxyIndex, galaxy := range u.galaxies {
		for i, x := range u.emptyColumns {
			if x < galaxy.x {
				u.galaxies[galaxyIndex].x += (len(u.emptyColumns)-i)*u.expansionRatio - (len(u.emptyColumns) - i)
				break
			}
		}
	}
}

func (u *universe) getShortestDistances() int {
	n := len(u.galaxies)
	sum := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			shortestDistance := int(math.Abs(float64(u.galaxies[i].y-u.galaxies[j].y)) + math.Abs(float64(u.galaxies[i].x-u.galaxies[j].x)))
			sum += shortestDistance
		}
	}
	return sum
}

func parseFile(filename string, expansionRatio int) universe {
	var u universe
	u.expansionRatio = expansionRatio
	rowsWithoutGalaxies := 0
	var populatedColumns []bool
	for l := range utils.GetInputLines(filename) {
		line := l.Text
		if line == "" {
			continue
		}
		if populatedColumns == nil {
			populatedColumns = make([]bool, len(line))
		}
		rowHasGalaxy := false
		for col, char := range line {
			if char == '#' {
				rowHasGalaxy = true
				populatedColumns[col] = true
				// expand rows on the way
				u.galaxies = append(u.galaxies, location{l.Index + (rowsWithoutGalaxies * u.expansionRatio) - rowsWithoutGalaxies, col})
			}
		}
		if !rowHasGalaxy {
			rowsWithoutGalaxies++
		}
	}
	for i, populatedColumn := range populatedColumns {
		if !populatedColumn {
			u.emptyColumns = append(u.emptyColumns, i)
		}
	}
	slices.Reverse(u.emptyColumns)
	return u
}

// parse in galaxies
// determine empty rows and columns along the way
// expand galaxies
// find shortest distance between each pair with abs(1.x-0.x) + abs(1.y-0.y)
// add up shortest distances
func Part1() (string, error) {
	universe := parseFile("input.txt", 2)
	universe.expandColumns()
	answer := universe.getShortestDistances()
	return fmt.Sprintf("%d", answer), nil
}
