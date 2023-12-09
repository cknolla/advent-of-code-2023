package day9

import (
	"advent-of-code-2023/utils"
	"fmt"
	"strconv"
	"strings"
)

type history struct {
	values [][]int
}

func (h *history) getNextValue() int {
	if len(h.values) == 1 {
		h.buildDifferences(h.values[0])
	}
	diff := 0
	for i := len(h.values) - 2; i >= 0; i-- {
		diff = diff + h.values[i][len(h.values[i])-1]
	}
	return diff
}

func (h *history) buildDifferences(values []int) {
	difference := make([]int, len(values)-1)
	nonZero := false
	for i := 0; i < len(values)-1; i++ {
		difference[i] = values[i+1] - values[i]
		if difference[i] != 0 {
			nonZero = true
		}
	}
	h.values = append(h.values, difference)
	if nonZero {
		h.buildDifferences(difference)
	}
}

func parseFile(filename string) []history {
	histories := make([]history, 200)
	for l := range utils.GetInputLines(filename) {
		line := l.Text
		if line == "" {
			continue
		}
		strValues := strings.Fields(line)
		histories[l.Index].values = append(histories[l.Index].values, make([]int, len(strValues)))
		for i, strValue := range strValues {
			intValue, _ := strconv.Atoi(strValue)
			histories[l.Index].values[0][i] = intValue
		}
	}
	return histories
}

func Part1() (string, error) {
	answer := 0
	histories := parseFile("input.txt")
	for _, history := range histories {
		answer += history.getNextValue()
	}
	return fmt.Sprintf("%d", answer), nil
}
