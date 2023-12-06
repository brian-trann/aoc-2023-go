package utils

import (
	"bufio"
	"log"
	"os"
	"strings"

	"strconv"
)

func OpenFileSliceStrings(fileName string) []string {
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

func OpenFileAsString(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var builder strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		builder.WriteString(scanner.Text() + "\n") // Add a newline character after each line
	}

	return builder.String()
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
func ConvertStringToIntSlice(input string, delimiter string) []int {
	var result []int
	substrings := strings.Split(input, delimiter)
	for _, s := range substrings {
		num, err := strconv.Atoi(s)
		if err != nil {

			// fmt.Println("error parsing strign to int")
			// continue
		} else {

			result = append(result, num)
		}
	}
	return result
}
