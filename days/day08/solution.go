package main

import "strings"

type point struct {
	x, y int
}

func part1(input string) any {
	grid := strings.Split(input, "\n")
	freqs := make(map[byte][]point)

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != '.' {
				freqs[grid[i][j]] = append(freqs[grid[i][j]], point{i, j})
			}
		}
	}

	seen := make(map[point]struct{})
	for _, antennas := range freqs {
		for i := 0; i < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				a1 := antennas[i]
				a2 := antennas[j]
				deltaX, deltaY := a1.x-a2.x, a1.y-a2.y
				c1 := point{a1.x + deltaX, a1.y + deltaY}
				c2 := point{a2.x - deltaX, a2.y - deltaY}
				if c1.x >= 0 && c1.x < len(grid) && c1.y >= 0 && c1.y < len(grid[0]) {
					seen[c1] = struct{}{}
				}
				if c2.x >= 0 && c2.x < len(grid) && c2.y >= 0 && c2.y < len(grid[0]) {
					seen[c2] = struct{}{}
				}
			}
		}
	}

	return len(seen)
}

func part2(input string) any {
	grid := strings.Split(input, "\n")
	freqs := make(map[byte][]point)

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != '.' {
				freqs[grid[i][j]] = append(freqs[grid[i][j]], point{i, j})
			}
		}
	}

	seen := make(map[point]struct{})
	for _, antennas := range freqs {
		for i := 0; i < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				a1 := antennas[i]
				a2 := antennas[j]
				seen[a1] = struct{}{}
				seen[a2] = struct{}{}
				deltaX, deltaY := a1.x-a2.x, a1.y-a2.y
				for k := 1; ; k++ {
					c1 := point{a1.x + k*deltaX, a1.y + k*deltaY}
					if c1.x >= 0 && c1.x < len(grid) && c1.y >= 0 && c1.y < len(grid[0]) {
						seen[c1] = struct{}{}
					} else {
						break
					}
				}
				for k := 1; ; k++ {
					c2 := point{a2.x - k*deltaX, a2.y - k*deltaY}
					if c2.x >= 0 && c2.x < len(grid) && c2.y >= 0 && c2.y < len(grid[0]) {
						seen[c2] = struct{}{}
					} else {
						break
					}
				}
			}
		}
	}

	return len(seen)
}
