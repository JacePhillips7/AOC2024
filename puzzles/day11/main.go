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
	blinks := 25
	for range blinks {
		stones = mutate(stones)
		fmt.Println(stones)
	}
	fmt.Println(len(stones))
}

func mutate(s []int) []int {
	stones := s
	for i := 0; i < len(stones); i++ {
		v := stones[i]
		if v == 0 {
			stones[i] = 1
			continue
		}
		stringV := strconv.Itoa(v)
		arr := strings.Split(stringV, "")
		if len(arr)%2 == 0 {
			leftV, rightV := splitHalf(stringV)
			stones[i] = leftV
			if i != len(stones) {
				stones = append(stones[:i+1], stones[i:]...)
				i++
				stones[i] = rightV
			} else {
				stones = append(stones, rightV)
			}
			continue
		}
		stones[i] = v * 2024
	}
	return stones
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
