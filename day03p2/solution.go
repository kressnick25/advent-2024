package day01p1

import (
	"aoc/utils"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	rex := regexp.MustCompile(`mul\((?<a>\d+),(?<b>\d+)\)|don't\(\)|do\(\)`)

	total := 0
	enabled := true
	for _, line := range lines {
		matches := rex.FindAllStringSubmatch(line, -1)
		for _, v := range matches {
			val := v[0]
			if strings.Contains(val, "mul(") && enabled {
				a, _ := strconv.ParseInt(v[1], 10, 64)
				b, _ := strconv.ParseInt(v[2], 10, 64)
				total += int(a) * int(b)
			} else if strings.Contains(val, "don't(") {
				enabled = false
			} else if strings.Contains(val, "do(") {
				enabled = true
			} else {
				if enabled {
					panic(fmt.Sprintf("not do, don't or mul: %s", val))
				}
			}
		}
	}
	return total
}
