package day15

import (
	"advent-of-code-2023/utils"
	"strings"
)

func parseFile(filename string) (steps []string) {
	for l := range utils.GetInputLines(filename) {
		line := l.Text
		if line == "" {
			continue
		}
		steps = strings.Split(line, ",")
	}
	return steps
}

func hasher(step string) int32 {
	var h int32
	for _, c := range step {
		h += c
		h *= 17
		h %= 256
	}
	return h
}

func Part1(filename string) int {
	answer := 0
	steps := parseFile(filename)
	for _, step := range steps {
		answer += int(hasher(step))
	}
	return answer
}
