package day3

import (
	"advent-of-code-2023/utils"
	"fmt"
	"unicode"
)

func (s *schematic) getGearRatios() int {
	ratioSum := 0
	for _, gear := range s.symbolLocations {
		var linkedParts []*partNumber
		combinedPartNumbers := s.partNumbers
		if s.prev != nil {
			combinedPartNumbers = append(combinedPartNumbers, s.prev.partNumbers...)
		}
		if s.next != nil {
			combinedPartNumbers = append(combinedPartNumbers, s.next.partNumbers...)
		}
		for pnIndex, pn := range combinedPartNumbers {
			if pn.startLocation-1 <= gear && pn.endLocation+1 >= gear {
				linkedParts = append(linkedParts, &combinedPartNumbers[pnIndex])
			}
		}
		if linkedParts != nil && len(linkedParts) == 2 {
			ratioSum += linkedParts[0].value * linkedParts[1].value
		}
	}
	return ratioSum
}

func parsePart2Schematic(line string) schematic {
	s := schematic{}
	numStr := ""
	for index, char := range line {
		if unicode.IsNumber(char) {
			numStr += string(char)
			continue
		}
		// ignore non-* symbols
		if char == '*' {
			s.symbolLocations = append(s.symbolLocations, index)
		}
		if numStr != "" {
			s.addPartNumber(index, numStr)
			numStr = ""
		}
	}
	if numStr != "" {
		s.addPartNumber(len(line), numStr)
	}
	return s
}

func Part2() (string, error) {
	sum := 0
	var prevSchematic *schematic
	var firstSchematic *schematic
	for l := range utils.GetInputLines("input.txt") {
		line := l.Text
		s := parseSchematic(line)
		if firstSchematic == nil {
			firstSchematic = &s
		}
		if prevSchematic != nil {
			prevSchematic.next = &s
			s.prev = prevSchematic
		}
		prevSchematic = &s
	}
	for schem := firstSchematic; schem != nil; schem = schem.next {
		sum += schem.getGearRatios()
	}
	return fmt.Sprintf("%d", sum), nil
}
