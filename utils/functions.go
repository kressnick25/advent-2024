// Package utils implements various utility functions and procedures which have
// been useful for previous Advents of Code.
package utils

import (
	"bufio"
	"fmt"
	"io"
)

// Greatest Common Denominator
func Gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// Least Common Multiple
func Lcm(a, b int64) int64 {
	return a * b / Gcd(a, b)
}

// Kernighan's Bit Counting Algorithm
func CountBits(n uint64) int64 {
	var count int64 = 0
	for n > 0 {
		n = n & (n - 1)
		count++
	}

	return count
}

// Check if error is not nil and panic with message if it is.
func Check(e error, format string, a ...any) {
	if e != nil {
		message := fmt.Sprintf(format, a...)
		panic(fmt.Errorf("%s: %s", message, e))
	}
}

// Read all lines from reader. Panic if there is an issue
func ReadLines(r io.Reader) []string {
	result := []string{}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	err := scanner.Err()
	Check(err, "error reading lines")

	return result
}

func Abs(a int) int {
	if a < 0 {
		return 0 - a
	}
	return a
}

func Remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func IndexOf(s []int, v int) int {
	for i, val := range s {
		if val == v {
			return i
		}
	}
	return -1
}

func ReadChars(reader io.Reader) ([][]rune, error) {
	scanner := bufio.NewScanner(reader)
	var result [][]rune

	for scanner.Scan() {
		line := scanner.Text()
		runeSlice := []rune(line)
		result = append(result, runeSlice)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
