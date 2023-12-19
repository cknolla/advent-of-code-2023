package day19

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFile(t *testing.T) {
	workflows, ratings := parseFile("sample.txt")
	assert.Equal(t, Workflow{
		name: "lnx",
		rules: []Rule{
			{
				property:    'm',
				operator:    '>',
				threshold:   1548,
				destination: "A",
			},
			{
				destination: "A",
			},
		},
	}, workflows["lnx"])
	assert.Equal(t, Workflow{
		name: "qqz",
		rules: []Rule{
			{
				property:    's',
				operator:    '>',
				threshold:   2770,
				destination: "qs",
			},
			{
				property:    'm',
				operator:    '<',
				threshold:   1801,
				destination: "hdj",
			},
			{
				destination: "R",
			},
		},
	}, workflows["qqz"])
	assert.Equal(t, Rating{787, 2655, 1222, 2876, false}, ratings[0])
}

func TestPart1Sample(t *testing.T) {
	answer := Part1("sample.txt")
	assert.Equal(t, 19114, answer)
}

func TestPart1(t *testing.T) {
	answer := Part1("input.txt")
	assert.Equal(t, 509597, answer)
}
