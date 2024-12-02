package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//load inputs
	read, err := os.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}
	//convert to string
	input := string(read)

	//make slices
	leftList := []int{}
	rightList := []int{}

	//assign to slices
	input_list := strings.Split(input, "\r\n")
	for i := 0; i < len(input_list); i++ {
		line := strings.Split(input_list[i], "   ")
		v1, err := strconv.Atoi(line[0])
		if err != nil {
			panic(err)
		}
		v2, err := strconv.Atoi(line[1])
		if err != nil {
			panic(err)
		}
		leftList = append(leftList, v1)
		rightList = append(rightList, v2)

	}

	//sort each list
	sortArray(leftList[:])
	sortArray(rightList[:])

	//make list of distance
	distance := []int{}

	for i := 0; i < len(leftList); i++ {
		distance = append(distance, int(math.Abs(float64(leftList[i])-float64(rightList[i]))))
	}

	//totalize
	total := 0
	for _, v := range distance {
		total += v
	}
	fmt.Println("total:" + string(total))
}

func sortArray(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
}
