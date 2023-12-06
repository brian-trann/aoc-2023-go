package utils

import (
	"bufio"
	"log"
	"os"

	"strconv"
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
func SumStringNumbers(strs []string) (int, error) {
	var sum int
	for _, str := range strs {
		num, err := strconv.Atoi(str)
		if err != nil {
			return 0, err
		}
		sum += num
	}
	return sum, nil
}
func OpenFileTo2dArray(fileName string) [][]rune {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln((err))
	}
	defer file.Close()
	var array [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		array = append(array, row)
	}
	return array
}
