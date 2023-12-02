package day1

import (
	"advent-of-code-2023/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func getPart1CalibrationValue(inputLine string) int {
	re := regexp.MustCompile("[0-9]")
	numbers := re.FindAllString(inputLine, -1)
	calibrationValue, err := strconv.Atoi(numbers[0] + numbers[len(numbers)-1])
	if err != nil {
		log.Fatalln(err)
	}
	return calibrationValue
}

// Part1 returns the challenge answer as a string (to be generic in case future puzzles
// have non-integer values) and an error if any.
// This particular function can't possibly return an error, but I wanted to create a
// standard function signature in case I build a project-level function that can call
// every Part function from every day.
func Part1() (string, error) {
	sum := 0
	for line := range utils.GetInputLines() {
		sum += getPart1CalibrationValue(line)
	}
	return fmt.Sprintf("%d", sum), nil
}
