package main

import (
	read_file "aoc2024/library"
	"fmt"
	"strconv"
	"strings"
)

type cord struct {
	x int
	y int
}

func main() {
	file := read_file.ReadFile("./input.txt")
	lines := read_file.SplitOnLine(file)
	//here we map out all the antennas
	//map of all cords like so "A":[{x:2,y:3}]
	//ignore these 2
	antennaMap := make(map[string][]cord)
	xBounds, yBounds := makeAntennaMap(lines, antennaMap)
	antinodeMap := make(map[string]string)
	mapAntinodes(antennaMap, antinodeMap, xBounds, yBounds)
	fmt.Println(len(antinodeMap))
}

func makeAntennaMap(lines []string, antennaMap map[string][]cord) (int, int) {
	for y, v := range lines {
		s := strings.Split(v, "")
		for x, l := range s {
			if l != "." && l != "#" {
				loc := cord{x, y}
				check := antennaMap[l]
				if check == nil {
					antennaMap[l] = []cord{
						loc,
					}
				} else {
					antennaMap[l] = append(antennaMap[l], loc)
				}
			}
		}
	}
	s := strings.Split(lines[0], "")
	return len(s) - 1, len(lines) - 1
}

func mapAntinodes(antennas map[string][]cord, nodes map[string]string, x int, y int) {
	for _, v := range antennas {
		found := findNodes(v, x, y)
		for _, f := range found {
			nodes[strconv.Itoa(f.x)+","+strconv.Itoa(f.y)] = "#"
		}
	}
}

func findNodes(antennas []cord, xbounds int, ybounds int) []cord {
	cords := []cord{}
	for _, loc := range antennas {
		for _, other := range antennas {
			if loc.x == other.x && loc.y == other.y { //we will ignore ourselves
				continue
			}
			//get dSlope and then double to find the node cord
			dSlope := cord{
				x: (other.x - loc.x) * 2,
				y: (other.y - loc.y) * 2,
			}

			node := cord{
				x: loc.x + dSlope.x,
				y: loc.y + dSlope.y,
			}
			if node.x > xbounds || node.x < 0 || node.y > ybounds || node.y < 0 { //off map, so we don't care
				continue
			}
			cords = append(cords, node)

		}
	}
	return cords
}
