package day6

import (
	"advent-of-code-2023/utils"
	"fmt"
	"strconv"
	"strings"
)

type race struct {
	time           int
	recordDistance int
}

func (r *race) getWinningTimes() int {
	var winningTimes int
	for t := 1; t < r.time; t++ {
		remainingTime := r.time - t
		distance := t * remainingTime
		if distance > r.recordDistance {
			winningTimes++
		}
	}
	return winningTimes
}

func parseFile() []race {
	var races []race
	for l := range utils.GetInputLines("input.txt") {
		line := l.Text
		parts := strings.Split(line, ":")
		if parts[0] == "Time" {
			fields := strings.Fields(parts[1])
			races = make([]race, len(fields))
			for i, field := range fields {
				t, _ := strconv.Atoi(field)
				races[i].time = t
			}
		} else {
			fields := strings.Fields(parts[1])
			for i, field := range fields {
				d, _ := strconv.Atoi(field)
				races[i].recordDistance = d
			}
		}
	}
	return races
}

func Part1() (string, error) {
	answer := 0
	races := parseFile()
	for _, r := range races {
		times := r.getWinningTimes()
		if answer == 0 {
			answer = times
		} else {
			answer *= times
		}
	}
	return fmt.Sprintf("%d", answer), nil
}
