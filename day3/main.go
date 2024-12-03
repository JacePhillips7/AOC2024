package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	file := readfile()
	re := regexp.MustCompile(`(?m)(mul\(+\d+,\d+\))`)
	total := 0
	for _, match := range re.FindAllString(file, -1) {
		total += runMul(match)
	}
	fmt.Println(total)

}
func readfile() string {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	read := string(file)
	return read
}
func runMul(s string) int {
	re := regexp.MustCompile(`(?m)\d+,\d+`)
	string_numbs := re.FindAllString(s, -1)
	if len(string_numbs) != 1 {
		panic("why is there more than 1 group?")
	}
	numbs := strings.Split(string_numbs[0], ",")
	num1, _ := strconv.Atoi(numbs[0])
	num2, _ := strconv.Atoi(numbs[1])

	return num1 * num2
}
