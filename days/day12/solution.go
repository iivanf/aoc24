package main

import (
	"strings"
)

type point struct {
	x, y int
}

func (p point) Add(other point) point {
	return point{p.x + other.x, p.y + other.y}
}

func (p point) IsInBounds(grid []string) bool {
	return p.x >= 0 && p.x < len(grid) && p.y >= 0 && p.y < len(grid[p.x])
}

func (p point) OfGrid(grid []string) rune {
	return rune(grid[p.x][p.y])
}

var dirs = []point{
	{1, 0}, {-1, 0}, {0, 1}, {0, -1},
}

func part1(input string) any {
	grid := strings.Split(strings.TrimSpace(input), "\n")

	addedGrid := make([][]bool, len(grid))
	for i := range addedGrid {
		addedGrid[i] = make([]bool, len(grid[i]))
	}

	result := 0
	for x := range grid {
		for y, val := range grid[x] {
			pos := point{x, y}
			if addedGrid[pos.x][pos.y] {
				continue
			}

			group := []point{pos}
			addedGrid[pos.x][pos.y] = true

			for i := 0; i < len(group); i++ {
				current := group[i]
				for _, dir := range dirs {
					nextPos := current.Add(dir)
					if nextPos.IsInBounds(grid) && !addedGrid[nextPos.x][nextPos.y] && nextPos.OfGrid(grid) == val {
						group = append(group, nextPos)
						addedGrid[nextPos.x][nextPos.y] = true
					}
				}
			}

			fenceLen := 0
			for _, p := range group {
				for _, dir := range dirs {
					checkPos := p.Add(dir)
					if !checkPos.IsInBounds(grid) || checkPos.OfGrid(grid) != val {
						fenceLen++
					}
				}
			}

			result += len(group) * fenceLen
		}
	}

	return result
}

func part2(input string) any {
	grid := strings.Split(strings.TrimSpace(input), "\n")

	addedGrid := make([][]bool, len(grid))
	for i := range addedGrid {
		addedGrid[i] = make([]bool, len(grid[i]))
	}

	result := 0
	for y := range grid {
		for x, val := range grid[y] {
			pos := point{x, y}
			if addedGrid[pos.y][pos.x] {
				continue
			}

			group := []point{pos}
			addedGrid[pos.y][pos.x] = true

			for i := 0; i < len(group); i++ {
				current := group[i]
				for _, dir := range dirs {
					nextPos := current.Add(dir)
					if nextPos.IsInBounds(grid) && !addedGrid[nextPos.y][nextPos.x] && nextPos.OfGrid(grid) == val {
						group = append(group, nextPos)
						addedGrid[nextPos.y][nextPos.x] = true
					}
				}
			}

			esquinasTotales := 0
			for _, p := range group {
				esquinasTotales += checkCorners(grid, p)
			}
			result += len(group) * esquinasTotales
		}
	}

	return result
}

func checkCorners(input []string, current point) int {
	count := 0
	gardenType := input[current.y][current.x]
	x, y := current.x, current.y

	if x == 0 && y == 0 {
		count++
	}
	if x == 0 && y == len(input)-1 {
		count++
	}
	if x == len(input[0])-1 && y == len(input)-1 {
		count++
	}
	if x == len(input[0])-1 && y == 0 {
		count++
	}

	if (x > 0 && y > 0 && input[y][x-1] != gardenType && input[y-1][x] != gardenType) ||
		(x > 0 && y == 0 && input[y][x-1] != gardenType) || (x == 0 && y > 0 && input[y-1][x] != gardenType) {
		count++
	}

	if (x < len(input[0])-1 && y > 0 && input[y][x+1] != gardenType && input[y-1][x] != gardenType) ||
		(x < len(input[0])-1 && y == 0 && input[y][x+1] != gardenType) || (x == len(input[0])-1 && y > 0 && input[y-1][x] != gardenType) {
		count++
	}

	if (x > 0 && y < len(input)-1 && input[y][x-1] != gardenType && input[y+1][x] != gardenType) ||
		(x > 0 && y == len(input)-1 && input[y][x-1] != gardenType) || (x == 0 && y < len(input)-1 && input[y+1][x] != gardenType) {
		count++
	}

	if (x < len(input[0])-1 && y < len(input)-1 && input[y][x+1] != gardenType && input[y+1][x] != gardenType) ||
		(x < len(input[0])-1 && y == len(input)-1 && input[y][x+1] != gardenType) || (x == len(input[0])-1 && y < len(input)-1 && input[y+1][x] != gardenType) {
		count++
	}

	if x < len(input[0])-1 && y < len(input)-1 && input[y][x+1] == gardenType && input[y+1][x] == gardenType && input[y+1][x+1] != gardenType {
		count++
	}

	if x > 0 && y < len(input)-1 && input[y][x-1] == gardenType && input[y+1][x] == gardenType && input[y+1][x-1] != gardenType {
		count++
	}

	if x < len(input[0])-1 && y > 0 && input[y][x+1] == gardenType && input[y-1][x] == gardenType && input[y-1][x+1] != gardenType {
		count++
	}

	if x > 0 && y > 0 && input[y][x-1] == gardenType && input[y-1][x] == gardenType && input[y-1][x-1] != gardenType {
		count++
	}

	return count
}
