package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file := readfile()
	lines := getLines(file)
	var full_array [][]string // => [[X,M,A,S],[X,M,A,S]]
	//
	totalCount := 0
	for _, v := range lines {
		line := strings.Split(v, "")
		full_array = append(full_array, line)
	}

	for y, y_line := range full_array {
		for x := range y_line {
			if checkSpace(full_array, "X", x, y) { //start by looking for the X
				//we found the X, now get the directions of M
				directions := searchFor(full_array, "M", x, y)
				for _, d := range directions {
					x_direction := d[0]
					y_direction := d[1]
					if matchString(full_array, "XMAS", x, y, x_direction, y_direction) {
						totalCount++
					}
				}

			}
		}
	}
	fmt.Println(totalCount)
}
func readfile() string {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	read := string(file)
	return read
}
func getLines(s string) []string {
	read := strings.ReplaceAll(s, "\r\n", "\n")
	return strings.Split(read, "\n")
}

// find letter and return the direction relative to the given input
// example, if we find X at 5,5
// and M is located at 6,5
// return 1,0
// showing x is up by 1, and y is 0
// spelling XMAS
// returns all possible directions to scan as [[x,y],[x,y]]
func searchFor(arr [][]string, search string, x int, y int) [][]int {
	combos := []int{-1, 0, 1} //where we are going to check for locations
	var directions [][]int
	for _, y_combo := range combos { //w checks the x
		for _, x_combo := range combos { //z checks the y
			check_x := x + x_combo
			check_y := y + y_combo
			//make sure check_y and check_x are not less than 0....
			if check_x < 0 || check_y < 0 {
				continue
			}
			//prevent outofbounds in the other direction now
			if check_y >= len(arr) || check_x >= len(arr[check_y]) {
				continue
			}
			if search == arr[check_y][check_x] {
				to_append := []int{x_combo, y_combo}
				directions = append(directions, to_append)
			}
		}
	}
	return directions
}

// follows the given direction in the 2d space
// returns true is the string is found in the direction provided
func matchString(arr [][]string, search string, x_start int, y_start int, x_direction int, y_direction int) bool {
	string_to_find := strings.Split(search, "")
	x := x_start
	y := y_start

	for _, s := range string_to_find {
		if !checkSpace(arr, s, x, y) {
			return false
		}
		x += x_direction
		y += y_direction
	}
	return true
}

// checks the value of the space in the 2d array
func checkSpace(arr [][]string, search string, x int, y int) bool {
	//prevent outofbounds
	if x < 0 || y < 0 {
		return false
	}
	//prevent outofbounds in the other direction now
	if y >= len(arr) || x >= len(arr[y]) {
		return false
	}
	return arr[y][x] == search
}
