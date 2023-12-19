package day18

import (
	"advent-of-code-2023/utils"
	"fmt"
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
	X        int
	Y        int
	color    string
	dug      bool
	filledIn bool
	char     rune
}

func charLookup(currentDirection string, nextDirection string) rune {
	switch currentDirection {
	case "R":
		if nextDirection == "U" {
			return 'J'
		}
		if nextDirection == "D" {
			return '7'
		}
	case "D":
		if nextDirection == "L" {
			return 'J'
		}
		if nextDirection == "R" {
			return 'L'
		}
	case "L":
		if nextDirection == "U" {
			return 'L'
		}
		if nextDirection == "D" {
			return 'F'
		}
	case "U":
		if nextDirection == "L" {
			return '7'
		}
		if nextDirection == "R" {
			return 'F'
		}
	}
	return '|'
}

type Grid struct {
	Positions [][]Position
	Width     int
	Height    int
	StartingX int
	StartingY int
}

var grid = Grid{}

func (g *Grid) executeDigPlan(digs []Dig) {
	currentX := g.StartingX
	currentY := g.StartingY
	axis := &currentX
	sign := 1
	char := '-'
	for digIndex, dig := range digs {
		switch dig.direction {
		case "R":
			axis = &currentX
			sign = 1
			char = '-'
		case "L":
			axis = &currentX
			sign = -1
			char = '-'
		case "U":
			axis = &currentY
			sign = -1
			char = '|'
		default:
			axis = &currentY
			sign = 1
			char = '|'
		}
		for i := 0; i < dig.distance; i++ {
			*axis += sign
			g.Positions[currentY][currentX].dug = true
			g.Positions[currentY][currentX].char = char
		}
		index := digIndex + 1
		if index == len(digs) {
			index = 0
		}
		nextDirection := digs[index].direction
		g.Positions[currentY][currentX].char = charLookup(dig.direction, nextDirection)
	}
}

func (g *Grid) getVolume() int {
	volume := 0
	for y, row := range g.Positions {
		inside := false
		for x, col := range row {
			if strings.ContainsRune("|JL", col.char) {
				inside = !inside
			}
			if inside {
				if !col.dug {
					g.Positions[y][x].filledIn = true
				}
			}
			if col.dug || g.Positions[y][x].filledIn {
				volume++
			}
		}
	}
	return volume
}

func (g *Grid) display() {
	for _, row := range g.Positions {
		for _, col := range row {
			c := '.'
			if col.dug {
				c = col.char
			} else if col.filledIn {
				c = '+'
			}
			fmt.Printf("%c", c)
		}
		fmt.Println("")
	}
}

func parseFile(filename string) (digPlan []Dig) {
	x := 0
	y := 0
	maxWidth := 0
	maxHeight := 0
	minWidth := 0
	minHeight := 0
	for l := range utils.GetInputLines(filename) {
		line := l.Text
		if line == "" {
			continue
		}
		dig := NewDig(line)
		switch dig.direction {
		case "R":
			x += dig.distance
		case "L":
			x -= dig.distance
		case "D":
			y += dig.distance
		default:
			y -= dig.distance
		}
		maxWidth = max(maxWidth, x)
		maxHeight = max(maxHeight, y)
		minWidth = min(minWidth, x)
		minHeight = min(minHeight, y)
		digPlan = append(digPlan, NewDig(line))
	}
	grid.Width = maxWidth - minWidth + 1
	grid.Height = maxHeight - minHeight + 1
	grid.StartingY = minHeight * -1
	grid.StartingX = minWidth * -1
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
	grid.display()
	return answer
}
