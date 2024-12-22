package main

import (
	"slices"
	"strconv"
	"strings"
)

func part1(input string) any {
	sl1 := []int{}
	sl2 := []int{}
	for _, line := range strings.Split(input, "\n") {
		s1, s2, _ := strings.Cut(line, "   ")
		n1, _ := strconv.Atoi(s1)
		n2, _ := strconv.Atoi(s2)
		sl1 = append(sl1, n1)
		sl2 = append(sl2, n2)
	}
	slices.Sort(sl1)
	slices.Sort(sl2)

	sum := 0
	for i := range sl1 {
		sum += abs(sl2[i] - sl1[i])
	}
	return sum
}

func part2(input string) any {
	sl1 := []int{}
	m := map[int]int{}
	for _, line := range strings.Split(input, "\n") {
		s1, s2, _ := strings.Cut(line, "   ")
		n1, _ := strconv.Atoi(s1)
		n2, _ := strconv.Atoi(s2)
		sl1 = append(sl1, n1)
		m[n2]++
	}
	sum := 0
	for i := range sl1 {
		sum += sl1[i] * m[sl1[i]]
	}

	return sum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
