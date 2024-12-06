package main

import (
	read_file "aoc2024/library"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	file := read_file.ReadFile("./input.txt")
	line_breaks := read_file.SplitOnLine(file)
	var room [][]string
	var current_pos [2]int
	current_direction := "^"  //guard ticker
	next_step := []int{0, -1} //guards next movement
	guardInRoom := true       //track if the guard has left or not
	markedMap := make(map[string]bool)
	//for loop builds the room and gets guard position
	for y, l := range line_breaks {
		//get the current pos if we find it
		x := strings.Index(l, current_direction)
		if x != -1 {
			current_pos[0] = x
			current_pos[1] = y
		}
		line := strings.Split(l, "")

		room = append(room, line)
	}

	// main movement logic
	for guardInRoom {
		x := current_pos[0]
		y := current_pos[1]

		next_x := x + next_step[0]
		next_y := y + next_step[1]
		step := check_step(room, next_x, next_y)
		switch step {
		case 0:
			guardInRoom = false
		case 1: //normal step
			update_map(room, x, y, "x")
			current_pos[0] = next_x
			current_pos[1] = next_y
			update_map(room, current_pos[0], current_pos[1], current_direction)
		case 2: //make a right turn
			current_direction, next_step[0], next_step[1] = rightTurn(current_direction)
			update_map(room, current_pos[0], current_pos[1], current_direction)
		case 3: //guard has already been there, do almost nothing
			current_pos[0] = next_x
			current_pos[1] = next_y
			update_map(room, current_pos[0], current_pos[1], current_direction)
		}
		markedMap[strconv.Itoa(x)+","+strconv.Itoa(y)] = true
		// outputToFile(room)
		// fmt.Println(markedSPaces)
	}

	fmt.Println(len(markedMap))
}

// make sure we can take a step returns value based on what the guard will do
// 0 - guard is off map, we are done
// 1 - normal step forward
// 2 - guard will turn right
// 3 - guard has already stepped here
func check_step(arr [][]string, x int, y int) int {

	//guard will step off map
	if y < 0 || x < 0 {
		return 0
	}
	if y >= len(arr) {
		return 0
	}
	if x >= len(arr[y]) {
		return 0
	}
	//guard in map
	space := arr[y][x]
	if space == "#" {
		return 2
	}
	if space == "x" {
		return 3
	}
	return 1
}

func update_map(arr [][]string, x int, y int, marker string) {
	arr[y][x] = marker
}

func rightTurn(s string) (string, int, int) {
	switch s {
	case "^":
		return ">", 1, 0
	case ">":
		return "v", 0, 1
	case "v":
		return "<", -1, 0
	case "<":
		return "^", 0, -1
	default:
		panic("Bro where are you looking??")
	}
}

func outputToFile(file [][]string) {
	time.Sleep(100 * time.Millisecond)
	output := ""

	for _, y := range file {
		for _, x := range y {
			output += x
		}
		output += "\n"
	}
	_, err := os.Create("./output.txt")
	if err != nil {
		panic(err)
	}
	os.WriteFile("./output.txt", []byte(output), 0666)
}
