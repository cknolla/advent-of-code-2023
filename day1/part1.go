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

func Part1() (string, error) {
	lines := utils.GetInputLines()
	sum := 0
	for _, line := range lines {
		sum += getPart1CalibrationValue(line)
	}
	return fmt.Sprintf("%d", sum), nil
}
