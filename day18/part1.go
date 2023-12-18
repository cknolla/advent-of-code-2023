package day18

import (
	"advent-of-code-2023/utils"
	"strconv"
	"strings"
)

type Dig struct {
	direction string
	distance  int
	color     string
}

func NewDig(line string) Dig {
	sections := strings.Fields(line)
	distance, _ := strconv.Atoi(sections[1])
	return Dig{
		direction: sections[0],
		distance:  distance,
		color:     sections[2][1:8],
	}
}

type Position struct {
	X     int
	Y     int
	color string
	dug   bool
}

type Grid struct {
	Positions [][]Position
	Width     int
	Height    int
}

var grid = Grid{}

func (g *Grid) executeDigPlan(digs []Dig) {
	currentX := 0
	currentY := 0
	axis := &currentX
	sign := 1
	g.Positions[0][0].dug = true
	for _, dig := range digs {
		switch dig.direction {
		case "R":
			axis = &currentX
			sign = 1
		case "L":
			axis = &currentX
			sign = -1
		case "U":
			axis = &currentY
			sign = -1
		default:
			axis = &currentY
			sign = 1
		}
		for i := 0; i < dig.distance; i++ {
			*axis += sign
			g.Positions[currentY][currentX].dug = true
		}
	}
}

func (g *Grid) getVolume() int {
	volume := 0
	for _, row := range g.Positions {
		inside := false
		prevDug := false
		for _, col := range row {
			if col.dug {
				if !prevDug {
					inside = !inside
				}
				volume++
				prevDug = true
			} else if inside {
				col.dug = true
				volume++
				prevDug = false
			} else {
				prevDug = false
			}
		}
	}
	return volume
}

func parseFile(filename string) (digPlan []Dig) {
	width := 0
	height := 0
	maxWidth := 0
	maxHeight := 0
	for l := range utils.GetInputLines(filename) {
		line := l.Text
		if line == "" {
			continue
		}
		dig := NewDig(line)
		switch dig.direction {
		case "R":
			width += dig.distance
		case "L":
			width -= dig.distance
		case "D":
			height += dig.distance
		default:
			height -= dig.distance
		}
		maxWidth = max(maxWidth, width)
		maxHeight = max(maxHeight, height)
		digPlan = append(digPlan, NewDig(line))
	}
	grid.Width = maxWidth + 1
	grid.Height = maxHeight + 1
	return digPlan
}

func Part1(filename string) int {
	answer := 0
	digPlan := parseFile(filename)
	grid.Positions = make([][]Position, grid.Height)
	for i := range grid.Positions {
		grid.Positions[i] = make([]Position, grid.Width)
	}
	grid.executeDigPlan(digPlan)
	answer = grid.getVolume()
	return answer
}
