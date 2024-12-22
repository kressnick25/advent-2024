package day01p1

import (
	"io"
	"strconv"
	"strings"

	"aoc/utils"
)

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
		if utils.Abs(s[i]-tmp) > 3 {
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
		} else {
			// just brute force all combinations for this report
			for i := range report {
				// remove the element i and try
				r := make([]int, len(report))
				copy(r, report)
				r = utils.Remove(r, i)
				if safeLevels(r) && (isDecreasing(r) || isIncreasing(r)) {
					safe += 1
					break
				}
			}
		}
	}

	return safe
}
