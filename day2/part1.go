package day2

import (
	"advent-of-code-2023/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type round struct {
	red   int
	green int
	blue  int
}

func newRound() round {
	return round{
		red:   0,
		green: 0,
		blue:  0,
	}
}

type game struct {
	id       int
	rounds   []round
	maxRound round
}

func parseGame(line string) game {
	var err error
	g := game{}
	majorParts := strings.Split(line, ":")
	gameParts := strings.Fields(majorParts[0])
	g.id, err = strconv.Atoi(gameParts[1])
	if err != nil {
		log.Fatalln(err)
	}
	roundParts := strings.Split(majorParts[1], ";")
	g.rounds = make([]round, len(roundParts))
	g.maxRound = newRound()
	for index, roundPart := range roundParts {
		r := newRound()
		cubes := strings.Split(roundPart, ",")
		for _, cube := range cubes {
			cubeParts := strings.Fields(cube)
			value, err := strconv.Atoi(cubeParts[0])
			if err != nil {
				log.Fatalln(err)
			}
			switch cubeParts[1] {
			case "red":
				r.red = value
				if r.red > g.maxRound.red {
					g.maxRound.red = r.red
				}
			case "green":
				r.green = value
				if r.green > g.maxRound.green {
					g.maxRound.green = r.green
				}
			case "blue":
				r.blue = value
				if r.blue > g.maxRound.blue {
					g.maxRound.blue = r.blue
				}
			}
		}
		g.rounds[index] = r
	}
	return g
}

func (g *game) isPossible(limit *round) bool {
	return limit.red >= g.maxRound.red && limit.green >= g.maxRound.green && limit.blue >= g.maxRound.blue
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
