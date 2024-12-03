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

func NewSliceWithoutElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}

	newSlice := make([]int, 0, len(slice)-1)
	newSlice = append(newSlice, slice[:index]...)
	newSlice = append(newSlice, slice[index+1:]...)

	return newSlice
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
