package day6

import (
	"advent-of-code-2023/utils"
	"fmt"
	"strconv"
	"strings"
)

func parseFilePart2() race {
	var r race
	for line := range utils.GetInputLines() {
		parts := strings.Split(line, ":")
		if parts[0] == "Time" {
			fields := strings.Fields(parts[1])
			concatTime := strings.Join(fields, "")
			t, _ := strconv.Atoi(concatTime)
			r.time = t
		} else {
			fields := strings.Fields(parts[1])
			concatDistance := strings.Join(fields, "")
			d, _ := strconv.Atoi(concatDistance)
			r.recordDistance = d
		}
	}
	return r
}

func Part2() (string, error) {
	answer := 0
	r := parseFilePart2()
	answer = r.getWinningTimes()
	return fmt.Sprintf("%d", answer), nil
}
