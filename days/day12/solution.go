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

	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}

	result := 0

	for y := range grid {
		for x := range grid[y] {
			pos := point{x, y}
			if visited[pos.y][pos.x] {
				continue
			}

			edgeSet := make(map[point]map[point]bool)
			calculateRegionEdges(grid, pos, edgeSet, visited)

			sides := 0
			for pos, edges := range edgeSet {
				for edge := range edges {
					destructivelyRemoveEdge(edgeSet, pos, edge)
					sides++
				}
			}

			result += len(edgeSet) * sides
		}
	}

	return result
}

func calculateRegionEdges(grid []string, pos point, edgeSet map[point]map[point]bool, visited [][]bool) {
	visited[pos.y][pos.x] = true

	if edgeSet[pos] == nil {
		edgeSet[pos] = make(map[point]bool)
	}

	for _, dir := range dirs {
		neighbor := pos.Add(dir)
		if neighbor.IsInBounds(grid) {
			if neighbor.OfGrid(grid) != pos.OfGrid(grid) {
				edgeSet[pos][dir] = true
			} else if !visited[neighbor.y][neighbor.x] {
				calculateRegionEdges(grid, neighbor, edgeSet, visited)
			}
		} else {
			edgeSet[pos][dir] = true
		}
	}
}

func destructivelyRemoveEdge(edgeSet map[point]map[point]bool, pos point, edge point) {
	if !edgeSet[pos][edge] {
		return
	}
	delete(edgeSet[pos], edge)

	for _, dir := range []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		destructivelyRemoveEdge(edgeSet, pos.Add(dir), edge)
	}
}
