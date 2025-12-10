package utils

import (
	"bufio"
	"os"
)

func GetFileLines(path string) (<-chan string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	ch := make(chan string)

	f := func() {
		defer file.Close()
		defer close(ch)

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			ch <- scanner.Text()
		}

	}

	go f()

	return ch, nil
}
