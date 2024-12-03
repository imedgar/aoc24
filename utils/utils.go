package utils

import (
	"os"
	"strconv"
)

func ReadFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}

func StrSliceToInt(input []string) []int {
	result := make([]int, len(input))
	for i, s := range input {
		result[i] = StrToInt(s)
	}
	return result
}

func StrToInt(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return n
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
