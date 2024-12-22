package main

import (
	"strings"
)

type point struct {
	x, y int
}

var directions = []point{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func part1(input string) any {
	grid := strings.Split(input, "\n")
	obstacles := make(map[point]struct{})
	path := make(map[point]struct{})
	dirIndex := 0 // up
	var start point

	for x, line := range grid {
		for y, c := range line {
			switch c {
			case '#':
				obstacles[point{x, y}] = struct{}{}
			case '^':
				start = point{x, y}
			}
		}
	}

	position := start
	path[position] = struct{}{}

	for {
		dir := directions[dirIndex]
		next := point{position.x + dir.x, position.y + dir.y}

		if next.x < 0 || next.x >= len(grid) || next.y < 0 || next.y >= len(grid[0]) {
			break
		}

		if _, ok := obstacles[next]; ok {
			dirIndex = (dirIndex + 1) % len(directions)
			continue
		}

		position = next
		path[position] = struct{}{}
	}

	return len(path)
}

func part2(input string) any {
	grid := strings.Split(input, "\n")
	obstacles := make(map[point]struct{})
	freePoints := []point{}

	var start point

	for x, line := range grid {
		for y, c := range line {
			switch c {
			case '#':
				obstacles[point{x, y}] = struct{}{}
			case '^':
				start = point{x, y}
			case '.':
				freePoints = append(freePoints, point{x, y})
			}
		}
	}

	loopCount := 0

	for _, obstacle := range freePoints {
		visited := make(map[point]point)
		obstacles[obstacle] = struct{}{}
		position := start
		loopDetected := false
		dirIndex := 0
		visited[position] = directions[dirIndex]

		for {
			dir := directions[dirIndex]
			next := point{position.x + dir.x, position.y + dir.y}
			if next.x < 0 || next.x >= len(grid) || next.y < 0 || next.y >= len(grid[0]) {
				break
			}

			if _, ok := obstacles[next]; ok {
				dirIndex = (dirIndex + 1) % len(directions)
				continue
			}

			if visited[next] == dir {
				loopDetected = true
				break
			}

			position = next
			visited[position] = dir
		}

		if loopDetected {
			loopCount++
		}

		delete(obstacles, obstacle)
	}

	return loopCount
}
