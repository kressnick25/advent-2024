package day04p2

import (
	"io"
	"strings"

	"aoc/utils"
)

const xmas = "XMAS"

func rightSamMas(lines []string, i int, j int) bool {
	ari, arj := i-1, j+1
	if ari < 0 || arj >= len(lines[0]) {
		return false
	}
	ar := strings.Split(lines[ari], "")[arj]
	bli, blj := i+1, j-1
	if bli >= len(lines) || blj < 0 {
		return false
	}
	bl := strings.Split(lines[bli], "")[blj]
	if ar == "M" && bl == "S" {
		return true
	}
	if ar == "S" && bl == "M" {
		return true
	}
	return false
}

func leftSamMas(lines []string, i int, j int) bool {
	ali, alj := i-1, j-1
	if ali < 0 || alj < 0 {
		return false
	}
	al := strings.Split(lines[ali], "")[alj]
	bri, brj := i+1, j+1
	if bri >= len(lines) || brj >= len(lines[0]) {
		return false
	}
	br := strings.Split(lines[bri], "")[brj]
	if al == "M" && br == "S" {
		return true
	}
	if al == "S" && br == "M" {
		return true
	}
	return false
}

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	total := 0
	for i, line := range lines {
		for j, char := range strings.Split(line, "") {
			if char == "A" {
				if leftSamMas(lines, i, j) && rightSamMas(lines, i, j) {
					total += 1
				}
			}
		}
	}

	return total
}
