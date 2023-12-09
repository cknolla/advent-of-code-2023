package day9

import "fmt"

func (h *history) getPrevValue() int {
	if len(h.values) == 1 {
		h.buildDifferences(h.values[0])
	}
	diff := 0
	for i := len(h.values) - 2; i >= 0; i-- {
		diff = h.values[i][0] - diff
	}
	return diff
}

func Part2() (string, error) {
	answer := 0
	histories := parseFile("input.txt")
	for _, history := range histories {
		answer += history.getPrevValue()
	}
	return fmt.Sprintf("%d", answer), nil
}
