package day12

import (
	"advent-of-code-2023/utils"
	"fmt"
	"strconv"
	"strings"
)

var cache = make(map[string]int)

type row2 struct {
	springs  string
	criteria []int
}

// ended up copying this. wasn't gonna figure it out
// https://raw.githubusercontent.com/Oupsman/AOC2023/main/d12/advent_12.go
func recursiveArrangements(springs string, criterias []int) int {
	key := springs
	for _, criteria := range criterias {
		key += strconv.Itoa(criteria) + ","
	}
	if v, ok := cache[key]; ok {
		return v
	}
	if len(springs) == 0 {
		if len(criterias) == 0 {
			return 1
		} else {
			return 0
		}
	}
	if strings.HasPrefix(springs, "?") {
		return recursiveArrangements(strings.Replace(springs, "?", ".", 1), criterias) +
			recursiveArrangements(strings.Replace(springs, "?", "#", 1), criterias)
	}
	if strings.HasPrefix(springs, ".") {
		res := recursiveArrangements(strings.TrimPrefix(springs, "."), criterias)
		cache[key] = res
		return res
	}

	if strings.HasPrefix(springs, "#") {
		if len(criterias) == 0 {
			cache[key] = 0
			return 0
		}
		if len(springs) < criterias[0] {
			cache[key] = 0
			return 0
		}
		if strings.Contains(springs[0:criterias[0]], ".") {
			cache[key] = 0
			return 0
		}
		if len(criterias) > 1 {
			if len(springs) < criterias[0]+1 || string(springs[criterias[0]]) == "#" {
				cache[key] = 0
				return 0
			}
			res := recursiveArrangements(springs[criterias[0]+1:], criterias[1:])
			cache[key] = res
			return res
		} else {
			res := recursiveArrangements(springs[criterias[0]:], criterias[1:])
			cache[key] = res
			return res
		}
	}

	return 0
}

func parseFilePart2(filename string) []row2 {
	rows := make([]row2, 1000)
	for l := range utils.GetInputLines(filename) {
		line := l.Text
		if line == "" {
			continue
		}
		segments := strings.Fields(l.Text)
		rows[l.Index].springs = segments[0] + "?" + segments[0] + "?" + segments[0] + "?" + segments[0] + "?" + segments[0]
		criteriaStrs := strings.Split(segments[1], ",")
		for i := 0; i < 5; i++ {
			for _, c := range criteriaStrs {
				criteriaInt, _ := strconv.Atoi(c)
				rows[l.Index].criteria = append(rows[l.Index].criteria, criteriaInt)
			}
		}
	}
	return rows
}

func Part2() (string, error) {
	sum := 0
	rows := parseFilePart2("input.txt")
	for _, row := range rows {
		validArrangements := recursiveArrangements(row.springs, row.criteria)
		sum += validArrangements
	}
	return fmt.Sprintf("%d", sum), nil
}
