package day5

import (
	"advent-of-code-2023/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type mapping struct {
	destination int
	source      int
	range_      int
}

func newMapping(line string) mapping {
	fields := strings.Fields(line)
	destination, _ := strconv.Atoi(fields[0])
	source, _ := strconv.Atoi(fields[1])
	range_, _ := strconv.Atoi(fields[2])
	return mapping{
		destination: destination,
		source:      source,
		range_:      range_,
	}
}

func (m *mapping) convert(input int) (int, bool) {
	if input >= m.source && input <= m.source+m.range_ {
		return m.destination + (input - m.source), true
	}
	return input, false
}

func convertStringSliceToInt(slice []string) []int {
	output := make([]int, len(slice))
	for index, str := range slice {
		output[index], _ = strconv.Atoi(str)
	}
	return output
}

func processSeeds(line string) []int {
	segments := strings.Split(line, ":")
	seedStrs := strings.Fields(segments[1])
	return convertStringSliceToInt(seedStrs)
}

func processMappings(mappings []mapping, inputs []int) []int {
	outputs := make([]int, len(inputs))
	var found bool
	for i, input := range inputs {
		for _, mapping := range mappings {
			var output int
			output, found = mapping.convert(input)
			if found {
				outputs[i] = output
				break
			}
		}
		if !found {
			outputs[i] = input
		}
	}
	return outputs
}

func parseFile() int {
	var seeds []int
	var mappings []mapping
	for line := range utils.GetInputLines() {
		if strings.HasPrefix(line, "seeds") {
			seeds = processSeeds(line)
		} else if strings.Contains(line, "map") {
			if mappings != nil {
				seeds = processMappings(mappings, seeds)
			}
			mappings = nil
		} else if line != "" {
			mappings = append(mappings, newMapping(line))
		}
	}
	seeds = processMappings(mappings, seeds)
	return slices.Min(seeds)
}

func Part1() (string, error) {
	nearestLoc := parseFile()
	return fmt.Sprintf("%d", nearestLoc), nil
}
