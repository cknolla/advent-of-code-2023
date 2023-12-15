package day15

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStore(t *testing.T) {
	steps := parseFile("sample.txt")
	processStep(steps[0])
	assert.Equal(t, 1, len(boxes[0].lenses))
	assert.Equal(t, 0, len(boxes[1].lenses))
}

func TestPart2Sample(t *testing.T) {
	answer := Part2("sample.txt")
	assert.Equal(t, 2, len(boxes[0].lenses))
	assert.Equal(t, lens{"rn", 0, 1}, boxes[0].lenses[0])
	assert.Equal(t, lens{"cm", 0, 2}, boxes[0].lenses[1])
	assert.Equal(t, 0, len(boxes[1].lenses))
	assert.Equal(t, 0, len(boxes[2].lenses))
	assert.Equal(t, 3, len(boxes[3].lenses))
	assert.Equal(t, lens{"ot", 3, 7}, boxes[3].lenses[0])
	assert.Equal(t, lens{"ab", 3, 5}, boxes[3].lenses[1])
	assert.Equal(t, lens{"pc", 3, 6}, boxes[3].lenses[2])
	assert.Equal(t, 0, len(boxes[4].lenses))
	assert.Equal(t, 145, answer)
}

func TestPart2(t *testing.T) {
	answer := Part2("input.txt")
	assert.Equal(t, 247763, answer)
}
