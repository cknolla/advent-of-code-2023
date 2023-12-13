package day12

import (
	"advent-of-code-2023/utils"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

type row struct {
	springs  []rune
	criteria []int
}

func (r *row) matchesCriteria() bool {
	contiguousSprings := 0
	var springCounts []int
	for _, spring := range r.springs {
		if spring == '.' {
			if contiguousSprings > 0 {
				springCounts = append(springCounts, contiguousSprings)
			}
			contiguousSprings = 0
		} else if spring == '#' {
			contiguousSprings++
		}
	}
	if contiguousSprings > 0 {
		springCounts = append(springCounts, contiguousSprings)
	}
	if slices.Equal(springCounts, r.criteria) {
		return true
	}
	return false
}

func (r *row) testAll() int {
	var indexes []int
	for index, spring := range r.springs {
		if spring == '?' {
			indexes = append(indexes, index)
		}
	}
	unknownCount := len(indexes)
	validArrangements := 0
	for i := 0; i < int(math.Pow(float64(2), float64(unknownCount))); i++ {
		for j := 0; j < unknownCount; j++ {
			if i&(1<<j) == 0 {
				r.springs[indexes[j]] = '.'
			} else {
				r.springs[indexes[j]] = '#'
			}
		}
		if r.matchesCriteria() {
			validArrangements++
		}
	}
	return validArrangements
}

func parseFile(filename string) []row {
	rows := make([]row, 1000)
	for l := range utils.GetInputLines(filename) {
		line := l.Text
		if line == "" {
			continue
		}
		segments := strings.Fields(l.Text)
		rows[l.Index].springs = []rune(segments[0])
		criteraStrs := strings.Split(segments[1], ",")
		for _, c := range criteraStrs {
			criteriaInt, _ := strconv.Atoi(c)
			rows[l.Index].criteria = append(rows[l.Index].criteria, criteriaInt)
		}
	}
	return rows
}

func Part1() (string, error) {
	sum := 0
	rows := parseFile("input.txt")
	for _, row := range rows {
		validArrangements := row.testAll()
		sum += validArrangements
	}
	return fmt.Sprintf("%d", sum), nil
}
