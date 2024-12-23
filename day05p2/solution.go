package day05p2

import (
	"io"
	"math"
	"slices"
	"strconv"
	"strings"

	"aoc/utils"
)

type Rule struct {
	page   int
	before int
}

func (r *Rule) isSatisfied(printOrder []int) bool {
	if !(slices.Contains(printOrder, r.page) && slices.Contains(printOrder, r.before)) {
		return true
	}
	pageIdx := utils.IndexOf(printOrder, r.page)
	beforeIdx := utils.IndexOf(printOrder, r.before)
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

	var invalidOrders [][]int
	for _, order := range printOrders {
		ruleBroken := false
		for _, rule := range rules {
			if !rule.isSatisfied(order) {
				ruleBroken = true
				break
			}
		}
		if !ruleBroken {
			continue
		}
		invalidOrders = append(invalidOrders, order)
	}

	rulesMap := make(map[int][]int)
	for _, r := range rules {
		rulesMap[r.before] = append(rulesMap[r.before], r.page)
	}

	total := 0
	for _, order := range invalidOrders {
		o := make([]int, len(order))
		copy(o, order)
		// bubble sort
		for i := 0; i < len(order)-1; i++ {
			swapped := true
			for swapped {
				rule := rulesMap[o[i]]
				for j := i + 1; j < len(o); j++ {
					if slices.Contains(rule, o[j]) {
						o[i], o[j] = o[j], o[i]
						swapped = true
						break
					}
					swapped = false
				}
			}
		}
		total += getMiddle(o)
	}

	return total
}
