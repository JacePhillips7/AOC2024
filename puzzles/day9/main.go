package main

import (
	read_file "aoc2024/library"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	file := read_file.ReadFile("./input.txt")
	blocks := createBlocks(file)
	// fmt.Println(blocks)
	part := diskPart(blocks)
	// fmt.Println(part)
	sum := checksum(part)
	fmt.Println(sum)
}

func createBlocks(s string) []rune {
	arr := strings.Split(s, "")
	runes := make([]rune, 0)
	isFile := true
	id := 1
	for _, v := range arr {
		placement := rune(0)
		if isFile {
			placement = rune(id)
			id++
		}

		size, _ := strconv.Atoi(v)
		for d := 0; d < size; d++ {
			runes = append(runes, placement)
		}

		isFile = !isFile
	}
	return runes
}
func diskPart(r []rune) []rune {

	dotIndex := getFirstdot(r)
	lastIdIndex, lastId := getLastId(r)
	for dotIndex < lastIdIndex {
		r[dotIndex] = lastId
		r[lastIdIndex] = '\x00'
		dotIndex = getFirstdot(r)
		lastIdIndex, lastId = getLastId(r)
		if lastIdIndex == -1 {
			break
		}
	}
	return r
}

func getLastId(b []rune) (int, rune) {
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] != '\x00' {
			return i, b[i]
		}
	}
	return -1, rune(0)

}

func getFirstdot(b []rune) int {
	for i := 0; i < len(b); i++ {
		if b[i] == '\x00' {
			return i
		}
	}
	return -1
}

func checksum(b []rune) int {
	sum := 0
	for i, v := range b {
		if v == '\x00' || int(v) == 0 {
			continue
		}
		sum += i * (int(v) - 1)
	}
	return sum
}
