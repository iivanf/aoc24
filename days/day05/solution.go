package main

import (
	"strconv"
	"strings"
)

func part1(input string) any {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	rules := make(map[int][]int)
	res := 0
	ind := 0

	// Parse the first part of the input
	for ; ind < len(lines) && lines[ind] != ""; ind++ {
		as, bs, _ := strings.Cut(lines[ind], "|")
		a, _ := strconv.Atoi(as)
		b, _ := strconv.Atoi(bs)
		rules[a] = append(rules[a], b)
	}

	// Process the second part of the input
	for _, line := range lines[ind+1:] {
		nums := strings.Split(line, ",")
		pages := make([]int, len(nums))
		index := make(map[int]int, len(nums))
		for i, num := range nums {
			num, _ := strconv.Atoi(num)
			pages[i] = num
			index[num] = i
		}

		isValid := true
		for numIndex, num := range pages {
			if numsAfter, ok := rules[num]; ok {
				for _, numAfter := range numsAfter {
					if afterIndex, ok := index[numAfter]; ok && afterIndex <= numIndex {
						isValid = false
						break
					}
				}
			}
			if !isValid {
				break
			}
		}

		if isValid {
			res += pages[len(pages)/2]
		}
	}

	return res
}

func part2(input string) any {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	rules := make(map[int][]int)
	res := 0
	ind := 0

	// Parse the first part of the input
	for ; ind < len(lines) && lines[ind] != ""; ind++ {
		as, bs, _ := strings.Cut(lines[ind], "|")
		a, _ := strconv.Atoi(as)
		b, _ := strconv.Atoi(bs)
		rules[a] = append(rules[a], b)
	}

	// Process the second part of the input
	for _, line := range lines[ind+1:] {
		nums := strings.Split(line, ",")
		pages := make([]int, len(nums))
		index := make(map[int]int, len(nums))
		for i, num := range nums {
			num, _ := strconv.Atoi(num)
			pages[i] = num
			index[num] = i
		}

		wasInvalid := false

		for {
			isValid := true
			for i, num := range pages {
				index[num] = i
			}
			for numIndex, num := range pages {
				if numsAfter, ok := rules[num]; ok {
					for _, numAfter := range numsAfter {
						if afterIndex, ok := index[numAfter]; ok && afterIndex <= numIndex {
							isValid = false
							wasInvalid = true
							pages[afterIndex], pages[numIndex] = pages[numIndex], pages[afterIndex]
							break
						}
					}
				}
				if !isValid {
					break
				}
			}
			if isValid {
				break
			}
		}

		if wasInvalid {
			res += pages[len(pages)/2]
		}
	}
	return res
}
