package main

import "strings"

type point struct {
	x, y int
}

var dirs = []point{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func part1(input string) any {
	grid := strings.Split(input, "\n")
	rows := len(grid)
	cols := len(grid[0])
	res := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '0' {
				res += trailheads(grid, point{i, j}, true)
			}
		}
	}

	return res
}

func part2(input string) any {
	grid := strings.Split(input, "\n")
	rows := len(grid)
	cols := len(grid[0])
	res := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '0' {
				res += trailheads(grid, point{i, j}, false)
			}
		}
	}

	return res
}

func trailheads(grid []string, start point, considerSeen bool) int {
	seen := make(map[point]bool)
	res := 0
	seen[start] = true
	queue := []point{start}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, dir := range dirs {
			next := point{cur.x + dir.x, cur.y + dir.y}
			if next.x < 0 || next.x >= len(grid) || next.y < 0 || next.y >= len(grid[0]) {
				continue
			}

			if grid[next.x][next.y]-grid[cur.x][cur.y] != 1 {
				continue
			}

			if seen[next] {
				continue
			}
			if considerSeen {
				seen[next] = true
			}
			if grid[next.x][next.y] == '9' {
				res++
				continue
			}
			queue = append(queue, next)
		}
	}

	return res
}
