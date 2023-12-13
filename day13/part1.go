package day13

import (
	"advent-of-code-2023/utils"
	"fmt"
)

type pattern struct {
	rows    []string
	columns []string
}

func (p *pattern) summarize() (summary int) {
	summary += 100 * summarize(p.rows)
	summary += summarize(p.columns)
	return summary
}

func summarize(rows []string) int {
	mirrorPoint := -1
	for left := 0; left < len(rows)-1; left++ {
		right := left + 1
		if mirrorPoint != -1 {
			mirrorIndex := mirrorPoint - (left - mirrorPoint)
			if mirrorIndex < 0 {
				// mirror complete
				break
			}
			if rows[right] != rows[mirrorIndex] {
				// not a true mirror. reset and keep looking
				mirrorPoint = -1
			}
		}
		if mirrorPoint == -1 && rows[left] == rows[right] {
			mirrorPoint = left
		}
	}
	return mirrorPoint + 1
}

func parseFile(filename string) []pattern {
	var patterns []pattern
	activePattern := pattern{}
	for l := range utils.GetInputLines(filename) {
		line := l.Text
		if line == "" {
			patterns = append(patterns, activePattern)
			activePattern = pattern{}
			continue
		}
		activePattern.rows = append(activePattern.rows, line)
		if activePattern.columns == nil {
			activePattern.columns = make([]string, len(line))
		}
		for i := 0; i < len(line); i++ {
			activePattern.columns[i] = activePattern.columns[i] + string(line[i])
		}
	}
	patterns = append(patterns, activePattern)
	return patterns
}

func Part1() (string, error) {
	answer := 0
	patterns := parseFile("input.txt")
	for _, pattern := range patterns {
		answer += pattern.summarize()
	}
	return fmt.Sprintf("%d", answer), nil
}
