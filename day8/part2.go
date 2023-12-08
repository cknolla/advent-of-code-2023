package day8

import (
	"fmt"
	"slices"
)

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func navigatePart2(directions string, moves map[string]move) (steps int) {
	steps = 0
	var aLocs []move
	for key, m := range moves {
		if key[2] == 'A' {
			aLocs = append(aLocs, m)
		}
	}
	zDistances := make([]int, len(aLocs))
OUTER:
	for {
		for _, char := range directions {
			steps++
			for i, loc := range aLocs {
				if char == 'L' {
					if zDistances[i] == 0 && loc.left[2] == 'Z' {
						zDistances[i] = steps
					}
					aLocs[i] = moves[loc.left]
				} else {
					if zDistances[i] == 0 && loc.right[2] == 'Z' {
						zDistances[i] = steps
					}
					aLocs[i] = moves[loc.right]
				}
			}
			if !slices.Contains(zDistances, 0) {
				break OUTER
			}
		}
	}
	lcm := LCM(zDistances[0], zDistances[1], zDistances[2:]...)
	return lcm
}

func Part2() (string, error) {
	answer := 0
	directions, moves := parseFile("input.txt")
	answer = navigatePart2(directions, moves)
	return fmt.Sprintf("%d", answer), nil
}
