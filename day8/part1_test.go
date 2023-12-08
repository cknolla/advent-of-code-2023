package day8

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestParseFile(t *testing.T) {
	directions, moves := parseFile("input.txt")
	assert.Equal(t, "LRLRRLRLRRRLRRRLRRLRLLRLRLRRRLRLRRLLRRLLRRRLLRRRLRRRLRRLLRLRRRLRRLRLRLLRRLLRRRLLRLRRRLRRRLLRLRRRLLRLLRRLRLRRRLLRLRLLRRRLLRLRRRLLLRRRLLLRRLLLRRRLLRLRLRLRRLLRRRLRRLRRRLRRLRRRLRLRRLRLRRRLRLRRRLRRLRRRLRLLLRLRRRLRLLRLRRLRRRLRRLRLRLRLRRLRRLLRLLLRLRLRRRLRRRLLRLLRLRRLRRRLRRLRRRLRLRRRR", directions)
	assert.Equal(t, move{"KFV", "CFQ"}, moves["JKT"])
	assert.Equal(t, move{"JFX", "BBF"}, moves["CBL"])
	assert.Equal(t, move{"BQX", "PMT"}, moves["BGF"])
	assert.Equal(t, move{"TSK", "ZZZ"}, moves["PLJ"])
	for key, value := range moves {
		assert.Equal(t, 3, len(key))
		assert.False(t, strings.Contains(key, "("))
		assert.False(t, strings.Contains(key, ")"))
		assert.False(t, strings.Contains(key, ","))
		assert.Equal(t, 3, len(value.left))
		assert.False(t, strings.Contains(value.left, "("))
		assert.False(t, strings.Contains(value.left, ")"))
		assert.False(t, strings.Contains(value.left, ","))
		assert.Equal(t, 3, len(value.right))
		assert.False(t, strings.Contains(value.right, "("))
		assert.False(t, strings.Contains(value.right, ")"))
		assert.False(t, strings.Contains(value.right, ","))
	}
}

func TestSample(t *testing.T) {
	directions, moves := parseFile("sample.txt")
	steps := navigate(directions, moves)
	assert.Equal(t, 2, steps)
}

func TestSample2(t *testing.T) {
	directions, moves := parseFile("sample2.txt")
	steps := navigate(directions, moves)
	assert.Equal(t, 6, steps)
}

func TestPart1(t *testing.T) {
	answer, err := Part1()
	assert.Nil(t, err)
	assert.Equal(t, "20221", answer)
}
