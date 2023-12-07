package day2

import (
	"advent-of-code-2023/utils"
	"fmt"
)

func (g *game) getPower() int {
	return g.maxRound.red * g.maxRound.green * g.maxRound.blue
}

func Part2() (string, error) {
	sum := 0
	for l := range utils.GetInputLines("input.txt") {
		line := l.Text
		g := parseGame(line)
		sum += g.getPower()
	}
	return fmt.Sprintf("%d", sum), nil
}
