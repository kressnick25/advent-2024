package day01p1

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"aoc/utils"
)

func indexOf(s []int, v int) (int, error) {
	for i, val := range s {
		if val == v {
			return i, nil
		}
	}
	return -1, fmt.Errorf("%d was not found in array", v)
}

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	col1 := make([]int, len(lines))
	col2 := make([]int, len(lines))
	for i, line := range lines {
		splits := strings.Fields(line)
		int1, _ := strconv.ParseInt(splits[0], 10, 64)
		int2, _ := strconv.ParseInt(splits[1], 10, 64)
		col1[i] = int(int1)
		col2[i] = int(int2)
	}

	keys2 := make(map[int]int)
	for _, val := range col2 {
		current, ok := keys2[val]
		if ok {
			keys2[val] = current + 1
		} else {
			keys2[val] = 1
		}
	}

	sim := 0
	for _, val := range col1 {
		col2Count, ok := keys2[val]
		if ok {
			sim += col2Count * val
		}
	}

	return sim
}
