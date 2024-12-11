package main

import (
	read_file "aoc2024/library"
	"fmt"
	"strconv"
	"strings"
)

type ICords struct {
	x int
	y int
}

func main() {
	file := read_file.ReadFile("input.txt")
	lines := read_file.SplitOnLine(file)
	myMap := [][]int{}
	for _, v := range lines {
		l := strings.Split(v, "")
		arr := []int{}
		for _, conv := range l {
			value, _ := strconv.Atoi(conv)
			arr = append(arr, value)
		}
		myMap = append(myMap, arr)
	}

	//start of logic
	sum := 0
	trailHeads := findTrailheads(myMap)
	for _, v := range trailHeads {
		heads := trailBlaze(myMap, v)
		sum += heads
		fmt.Printf("[%v,%v] found %v paths", v.x, v.y, heads)
		fmt.Println()
	}
	fmt.Println(sum)
}

func findTrailheads(trail [][]int) []ICords {
	foundCords := []ICords{}
	for y, row := range trail {
		for x, value := range row {
			if value == 0 {
				cord := ICords{x: x, y: y}
				foundCords = append(foundCords, cord)
			}
		}
	}
	return foundCords
}

func trailBlaze(trail [][]int, loc ICords) int {
	value := trail[loc.y][loc.x]
	if value == 9 {
		return 1
	}
	paths := getNextSteps(trail, loc)
	total := 0

	for _, nextLoc := range paths {
		total += trailBlaze(trail, nextLoc)
	}
	return total
}

func getNextSteps(trail [][]int, loc ICords) []ICords {
	canBlaze := []ICords{}
	locationValue := trail[loc.y][loc.x]
	// we have 4 locations to check
	x1 := ICords{x: loc.x + 1, y: loc.y}
	x2 := ICords{x: loc.x - 1, y: loc.y}
	y1 := ICords{x: loc.x, y: loc.y + 1}
	y2 := ICords{x: loc.x, y: loc.y - 1}
	possible := []ICords{
		x1, x2, y1, y2,
	}
	for _, v := range possible {
		//checking out of bounds
		if v.x < 0 || v.y < 0 {
			continue
		}
		if v.x >= len(trail[0]) || v.y >= len(trail) {
			continue
		}
		//check if we can continue
		testValue := trail[v.y][v.x]
		if testValue == (locationValue + 1) {
			canBlaze = append(canBlaze, v)
		}
	}
	return canBlaze
}
