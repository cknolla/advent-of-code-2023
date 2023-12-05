package day5

import (
	"advent-of-code-2023/utils"
	"fmt"
	"strconv"
	"strings"
)

type seedRange struct {
	start int
	end   int
}

func (m *mapping) convertPart2(input *seedRange) ([]seedRange, bool) {
	var outputRanges []seedRange
	if input.start >= m.source && input.start <= m.source+m.range_ {
		if input.end <= m.source+m.range_ {
			// input fully enclosed in mapping range
			outputRanges = append(outputRanges, seedRange{
				start: m.destination + (input.start - m.source),
				end:   m.destination + (input.end - m.source),
			})
		} else {
			// input goes beyond mapping, breaks into two
			mappingEnd := m.source + m.range_
			outputRanges = append(outputRanges, seedRange{
				start: m.destination + (input.start - m.source),
				end:   m.destination + (mappingEnd - m.source),
			})
			outputRanges = append(outputRanges, seedRange{
				start: mappingEnd + 1,
				end:   input.end,
			})
		}
		return outputRanges, true
	}
	// input not found in mapping
	outputRanges = append(outputRanges, *input)
	return outputRanges, false
}

func processSeedsPart2(line string) []seedRange {
	segments := strings.Split(line, ":")
	seedStrs := strings.Fields(segments[1])
	seedRanges := make([]seedRange, len(seedStrs)/2)
	for i := 0; i < len(seedStrs); i += 2 {
		start, _ := strconv.Atoi(seedStrs[i])
		range_, _ := strconv.Atoi(seedStrs[i+1])
		seedRanges[i/2] = seedRange{
			start: start,
			end:   start + range_,
		}
	}
	return seedRanges
}

func processMappingsPart2(mappings []mapping, inputs []seedRange) []seedRange {
	var outputs []seedRange
	var found bool
	for i, input := range inputs {
		for _, mapping := range mappings {
			var output []seedRange
			// weakness: seed ranges can be broken into two parts after conversion, but
			// the out-of-mapping range doesn't get rechecked against other mappings.
			// Have to assume no seed ranges split across multiple mappings
			output, found = mapping.convertPart2(&input)
			if found {
				for _, sr := range output {
					outputs = append(outputs, sr)
				}
				break
			}
		}
		if !found {
			outputs[i] = input
		}
	}
	return outputs
}

func getMinLocation(seeds []seedRange) int {
	minLoc := 999999999999999
	for _, seed := range seeds {
		if seed.start < minLoc {
			minLoc = seed.start
		}
	}
	return minLoc
}

func parseFilePart2() int {
	var seeds []seedRange
	var mappings []mapping
	for line := range utils.GetInputLines() {
		if strings.HasPrefix(line, "seeds") {
			seeds = processSeedsPart2(line)
		} else if strings.Contains(line, "map") {
			if mappings != nil {
				seeds = processMappingsPart2(mappings, seeds)
			}
			mappings = nil
		} else if line != "" {
			mappings = append(mappings, newMapping(line))
		}
	}
	seeds = processMappingsPart2(mappings, seeds)
	minLoc := getMinLocation(seeds)
	return minLoc
}

func Part2() (string, error) {
	nearestLoc := parseFilePart2()
	return fmt.Sprintf("%d", nearestLoc), nil
}
