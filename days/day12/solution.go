package main

import "strings"

type point struct {
	x, y int
}

var dirs = []point{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func part1(input string) any {
	garden := strings.Split(input, "\n")
	seen := make(map[point]struct{})
	// res := 0
	for i := 0; i < len(garden); i++ {
		for j := 0; j < len(garden[0]); j++ {
			c := point{i, j}
			if _, ok := seen[c]; ok {
				continue
			}
			position := c
			seen[position] = struct{}{}
		}
	}

	return nil
}

func part2(input string) any {
	for _, line := range strings.Split(input, "\n") {
		_ = line
	}

	return nil
}
