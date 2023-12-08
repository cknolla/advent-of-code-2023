package day8

import (
	"advent-of-code-2023/utils"
	"fmt"
)

type move struct {
	left  string
	right string
}

func navigate(directions string, moves map[string]move) (steps int) {
	steps = 0
	currentLoc := "AAA"
OUTER:
	for {
		for _, char := range directions {
			steps++
			m := moves[currentLoc]
			if char == 'L' {
				currentLoc = m.left
			} else {
				currentLoc = m.right
			}
			if currentLoc == "ZZZ" {
				break OUTER
			}
		}
	}
	return steps
}

// conveniently, all lines are the exact same length, so I can just use indexing
func parseMove(line string) (location string, m move) {
	location = line[0:3]
	m.left = line[7:10]
	m.right = line[12:15]
	return location, m
}

func parseFile(filename string) (directions string, moves map[string]move) {
	moves = make(map[string]move, 750)
	for l := range utils.GetInputLines(filename) {
		if l.Index > 0 {
			line := l.Text
			if line == "" {
				continue
			}
			key, value := parseMove(line)
			moves[key] = value
		} else {
			directions = l.Text
		}
	}
	return directions, moves
}

func Part1() (string, error) {
	answer := 0
	directions, moves := parseFile("input.txt")
	answer = navigate(directions, moves)
	return fmt.Sprintf("%d", answer), nil
}
