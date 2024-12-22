package main

import (
	"math"
	"strconv"
	"strings"
)

func part1(input string) any {
	res := 0
	for _, line := range strings.Split(input, "\n") {
		nums := strings.Fields(line)
		var sl1 []int
		for _, num := range nums {
			n1, err := strconv.Atoi(num)
			if err != nil {
				return err
			}
			sl1 = append(sl1, n1)
		}

		if isSafe(sl1) {
			res++
		}
	}

	return res
}

func isSafe(sl1 []int) bool {
	if len(sl1) < 2 {
		return true
	}

	increasing := true
	decreasing := true

	for i := 1; i < len(sl1); i++ {
		diff := math.Abs(float64(sl1[i] - sl1[i-1]))
		if diff < 1 || diff > 3 {
			return false
		}
		if sl1[i] < sl1[i-1] {
			increasing = false
		}
		if sl1[i] > sl1[i-1] {
			decreasing = false
		}
	}

	return increasing || decreasing
}

func part2(input string) any {
	res := 0
	for _, line := range strings.Split(input, "\n") {
		nums := strings.Fields(line)
		var sl1 []int
		for _, num := range nums {
			n1, err := strconv.Atoi(num)
			if err != nil {
				return err
			}
			sl1 = append(sl1, n1)
		}

		if validateByRemoving(sl1) {
			res++
			continue
		}
	}
	return res
}

func validateByRemoving(data []int) bool {
	for i := 0; i < len(data); i++ {
		rmv1 := append([]int{}, data[:i]...)
		rmv1 = append(rmv1, data[i+1:]...)
		if isSafe(rmv1) {
			return true
		}
	}
	return false
}
