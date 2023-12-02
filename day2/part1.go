package day2

import (
	"advent-of-code-2023/utils"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type round struct {
	red   int
	green int
	blue  int
}

func NewRound() round {
	return round{
		red:   0,
		green: 0,
		blue:  0,
	}
}

type game struct {
	id     int
	rounds []round
}

func parseGame(line string) game {
	var err error
	g := game{}
	majorParts := strings.Split(line, ":")
	gameParts := strings.Split(majorParts[0], " ")
	g.id, err = strconv.Atoi(gameParts[1])
	if err != nil {
		log.Fatalln(err)
	}
	roundParts := strings.Split(majorParts[1], ";")
	for _, roundPart := range roundParts {
		r := NewRound()
		cubes := strings.Split(roundPart, ",")
		for _, cube := range cubes {
			cubeParts := strings.Split(strings.Trim(cube, " "), " ")
			value, err := strconv.Atoi(cubeParts[0])
			if err != nil {
				log.Fatalln(err)
			}
			switch cubeParts[1] {
			case "red":
				r.red = value
			case "green":
				r.green = value
			case "blue":
				r.blue = value
			}
		}
		g.rounds = append(g.rounds, r)
	}
	return g
}

func (g *game) roundMaximum() round {
	maxCubes := round{
		red:   0,
		green: 0,
		blue:  0,
	}
	for _, r := range g.rounds {
		maxCubes.red = int(math.Max(float64(maxCubes.red), float64(r.red)))
		maxCubes.green = int(math.Max(float64(maxCubes.green), float64(r.green)))
		maxCubes.blue = int(math.Max(float64(maxCubes.blue), float64(r.blue)))
	}
	return maxCubes
}

func (g *game) isPossible(limit *round) bool {
	roundMax := g.roundMaximum()
	return limit.red >= roundMax.red && limit.green >= roundMax.green && limit.blue >= roundMax.blue
}

func Part1() (string, error) {
	sum := 0
	limit := round{
		red:   12,
		green: 13,
		blue:  14,
	}
	for line := range utils.GetInputLines() {
		g := parseGame(line)
		if g.isPossible(&limit) {
			sum += g.id
		}
	}
	return fmt.Sprintf("%d", sum), nil
}
