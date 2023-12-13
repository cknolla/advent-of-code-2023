package day13

import (
	"advent-of-code-2023/utils"
	"math/bits"
)

type pattern struct {
	rows    []uint64
	columns []uint64
	width   int
}

func (p *pattern) reverseColumns() {
	for i, column := range p.columns {
		reversed := bits.Reverse64(column)
		p.columns[i] = reversed >> (64 - len(p.rows))
	}
}

func (p *pattern) summarize() (summary int) {
	summary += 100 * summarize(p.rows, -1)
	summary += summarize(p.columns, -1)
	return summary
}

func summarize(rows []uint64, excludedMirrorPoint int) int {
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
		if mirrorPoint == -1 && left != excludedMirrorPoint-1 && rows[left] == rows[right] {
			mirrorPoint = left
		}
	}
	return mirrorPoint + 1
}

func binaryConvert(line string) (value uint64) {
	for i, char := range line {
		if char == '#' {
			value |= 1 << (len(line) - 1 - i)
		}
	}
	return value
}

func parseFile(filename string) []pattern {
	var patterns []pattern
	activePattern := pattern{}
	for l := range utils.GetInputLines(filename) {
		line := l.Text
		if line == "" {
			activePattern.reverseColumns()
			patterns = append(patterns, activePattern)
			activePattern = pattern{}
			continue
		}
		activePattern.width = len(line)
		activePattern.rows = append(activePattern.rows, binaryConvert(line))
		if activePattern.columns == nil {
			activePattern.columns = make([]uint64, len(line))
		}
		for i := 0; i < len(line); i++ {
			if line[i] == '#' {
				// ideally this would be height - len(activePattern.rows) - 1, but height is unknown
				// so build columns in reverse order, and reverse them once fully built
				activePattern.columns[i] |= 1 << (len(activePattern.rows) - 1)
			}
		}
	}
	activePattern.reverseColumns()
	patterns = append(patterns, activePattern)
	return patterns
}

func Part1(filename string) int {
	answer := 0
	patterns := parseFile(filename)
	for _, pattern := range patterns {
		answer += pattern.summarize()
	}
	return answer
}
