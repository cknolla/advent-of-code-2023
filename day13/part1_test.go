package day13

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFile(t *testing.T) {
	patterns := parseFile("sample.txt")
	assert.Equal(t, 358, patterns[0].rows[0])
	assert.Equal(t, 90, patterns[0].rows[1])
	assert.Equal(t, 89, patterns[0].columns[0])
	assert.Equal(t, 24, patterns[0].columns[1])

	assert.Equal(t, 281, patterns[1].rows[0])
	assert.Equal(t, 265, patterns[1].rows[6])
}

func TestPart1Sample(t *testing.T) {
	answer := Part1("sample.txt")
	assert.Equal(t, 405, answer)
}

func TestPart1(t *testing.T) {
	answer := Part1("input.txt")
	assert.Equal(t, 37718, answer)
}
