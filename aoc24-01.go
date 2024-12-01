package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("aoc_01.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	left := []int{}
	right := []int{}
	leftMin, rightMin := -1, -1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")

		if len(parts) == 2 {
			leftVal, rightVal := strToInt(parts[0], parts[1])

			if isSmaller(leftMin, leftVal) {
				left = insertSmaller(left, leftVal)
				leftMin = leftVal
			} else {
				left = insertBigger(left, leftVal)
			}

			if isSmaller(rightMin, rightVal) {
				right = insertSmaller(right, rightVal)
				rightMin = rightVal
			} else {
				right = insertBigger(right, rightVal)
			}

		} else {
			fmt.Println("Invalid format:", line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	distance := 0
	for i := 0; i < len(left); i++ {
		curr := 0
		if left[i] < right[i] {
			curr = right[i] - left[i]
		} else if right[i] < left[i] {
			curr = left[i] - right[i]
		}
		distance += curr
	}
	fmt.Println("Total distance is", distance)
}

func isSmaller(min, val int) bool {
	if min == -1 || min >= val {
		return true
	} else {
		return false
	}
}

func insertSmaller(slice []int, value int) []int {
	slice = append([]int{value}, slice...)
	return slice
}

func insertBigger(slice []int, value int) []int {
	if slice[len(slice)-1] <= value {
		return append(slice, []int{value}...)
	}

	for i, v := range slice {
		if v >= value {
			slice = append(slice[:i+1], slice[i:]...)
			slice[i] = value
			return slice
		}
	}
	return slice
}

func strToInt(left, right string) (int, int) {
	leftVal, err := strconv.Atoi(left)
	if err != nil {
		panic(err)
	}
	rightVal, err := strconv.Atoi(right)
	if err != nil {
		panic(err)
	}
	return leftVal, rightVal
}
