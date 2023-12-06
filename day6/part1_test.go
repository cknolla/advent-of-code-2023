package day6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFile(t *testing.T) {
	races := parseFile()
	assert.Equal(t, []race{
		{55, 401},
		{99, 1485},
		{97, 2274},
		{93, 1405},
	},
		races,
	)
}

func TestPart1(t *testing.T) {
	answer, err := Part1()
	assert.Nil(t, err)
	assert.Equal(t, "2374848", answer)
}
