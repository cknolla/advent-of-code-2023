package day12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFile(t *testing.T) {
	rows := parseFile("input.txt")
	assert.Equal(t, []rune{'?', '?', '?', '?', '?', '?', '#', '?', '?', '#', '?', '?'}, rows[0].springs)
	assert.Equal(t, []int{1, 1, 5, 1}, rows[0].criteria)
	assert.Equal(t, []rune("??.?????#?"), rows[999].springs)
	assert.Equal(t, []int{1, 5}, rows[999].criteria)
}

func TestMatchesCriteria(t *testing.T) {
	r := row{[]rune("?###????????"), []int{3, 2, 1}}
	r.springs = []rune(".###.##.#...")
	assert.True(t, r.matchesCriteria())
	r.springs = []rune(".###.##....#")
	assert.True(t, r.matchesCriteria())
	r.springs = []rune(".######.#...")
	assert.False(t, r.matchesCriteria())
	r.springs = []rune(".###.##.#.#.")
	assert.False(t, r.matchesCriteria())
	r.springs = []rune(".###.##.##..")
	assert.False(t, r.matchesCriteria())
}

func TestTestAll(t *testing.T) {
	rows := parseFile("sample.txt")
	sum := 0
	for i := 0; i < 6; i++ {
		validArrangements := rows[i].testAll()
		sum += validArrangements
	}
	assert.Equal(t, 21, sum)
	r := row{[]rune("?###??????????###????????"), []int{3, 2, 1, 3, 2, 1}}
	arrangements := r.testAll()
	assert.Equal(t, 150, arrangements)
}

func TestPart1(t *testing.T) {
	answer, err := Part1()
	assert.Nil(t, err)
	assert.Equal(t, "6488", answer)
}
