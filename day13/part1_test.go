package day13

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFile(t *testing.T) {
	patterns := parseFile("sample.txt")
	assert.Equal(t, "#.##..##.", patterns[0].rows[0])
	assert.Equal(t, "..#.##.#.", patterns[0].rows[1])
	assert.Equal(t, "#.##..#", patterns[0].columns[0])
	assert.Equal(t, "..##...", patterns[0].columns[1])

	assert.Equal(t, "#...##..#", patterns[1].rows[0])
	assert.Equal(t, "#....#..#", patterns[1].rows[6])
}

func TestPart1Sample(t *testing.T) {
	answer := 0
	patterns := parseFile("sample.txt")
	for _, pattern := range patterns {
		answer += pattern.summarize()
	}
	assert.Equal(t, 405, answer)
}

func TestPart1(t *testing.T) {
	answer, err := Part1()
	assert.Nil(t, err)
	assert.Equal(t, "37718", answer)
}
