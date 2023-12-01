package utils

import (
	"bufio"
	"os"
)

// GetInputLines returns 1 line at a time when looped over
func GetInputLines() (channel chan string) {
	channel = make(chan string, 1)
	go func() {
		file, err := os.Open("input.txt")
		if err != nil {
			close(channel)
			return
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for {
			if scanner.Scan() {
				channel <- scanner.Text()
			} else {
				close(channel)
				return
			}
		}

	}()
	return channel
}
