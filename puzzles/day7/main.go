package main

import (
	read_file "aoc2024/library"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	file := read_file.ReadFile("./input.txt")
	problems := read_file.SplitOnLine(file)
	var total uint64
	total = 0
	for _, p := range problems {
		cleanString := strings.ReplaceAll(p, ":", "")
		parts := strings.Split(cleanString, " ")
		s, n := arrPop(parts)
		solution, _ := strconv.Atoi(s)
		numbs := make([]uint64, len(n))
		for i, v := range n {
			x, _ := strconv.Atoi(v)
			numbs[i] = uint64(x)
		}
		can, t := canSolve(uint64(solution), numbs)
		if can {
			total += t
		}
	}
	fmt.Println(total)
}

func canSolve(solution uint64, numbs []uint64) (bool, uint64) {
	//fill the operators
	operators := createAllPossible(len(numbs))
	return runTest(numbs, operators, solution)
}

func createAllPossible(n int) [][]string {
	result := [][]string{}
	arr := make([]string, n)

	var generate func(index int)
	generate = func(index int) {
		if index == n {
			// Make a copy of the current combination and add it to result
			combination := make([]string, n)
			copy(combination, arr)
			result = append(result, combination)
			return
		}

		// Try "+" at the current position and recurse
		arr[index] = "+"
		generate(index + 1)

		// Try "*" at the current position and recurse
		arr[index] = "*"
		generate(index + 1)

		//part 2 add the || operator
		arr[index] = "||"
		generate(index + 1)
	}

	generate(0) // Start recursion
	return result
}

func runTest(numbs []uint64, possible_operators [][]string, solution uint64) (bool, uint64) {
	j := 0
	for j < len(possible_operators) {
		operators := possible_operators[j]
		test_solution := numbs[0]
		for i := 1; i < len(numbs); i++ {
			if operators[i-1] == "+" {
				test_solution += numbs[i]
			} else if operators[i-1] == "*" {
				test_solution = test_solution * numbs[i]
			} else if operators[i-1] == "||" {
				s := strconv.Itoa(int(test_solution)) + strconv.Itoa(int(numbs[i]))
				parseme, _ := strconv.Atoi(s)
				test_solution = uint64(parseme)

			}
		}
		if test_solution == solution {
			return true, test_solution
		}
		j++
	}

	return false, 0
}

func arrPop(arr []string) (string, []string) {
	popper := arr[0]
	return popper, arr[1:]
}
