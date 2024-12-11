// 15753952992801 == too high
// 6368943744467 == too low

/*
*
I'm not doing part 2, I give up

using example input
last failed output:
00...111...2...333.44.5555.6666.777.888899
0099.111...2...333.44.5555.6666.777.8888..
00992111.......333.44.5555.6666.777.8888..
009921118888...333.44.5555.6666.777.......
009921118888777333.44.5555.6666...........
2102
*
*/
package main

import (
	read_file "aoc2024/library"
	"fmt"
	"strconv"
	"strings"
)

type fileblock struct {
	r     rune
	index int
	size  int
}

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

	files := getAllFileBlock(r)

	inBlock := false
	blockStart := 0
	blockSize := 0
	for i, v := range r {
		if v == 0 && !inBlock {
			inBlock = true
			blockStart = i
			blockSize++
		} else if v == 0 && inBlock {
			blockSize++
		} else if v != 0 && inBlock {
			for blockSize > 0 {
				var file fileblock
				var loc int
				file, files, loc = getFileBlockBySize(files, blockSize)
				if loc != -1 {
					printReadable(r)
					moveBlock(blockStart, file.size, r, file.r, file.index)
					blockStart += file.size
					blockSize = blockSize - file.size
				} else {
					blockSize = 0
				}

			}
			inBlock = false
			blockStart = 0
			blockSize = 0
		}

	}
	printReadable(r)
	return r

	//part 1 stuff
	// dotIndex := getFirstdot(r)
	// lastIdIndex, lastId := getLastId(r)
	// for dotIndex < lastIdIndex {
	// 	r[dotIndex] = lastId
	// 	r[lastIdIndex] = '\x00'
	// 	dotIndex = getFirstdot(r)
	// 	lastIdIndex, lastId = getLastId(r)
	// 	if lastIdIndex == -1 {
	// 		break
	// 	}
	// }
	// return r
}

func getAllFileBlock(r []rune) []fileblock {
	blocks := []fileblock{}
	inBlock := false
	blockSize := 0
	blockStart := 0
	blockRune := rune(0)
	for i, v := range r {
		if !inBlock && v == 0 {
			continue
		}
		if inBlock && (v == 0 || v != blockRune) {
			file := fileblock{
				blockRune,
				blockStart,
				blockSize,
			}
			blocks = append(blocks, file)
			inBlock = false
			blockSize = 0
		}
		if !inBlock && v != 0 {
			inBlock = true
			blockStart = i
			blockRune = v
		}
		if v != 0 {
			blockSize++
		}
	}
	//check if we ended on a block
	if inBlock {
		file := fileblock{
			blockRune,
			blockStart,
			blockSize,
		}
		blocks = append(blocks, file)
	}
	return blocks
}
func getFileBlockBySize(f []fileblock, s int) (fileblock, []fileblock, int) {
	for i := len(f) - 1; i > 0; i-- {
		file := f[i]
		if file.index > i && file.size <= s {
			list := RemoveIndex(f, i)
			return file, list, 0
		}
	}
	file := fileblock{
		rune(0),
		0,
		0,
	}
	return file, f, -1
}

func moveBlock(blockStart int, num int, r []rune, blockRune rune, removeStart int) {
	for j := blockStart; j < blockStart+num; j++ {
		r[j] = blockRune
	}

	for re := 0; re < num; re++ {
		r[removeStart+re] = rune(0)
	}
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
func printReadable(b []rune) {
	for _, v := range b {
		if v == 0 {
			fmt.Print(".")
		} else {
			fmt.Print(v - 1)
		}
	}
	fmt.Println()
}
func RemoveIndex(s []fileblock, i int) []fileblock {
	newSlice := make([]fileblock, len(s)-1)
	copy(newSlice, s[:i])
	copy(newSlice[i:], s[i+1:])
	return newSlice
}
