package main

import (
	"strings"
)

var directions [][]int = [][]int{
	{0, 1},   // up
	{0, -1},  //down
	{1, 0},   // right
	{-1, 0},  // left
	{1, 1},   // up right
	{-1, 1},  // up left
	{1, -1},  // down right
	{-1, -1}, // down left
}

var wordList []string = []string{"X", "M", "A", "S"}

func part1(input string) any {
	grid := [][]string{}
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		row := strings.Split(line, "")
		grid = append(grid, row)
	}
	res := 0
	for x, row := range grid {
		for y, char := range row {
			if char == wordList[0] {
				for _, direction := range directions {
					if findXMAS(x, y, 1, direction, grid) {
						res++
					}
				}
			}
		}
	}
	return res
}

func part2(input string) any {
	grid := [][]string{}
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		row := strings.Split(line, "")
		grid = append(grid, row)
	}
	res := 0
	for x, row := range grid {
		for y, char := range row {
			if char == wordList[2] && findMAS(x, y, grid) {
				res++
			}
		}
	}
	return res
}

func findXMAS(x, y, wordPosition int, direction []int, grid [][]string) bool {
	for wordPosition < len(wordList) {
		x += direction[0]
		y += direction[1]
		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[x]) || grid[x][y] != wordList[wordPosition] {
			return false
		}
		wordPosition++
	}
	return true
}

func findMAS(x, y int, grid [][]string) bool {
	if x-1 >= 0 && y+1 < len(grid[0]) && x+1 < len(grid) && y-1 >= 0 {
		if (grid[x-1][y+1] == wordList[1] && grid[x+1][y-1] == wordList[3]) || (grid[x-1][y+1] == wordList[3] && grid[x+1][y-1] == wordList[1]) {
			if (grid[x+1][y+1] == wordList[1] && grid[x-1][y-1] == wordList[3]) || (grid[x+1][y+1] == wordList[3] && grid[x-1][y-1] == wordList[1]) {
				return true
			}
		}
	}
	return false
}
