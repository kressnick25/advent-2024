package day01p1

import (
	"fmt"
	"io"
	"sort"
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

	sorted1 := make([]int, len(col1))
	copy(sorted1, col1)
	sorted2 := make([]int, len(col2))
	copy(sorted2, col2)

	sort.Slice(sorted1, func(i, j int) bool {
		return sorted1[i] < sorted1[j] 
	})
	sort.Slice(sorted2, func(i, j int) bool {
		return sorted2[i] < sorted2[j] 
	})

	total_distance := 0
	for i, val1 := range sorted1 {
		val2 := sorted2[i]
		// index1, err := indexOf(col1, val1)
		// if err != nil { panic("invalid value in col1") }
		// index2, err := indexOf(col2, val2)
		// if err != nil { panic("invalid value in col2") }

		distance := val2 - val1
		if distance < 0 {
			total_distance -= distance
		} else {
			total_distance += distance
		}
	}
 

	return total_distance
}
