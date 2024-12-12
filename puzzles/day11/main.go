// 65601038650482 -- too low
package main

import (
	read_file "aoc2024/library"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	file := read_file.ReadFile("./input.txt")
	s := strings.Split(file, " ")
	stones := toInts(s)
	stoneMap := make(map[int]int)
	blinks := 75
	//set map
	for _, v := range stones {
		stoneMap[v]++
	}
	for range blinks {
		stoneMap = mutate(stoneMap)
		// fmt.Println(stoneMap)
		// fmt.Println(len(stoneMap))
	}
	total := 0
	for v := range stoneMap {
		total += stoneMap[v]
	}
	fmt.Println(total)
}

func mutate(stoneMap map[int]int) map[int]int {
	updatedMap := make(map[int]int)
	for v := range stoneMap {
		if v == 0 {
			updatedMap[1] += stoneMap[v]
			continue
		}
		stringV := strconv.Itoa(v)
		arr := strings.Split(stringV, "")
		if len(arr)%2 == 0 {
			leftV, rightV := splitHalf(stringV)
			updatedMap[leftV] += stoneMap[v]
			updatedMap[rightV] += stoneMap[v]
			continue
		}
		val := v * 2024
		updatedMap[val] += stoneMap[v]
	}
	return updatedMap
}

func toInts(s []string) []int {
	ins := []int{}
	for _, v := range s {
		value, _ := strconv.Atoi(v)
		ins = append(ins, value)
	}
	return ins
}

func splitHalf(s string) (int, int) {
	left := ""
	right := ""
	arr := strings.Split(s, "")
	for i, v := range arr {
		if i < len(arr)/2 {
			left += v
		} else {
			right += v
		}
	}
	leftV, _ := strconv.Atoi(left)
	rightV, _ := strconv.Atoi(right)
	return leftV, rightV
}
