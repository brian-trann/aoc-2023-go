package utils

import (
	"bufio"
	"log"
	"os"
)

func OpenFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln((err))
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
