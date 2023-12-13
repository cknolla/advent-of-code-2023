package day13

func (p *pattern) fixSmudge() int {
	originalRowSummary := summarize(p.rows, -1)
	originalColumnSummary := summarize(p.columns, -1)
	for rowIndex := range p.rows {
		for columnIndex := range p.columns {
			p.rows[rowIndex] ^= 1 << (len(p.columns) - 1 - columnIndex)
			p.columns[columnIndex] ^= 1 << (len(p.rows) - 1 - rowIndex)
			newRowSummary := summarize(p.rows, originalRowSummary)
			if newRowSummary != 0 {
				return newRowSummary * 100
			}
			newColumnSummary := summarize(p.columns, originalColumnSummary)
			if newColumnSummary != 0 {
				return newColumnSummary
			}
			p.rows[rowIndex] ^= 1 << (len(p.columns) - 1 - columnIndex)
			p.columns[columnIndex] ^= 1 << (len(p.rows) - 1 - rowIndex)
		}
	}
	return 0
}

func Part2(filename string) int {
	answer := 0
	patterns := parseFile(filename)
	for _, pattern := range patterns {
		answer += pattern.fixSmudge()
	}
	return answer
}
