package main

//first attempt: 2151 -- too high
//second attempt: 1990 -- too high
import (
	read_file "aoc2024/library"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type gameState struct {
	room              [][]string
	current_pos       [2]int
	next_step         [2]int
	movementMarker    string
	current_direction string
}

func main() {

	file := read_file.ReadFile("./input.txt")
	line_breaks := read_file.SplitOnLine(file)
	var room [][]string
	var current_pos [2]int
	current_direction := "^"  //guard ticker
	next_step := []int{0, -1} //guards next movement
	guardInRoom := true       //track if the guard has left or not
	foundParadox := 0
	foundParadoxMap := make(map[string]bool)
	movementMarker := "|"
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
	for yIndex, room_line := range room {
		for xIndex, _ := range room_line {
			gameRoom := deepCopyRoom(room)
			update_map(gameRoom, xIndex, yIndex, "#")
			game := gameState{gameRoom, current_pos, [2]int(next_step), movementMarker, current_direction}
			gameTest, paradoxX, paradoxY := runGame(game)
			if gameTest {
				foundParadox++
				foundParadoxMap[strconv.Itoa(paradoxX)+","+strconv.Itoa(paradoxY)] = true
			}
		}
	}

	fmt.Println("Paradox amount: " + strconv.Itoa(foundParadox))
	fmt.Println("Paradox map amount: " + strconv.Itoa(len(foundParadoxMap)))
	return
	// main movement logic
	update_map(room, current_pos[0], current_pos[1], movementMarker)
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
			current_pos[0] = next_x
			current_pos[1] = next_y
			update_map(room, current_pos[0], current_pos[1], movementMarker)
		case 2: //make a right turn
			current_direction, next_step[0], next_step[1], movementMarker = rightTurn(current_direction)
			update_map(room, current_pos[0], current_pos[1], "+")
		case 3: //guard has already been there, do almost nothing
			current_pos[0] = next_x
			current_pos[1] = next_y
			update_map(room, current_pos[0], current_pos[1], "+")
		}
		markedMap[strconv.Itoa(x)+","+strconv.Itoa(y)] = true
		// outputToFile(room)
	}

	outputToFile(room)
	fmt.Println("unique steps:" + strconv.Itoa(len(markedMap)))
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
	if space == "x" || space == "-" || space == "|" || space == "o" {
		return 3
	}
	return 1
}

func update_map(arr [][]string, x int, y int, marker string) {
	arr[y][x] = marker
}

func rightTurn(s string) (string, int, int, string) {
	switch s {
	case "^":
		return ">", 1, 0, "-"
	case ">":
		return "v", 0, 1, "|"
	case "v":
		return "<", -1, 0, "-"
	case "<":
		return "^", 0, -1, "|"
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

// returns true on paradox, false on guard exit, along with paradox location
func runGame(game gameState) (bool, int, int) {
	room := game.room
	current_pos := game.current_pos
	next_step := game.next_step
	// movementMarker := game.movementMarker
	current_direction := game.current_direction
	guardInRoom := true

	next_x := current_pos[0] + next_step[0]
	next_y := current_pos[1] + next_step[1]
	//before we start, add a object directly in our path
	newObjLoc := []int{next_x, next_y}
	//prevent making an obj out of bounds
	checkStep := check_step(room, next_x, next_y)
	if checkStep == 0 {
		return false, 0, 0
	}
	// update_map(room, next_x, next_y, "#") removed for bruteforce

	paradoxMap := make(map[string]bool)

	//start game here
	for guardInRoom {
		x := current_pos[0]
		y := current_pos[1]

		next_x := x + next_step[0]
		next_y := y + next_step[1]
		step := check_step(room, next_x, next_y)
		switch step {
		case 0:
			guardInRoom = false
			return false, 0, 0
		case 1: //normal step
			current_pos[0] = next_x
			current_pos[1] = next_y
			// update_map(room, current_pos[0], current_pos[1], movementMarker)
		case 2: //make a right turn
			current_direction, next_step[0], next_step[1], _ = rightTurn(current_direction)
			// update_map(room, current_pos[0], current_pos[1], "+")
		case 3: //guard has already been there, do almost nothing
			current_pos[0] = next_x
			current_pos[1] = next_y
			// update_map(room, current_pos[0], current_pos[1], "+")
		}
		paradoxString := strconv.Itoa(current_pos[0]) + "," + strconv.Itoa(current_pos[1]) + "," + current_direction
		if paradoxMap[paradoxString] {
			fmt.Print("Paradox at: ")
			fmt.Print(newObjLoc)
			fmt.Println()
			return true, newObjLoc[0], newObjLoc[1]
		}
		paradoxMap[paradoxString] = true
	}

	return false, 0, 0
}
func deepCopyRoom(original [][]string) [][]string {
	copyRoom := make([][]string, len(original))
	for i := range original {
		copyRoom[i] = make([]string, len(original[i]))
		copy(copyRoom[i], original[i])
	}
	return copyRoom
}
