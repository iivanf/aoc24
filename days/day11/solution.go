package main

import (
	"strconv"
	"strings"
)

func part1(input string) any {
	stonesStr := strings.Fields(strings.TrimSpace(input))
	stones := []int{}

	for _, stoneStr := range stonesStr {
		stone, _ := strconv.Atoi(stoneStr)
		stones = append(stones, stone)
	}

	for range 25 {
		newStones := []int{}

		for _, stone := range stones {
			if stone == 0 {
				newStones = append(newStones, 1)
				continue
			}

			v := strconv.Itoa(stone)
			if len(v)%2 == 0 {
				first, _ := strconv.Atoi(v[:len(v)/2])
				second, _ := strconv.Atoi(v[len(v)/2:])
				newStones = append(newStones, first, second)
				continue
			}

			newStones = append(newStones, stone*2024)
		}
		stones = newStones
	}

	return len(stones)
}

func part2(input string) any {
	stonesStr := strings.Fields(strings.TrimSpace(input))
	stones := make(map[int]int)

	for _, stoneStr := range stonesStr {
		stone, _ := strconv.Atoi(stoneStr)
		stones[stone]++
	}

	for i := 0; i < 75; i++ {
		newStones := make(map[int]int)

		for stone, count := range stones {
			if stone == 0 {
				newStones[1] += count
				continue
			}

			digits := 0
			temp := stone
			for temp > 0 {
				temp /= 10
				digits++
			}

			if digits%2 == 0 {
				divisor := 1
				for k := 0; k < digits/2; k++ {
					divisor *= 10
				}
				first := stone / divisor
				second := stone % divisor
				newStones[first] += count
				newStones[second] += count
			} else {
				newStones[stone*2024] += count
			}
		}

		stones = newStones
	}

	totalCount := 0
	for _, count := range stones {
		totalCount += count
	}

	return totalCount
}
