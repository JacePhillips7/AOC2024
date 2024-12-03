package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	reports := getReports()
	total_safe := 0
	for _, r := range reports {
		line := strings.Split(r, " ")
		safe := isSafe(line)
		if safe {
			total_safe++
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
func isSafe(line []string) bool {
	allowed_safe := []int{1, 2, 3}

	my_ints := []int{}
	for _, v := range line {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		my_ints = append(my_ints, i)
	}
	//mode_increase to see if we are increasing or decreasing
	mode_increase := true //increasing
	if my_ints[0] > my_ints[1] {
		mode_increase = false //now set mode to decrease
	}
	//now we can see if it is safe
	running_value := my_ints[0]
	for i := 1; i < len(my_ints); i++ {
		check := getAbs(running_value - my_ints[i])
		safe := slices.Contains(allowed_safe, check)
		if !safe {
			return false
		}
		if mode_increase && running_value > my_ints[i] {
			return false
		}
		if !mode_increase && running_value < my_ints[i] {
			return false
		}
		running_value = my_ints[i]
	}
	return true
}

func getAbs(v int) int {
	return int(math.Abs(float64(v)))
}
