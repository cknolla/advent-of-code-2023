package utils

import (
	"bufio"
	"os"
)

func GetInputLines() (channel chan string) {
	channel = make(chan string, 1)
	buffer := ""
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
				buffer = scanner.Text()
				channel <- buffer
				buffer = ""
			} else {
				close(channel)
				return
			}
		}

	}()
	return channel
}
