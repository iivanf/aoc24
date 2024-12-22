package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func part1(input string) any {
	res := 0
	for _, line := range strings.Split(input, "\n") {
		colonInd := strings.Index(line, ":")
		want, _ := strconv.Atoi(line[:colonInd])
		numsS := strings.Fields(line[colonInd+2:])
		nums := make([]int, 0, len(numsS))
		for _, numS := range numsS {
			num, _ := strconv.Atoi(numS)
			nums = append(nums, num)
		}

		result := []int{nums[0]}
		for i := 1; i < len(nums); i++ {
			partials := make([]int, 0, len(result)*2)
			for _, res := range result {
				partials = append(partials, res+nums[i], res*nums[i])
			}
			result = partials
		}

		if slices.Contains(result, want) {
			res += want
		}
	}

	return res
}

func part2(input string) any {
	res := 0
	for _, line := range strings.Split(input, "\n") {
		colonInd := strings.Index(line, ":")
		want, _ := strconv.Atoi(line[:colonInd])
		numsS := strings.Fields(line[colonInd+2:])
		nums := make([]int, 0, len(numsS))
		for _, numS := range numsS {
			num, _ := strconv.Atoi(numS)
			nums = append(nums, num)
		}

		result := []int{nums[0]}
		for i := 1; i < len(nums); i++ {
			partials := make([]int, 0, len(result)*3)
			for _, res := range result {
				partials = append(partials, res+nums[i], res*nums[i], concat(res, nums[i]))
			}
			result = partials
		}

		if slices.Contains(result, want) {
			res += want
		}
	}

	return res
}

func concat(a, b int) int {
	n, _ := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	return n
}
