package utils

import (
	"os"
	"sort"
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

func RunesToIntSlice(input []rune) []int {
	output := []int{}
	for _, r := range input {
		if num, ok := RuneToInt(r); ok {
			output = append(output, num)
		}
	}
	return output
}

func SortRuneSlice(runes []rune) {
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
}

func RemoveElementFrom[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return slice
	}

	newSlice := make([]T, 0, len(slice)-1)
	newSlice = append(newSlice, slice[:index]...)
	newSlice = append(newSlice, slice[index+1:]...)

	return newSlice
}

func InsertAtAny[T any](slice []T, ele T, idx int) []T {
	if idx < 0 || idx > len(slice) {
		panic("Position out of bounds")
	}
	slice = append(slice[:idx], append([]T{ele}, slice[idx:]...)...)
	return slice
}

func DeepCopy[T any](src [][]T) [][]T {
	copy := make([][]T, len(src)) // Create a new slice of slices
	for i := range src {
		copy[i] = make([]T, len(src[i])) // Create a new sub-slice for each row
		for j := range src[i] {
			copy[i][j] = src[i][j] // Copy each element
		}
	}
	return copy
}

func CopySlice[T any](src []T) []T {
	dst := make([]T, len(src))
	copy(dst, src)
	return dst
}

func ReverseSlice[T any](s []T) []T {
	n := len(s)
	for i := 0; i < n/2; i++ {
		s[i], s[n-1-i] = s[n-1-i], s[i]
	}
	return s
}

func FindIndexes[T comparable](slice []T, match T) []int {
	var indexes []int
	for i, el := range slice {
		if el == match {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

func MoveTo[T any](slice []T, from, to int) []T {
	if from < 0 || from >= len(slice) || to < 0 || to >= len(slice) {
		panic("index out of bounds")
	}
	if from == to {
		return slice
	}

	ele := slice[from]
	if from < to {
		copy(slice[from:], slice[from+1:to+1]) // Shift left
	} else {
		copy(slice[to+1:from+1], slice[to:from]) // Shift right
	}
	slice[to] = ele

	return slice
}

func StrToInt(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return n
}

func RuneToInt(r rune) (int, bool) {
	if r >= '0' && r <= '9' {
		return int(r - '0'), true
	}
	return 0, false
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) Add(element T) {
	s[element] = struct{}{}
}

func (s Set[T]) Remove(element T) {
	delete(s, element)
}

func (s Set[T]) Contains(element T) bool {
	_, exists := s[element]
	return exists
}

func (s Set[T]) Size() int {
	return len(s)
}
