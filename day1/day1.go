package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

type InputScanner struct {
	*bufio.Scanner
}

func (s *InputScanner) scanLine() int {
	if s.Scan() {
		line := s.Text()
		re := regexp.MustCompile("[0-9]")
		numbers := re.FindAllString(line, -1)
		calibrationValue, err := strconv.Atoi(numbers[0] + numbers[len(numbers)-1])
		if err != nil {
			log.Fatalln(err)
		}
		return calibrationValue
	}
	return -1
}

func main() {
	scanner := InputScanner{bufio.NewScanner(os.Stdin)}
	sum := 0
	for calibration := scanner.scanLine(); calibration > 0; calibration = scanner.scanLine() {
		sum += calibration
	}
	log.Println(sum)
}
