package day01p1

import (
	"aoc/utils"
	"io"
	"regexp"
	"strconv"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	rex := regexp.MustCompile(`mul\((?<a>\d+),(?<b>\d+)\)`)

	total := 0
	for _, line := range lines {
		matches := rex.FindAllStringSubmatch(line, -1)
		for _, v := range matches {
			a, _ := strconv.ParseInt(v[1], 10, 64)
			b, _ := strconv.ParseInt(v[2], 10, 64)
			total += int(a) * int(b)
		}
	}
	return total
}
