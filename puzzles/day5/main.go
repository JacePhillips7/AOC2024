package main

import (
	read_file "aoc2024/library"
	"fmt"
	"strconv"
	"strings"
)

type Rules [][]string

type Data []string

func main() {
	file := read_file.ReadFile("./input.txt")
	rules, data := buildData(file)

	correctTotal := 0

	for _, d := range data {
		if checkRules(rules, d) {
			mid := getMiddleNumber(d)
			correctTotal += mid
		}
	}
	fmt.Println(correctTotal)
}

func buildData(file string) (Rules, Data) {
	lines := read_file.SplitOnLine(file)
	var rules Rules
	var data Data

	//build rules and slice
	for _, line := range lines {
		if strings.Contains(line, "|") {
			rules = append(rules, strings.Split(line, "|"))

		} else if strings.Contains(line, ",") {
			data = append(data, line)
		}
	}
	return rules, data
}

// sting to rule
// 45|32 -> [45,32]
func stringToIntArry(s string, d string) []int {
	split := strings.Split(s, d)
	var arr []int
	for _, v := range split {
		value, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		arr = append(arr, value)
	}
	return arr
}

func checkRules(r Rules, data string) bool {
	for _, rule := range r {
		index1 := strings.Index(data, (rule[0]))
		index2 := strings.Index(data, (rule[1]))

		//rule not found for line
		if index1 == -1 || index2 == -1 {
			continue
		}
		//if index 1 is not first, we have broken the rule
		if index1 > index2 {
			return false
		}
	}
	return true
}
func getMiddleNumber(line string) int {
	arr := stringToIntArry(line, ",")
	return arr[len(arr)/2]
}
