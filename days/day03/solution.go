package main

import (
	"regexp"
	"strconv"
	"strings"
)

func part1(input string) any {
	res := 0
	for _, line := range strings.Split(input, "\n") {
		re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
		matches := re.FindAllString(line, -1)
		for _, match := range matches {
			inner := strings.TrimSuffix(strings.TrimPrefix(match, "mul("), ")")
			parts := strings.Split(inner, ",")
			n1, _ := strconv.Atoi(parts[0])
			n2, _ := strconv.Atoi(parts[1])
			res += n1 * n2
		}
	}

	return res
}

func part2(input string) any {
	res := 0
	enabled := true
	for _, line := range strings.Split(input, "\n") {
		re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
		matches := re.FindAllString(line, -1)
		for _, match := range matches {
			switch match {
			case "do()":
				enabled = true
			case "don't()":
				enabled = false
			default:
				if enabled {
					inner := strings.TrimSuffix(strings.TrimPrefix(match, "mul("), ")")
					parts := strings.Split(inner, ",")
					n1, _ := strconv.Atoi(parts[0])
					n2, _ := strconv.Atoi(parts[1])
					res += n1 * n2
				}
			}
		}
	}

	return res
}
