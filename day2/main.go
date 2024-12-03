package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports := getReports()
	total_safe := 0
	for _, r := range reports {
		line := strings.Split(r, " ")
		//part 1 here
		safe, _ := isSafe(line)
		if safe {
			total_safe++
			continue
		}
		//part 2 fuck it we brute force
		for i := 0; i < len(line); i++ {
			brute_line := RemoveIndex(line, i)
			brute_safe, _ := isSafe(brute_line)
			if brute_safe {
				total_safe++
				break
			}
		}
	}
	fmt.Println(total_safe)
}

func getReports() []string {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	read := string(file)
	read = strings.ReplaceAll(read, "\r\n", "\n")
	reports := strings.Split(read, "\n")
	return reports
}
func isSafe(line []string) (bool, int) {
	my_ints := []int{}
	for _, v := range line {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		my_ints = append(my_ints, i)
	}
	//is_increasing to see if we are increasing or decreasing
	is_increasing := true //increasing
	if my_ints[0] > my_ints[1] {
		is_increasing = false //now set mode to decrease
	}
	//now we can see if it is safe
	previous_value := my_ints[0]
	for i := 1; i < len(my_ints); i++ {
		delta := getAbs(previous_value - my_ints[i])
		if delta == 0 || delta > 3 || is_increasing && previous_value > my_ints[i] || !is_increasing && previous_value < my_ints[i] {
			return false, i
		}
		previous_value = my_ints[i]
	}
	return true, -1
}

func getAbs(v int) int {
	return int(math.Abs(float64(v)))
}
func RemoveIndex(s []string, i int) []string {
	newSlice := make([]string, len(s)-1)
	copy(newSlice, s[:i])
	copy(newSlice[i:], s[i+1:])
	return newSlice
}
