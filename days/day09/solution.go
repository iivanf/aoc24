package main

import (
	"strings"
)

func part1(input string) any {
	line := strings.TrimSpace(input)
	res := make([]int, 0, len(line))
	curInd := 0
	freePos := -1

	for i := 0; i < len(line); i += 2 {
		n := int(line[i] - '0')
		free := 0
		if i+1 < len(line) {
			free = int(line[i+1] - '0')
		}

		for j := 0; j < n; j++ {
			res = append(res, curInd)
		}
		for j := 0; j < free; j++ {
			if freePos == -1 {
				freePos = len(res)
			}
			res = append(res, -1)
		}
		curInd++
	}

	curPos := len(res) - 1
	for freePos < len(res) && freePos < curPos {
		res[freePos] = res[curPos]
		res[curPos] = -1
		for freePos < len(res) && res[freePos] != -1 {
			freePos++
		}
		for curPos > freePos && res[curPos] == -1 {
			curPos--
		}
	}

	checksum := 0
	for i, n := range res {
		if n != -1 {
			checksum += i * n
		}
	}

	return checksum
}

type block struct {
	fileID int
	cnt    int
}

func part2(input string) any {
	line := strings.TrimSpace(input)
	res := make([]block, 0, len(line)/2)
	curInd := 0

	for i := 0; i < len(line); i += 2 {
		n := int(line[i] - '0')
		res = append(res, block{curInd, n})
		if i+1 < len(line) {
			free := int(line[i+1] - '0')
			res = append(res, block{-1, free})
		}
		curInd++
	}

	for fileID := curInd - 1; fileID >= 0; fileID-- {
		var blockInd, freeInd int
		var mblock block
		for blockInd = 0; blockInd < len(res); blockInd++ {
			if res[blockInd].fileID == fileID {
				mblock = res[blockInd]
				break
			}
		}

		for freeInd = 0; freeInd < len(res); freeInd++ {
			if res[freeInd].fileID == -1 && res[freeInd].cnt >= mblock.cnt {
				break
			}
		}
		if freeInd == len(res) || freeInd >= blockInd {
			continue
		}

		res[blockInd].fileID = -1
		res[freeInd].cnt -= mblock.cnt
		if res[freeInd].cnt == 0 {
			res[freeInd] = mblock
		} else {
			res = append(res[:freeInd+1], res[freeInd:]...)
			res[freeInd] = mblock
		}
	}

	checksum := 0
	t := 0
	for _, mblock := range res {
		for j := 0; j < mblock.cnt; j++ {
			if mblock.fileID != -1 {
				checksum += t * mblock.fileID
			}
			t++
		}
	}

	return checksum
}
