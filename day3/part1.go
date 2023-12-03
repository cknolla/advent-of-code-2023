package day3

import (
	"advent-of-code-2023/utils"
	"fmt"
	"strconv"
	"unicode"
)

// schematic represents a single line in the input
type schematic struct {
	prev            *schematic
	next            *schematic
	symbolLocations []int
	partNumbers     []partNumber
}

type partNumber struct {
	value         int
	valid         bool
	startLocation int
	endLocation   int
}

func (s *schematic) addPartNumber(index int, numStr string) {
	num, _ := strconv.Atoi(numStr)
	s.partNumbers = append(s.partNumbers, partNumber{
		value:         num,
		endLocation:   index - 1,
		startLocation: index - len(numStr),
	})
}

func (s *schematic) validatePartNumbers() int {
	validSum := 0
	combinedSymbolLocations := s.symbolLocations
	if s.prev != nil {
		combinedSymbolLocations = append(combinedSymbolLocations, s.prev.symbolLocations...)
	}
	if s.next != nil {
		combinedSymbolLocations = append(combinedSymbolLocations, s.next.symbolLocations...)
	}
	for index, pn := range s.partNumbers {
		s.partNumbers[index].valid = pn.isValid(combinedSymbolLocations)
		if s.partNumbers[index].valid {
			validSum += s.partNumbers[index].value
		}
	}
	return validSum
}

func parseSchematic(line string) schematic {
	s := schematic{}
	numStr := ""
	for index, char := range line {
		if unicode.IsNumber(char) {
			numStr += string(char)
			continue
		}
		if char != '.' {
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

func (p *partNumber) isValid(symbolLocations []int) bool {
	for _, loc := range symbolLocations {
		if p.startLocation-1 <= loc && p.endLocation+1 >= loc {
			return true
		}
	}
	return false
}

func Part1() (string, error) {
	sum := 0
	var prevSchematic *schematic
	var firstSchematic *schematic
	for line := range utils.GetInputLines() {
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
		sum += schem.validatePartNumbers()
	}
	return fmt.Sprintf("%d", sum), nil
}
