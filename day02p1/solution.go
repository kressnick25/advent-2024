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

func isIncreasing(s []int) bool {
	tmp := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] <= tmp {
			return false
		}
		tmp = s[i]
	}
	return true
}

func isDecreasing(s []int) bool {
	tmp := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] >= tmp {
			return false
		}
		tmp = s[i]
	}
	return true
}

func safeLevels(s []int) bool {
	tmp := s[0]
	for i := 1; i < len(s); i++ {
		if utils.Abs(s[i] - tmp) > 3 {
			return false
		}
		tmp = s[i]
	}
	return true
}

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	reports := make([][]int, len(lines))
	for i, l := range lines {
		report_vals := strings.Fields(l)
		report_val_ints := make([]int, len(report_vals))
		for j, val := range report_vals {
			int1, _ := strconv.ParseInt(val, 10, 64)
			report_val_ints[j] = int(int1)
		}
		reports[i] = report_val_ints
	}
	
	safe := 0
	for _, report := range reports {
		if safeLevels(report) && (isDecreasing(report) || isIncreasing(report)) {
			safe += 1
		}
	}

	return safe
}
