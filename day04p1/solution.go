package day04p1

import (
	"io"
	"strings"

	"aoc/utils"
)

const xmas = "XMAS"

func checkRight(line string, start int) bool {
	s := 0
	for i := start; i < (start + 4); i++ {
		if i > len(line)-1 {
			return false
		}

		if line[i] != xmas[s] {
			return false
		}
		s++
	}
	return true
}

func checkLeft(line string, start int) bool {
	s := 0
	for i := start; i > (start - 4); i-- {
		if i < 0 {
			return false
		}

		if line[i] != xmas[s] {
			return false
		}
		s++
	}
	return true
}

func checkDown(lines []string, startRow int, startCol int) bool {
	s := 0
	for i := startRow; i < (startRow + 4); i++ {
		if i > len(lines)-1 {
			return false
		}

		if lines[i][startCol] != xmas[s] {
			return false
		}
		s++
	}
	return true
}

func checkUp(lines []string, startRow int, startCol int) bool {
	s := 0
	for i := startRow; i > (startRow - 4); i-- {
		if i < 0 {
			return false
		}

		if lines[i][startCol] != xmas[s] {
			return false
		}
		s++
	}
	return true
}

func checkDiagUpLeft(lines []string, startRow int, startCol int) bool {
	s := 0
	for i, j := startRow, startCol; i > (startRow - 4); i, j = i-1, j-1 {
		if i < 0 || j < 0 {
			return false
		}

		if lines[i][j] != xmas[s] {
			return false
		}
		s++
	}
	return true
}

func checkDiagUpRight(lines []string, startRow int, startCol int) bool {
	s := 0
	for i, j := startRow, startCol; i > (startRow - 4); i, j = i-1, j+1 {
		if i < 0 || j > (len(lines[0])-1) {
			return false
		}

		if lines[i][j] != xmas[s] {
			return false
		}
		s++
	}
	return true
}

func checkDiagDownLeft(lines []string, startRow int, startCol int) bool {
	s := 0
	for i, j := startRow, startCol; i < (startRow + 4); i, j = i+1, j-1 {
		if (i > len(lines)-1) || j < 0 {
			return false
		}

		if lines[i][j] != xmas[s] {
			return false
		}
		s++
	}
	return true
}

func checkDiagDownRight(lines []string, startRow int, startCol int) bool {
	s := 0
	for i, j := startRow, startCol; i < (startRow + 4); i, j = i+1, j+1 {
		if (i > len(lines)-1) || j > (len(lines[0])-1) {
			return false
		}

		if lines[i][j] != xmas[s] {
			return false
		}
		s++
	}
	return true
}

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	total := 0
	for i, line := range lines {
		for j, char := range strings.Split(line, "") {
			if char == "X" {
				// fmt.Printf("co-ords: [%v, %v] ", i, j)
				// fmt.Printf("checkRight: %v ", checkRight(line, j))
				// fmt.Printf("checkUp: %v ", checkUp(lines, i, j))
				// fmt.Printf("checkDown: %v ", checkDown(lines, i, j))
				// fmt.Printf("checkLeft: %v\n", checkLeft(line, j))
				if checkRight(line, j) {
					total += 1
				}
				if checkLeft(line, j) {
					total += 1
				}
				if checkUp(lines, i, j) {
					total += 1
				}
				if checkDown(lines, i, j) {
					total += 1
				}
				if checkDiagUpLeft(lines, i, j) {
					total += 1
				}
				if checkDiagUpRight(lines, i, j) {
					total += 1
				}
				if checkDiagDownLeft(lines, i, j) {
					total += 1
				}
				if checkDiagDownRight(lines, i, j) {
					total += 1
				}
			}
		}
	}

	return total
}
