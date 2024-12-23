package day05p1

import (
	"io"
	"math"
	"strconv"
	"strings"

	"aoc/utils"
)

type Rule struct {
	page   int
	before int
}

func (r *Rule) isSatisfied(printOrder []int) bool {
	pageIdx := utils.IndexOf(printOrder, r.page)
	if pageIdx == -1 {
		return true
	}
	beforeIdx := utils.IndexOf(printOrder, r.before)
	if beforeIdx == -1 {
		return true
	}
	if beforeIdx < pageIdx {
		return false
	} else {
		return true
	}
}

func getMiddle(printOrder []int) int {
	middleIdx := int(math.Floor(float64(len(printOrder)) / 2))
	return printOrder[middleIdx]
}

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	var rules []Rule
	var printOrders [][]int

	// parse rules
	for _, line := range lines {
		if strings.Split(line, "\n")[0] == "" {
			// empty line encountered
			break
		}
		split := strings.Split(line, "|")
		page, err := strconv.ParseInt(split[0], 10, 64)
		if err != nil {
			panic(err)
		}
		before, err := strconv.ParseInt(split[1], 10, 64)
		if err != nil {
			panic(err)
		}
		r := Rule{
			page:   int(page),
			before: int(before),
		}
		rules = append(rules, r)
	}

	// parse printOrders
	for _, line := range lines {
		if !strings.Contains(line, ",") {
			continue
		}
		numbers := strings.Split(line, ",")
		ints := make([]int, len(numbers))
		for i, n := range numbers {
			parsed, err := strconv.ParseInt(n, 10, 64)
			if err != nil {
				panic(err)
			}
			ints[i] = int(parsed)
		}
		printOrders = append(printOrders, ints)
	}

	var validOrders [][]int
	for _, order := range printOrders {
		ruleBroken := false
		for _, rule := range rules {
			if !rule.isSatisfied(order) {
				ruleBroken = true
				break
			}
		}
		if ruleBroken {
			continue
		}
		validOrders = append(validOrders, order)
	}

	total := 0
	for _, order := range validOrders {
		total += getMiddle(order)
	}

	return total
}
