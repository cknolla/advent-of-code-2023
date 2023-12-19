package utils

import (
	"bufio"
	"os"
	"strconv"
)

type Line struct {
	Index int
	Text  string
}

// GetInputLines returns 1 line at a time when looped over
func GetInputLines(filename string) (channel chan Line) {
	channel = make(chan Line, 1)
	currentLine := 0
	go func() {
		file, err := os.Open(filename)
		if err != nil {
			close(channel)
			return
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for {
			if scanner.Scan() {
				channel <- Line{currentLine, scanner.Text()}
				currentLine++
			} else {
				close(channel)
				return
			}
		}

	}()
	return channel
}

func ToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
