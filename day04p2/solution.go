package day04p2

import (
	"aoc/utils"
	"io"
)

const xmas = "XMAS"

func rightSamMas(lines [][]rune, i int, j int, lenRow int, lenCol int) bool {
	ari, arj := i-1, j+1
	if ari < 0 || arj >= lenCol {
		return false
	}
	ar := lines[ari][arj]
	bli, blj := i+1, j-1
	if bli >= lenRow || blj < 0 {
		return false
	}
	bl := lines[bli][blj]
	if ar == 'M' && bl == 'S' {
		return true
	}
	if ar == 'S' && bl == 'M' {
		return true
	}
	return false
}

func leftSamMas(lines [][]rune, i int, j int, lenRow int, lenCol int) bool {
	ali, alj := i-1, j-1
	if ali < 0 || alj < 0 {
		return false
	}
	al := lines[ali][alj]
	bri, brj := i+1, j+1
	if bri >= lenRow || brj >= lenCol {
		return false
	}
	br := lines[bri][brj]
	if al == 'M' && br == 'S' {
		return true
	}
	if al == 'S' && br == 'M' {
		return true
	}
	return false
}

func Solve(r io.Reader) any {
	lines, err := utils.ReadChars(r)
	if err != nil {
		panic(err)
	}

	total := 0
	lenRow := len(lines)
	lenCol := len(lines[0])
	for i, line := range lines {
		for j, char := range []rune(line) {
			if char == 'A' {
				if leftSamMas(lines, i, j, lenRow, lenCol) && rightSamMas(lines, i, j, lenRow, lenCol) {
					total += 1
				}
			}
		}
	}

	return total
}
