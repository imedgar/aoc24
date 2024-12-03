package day1

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/imedgar/aoc24-imedgar/utils"
)

func Day1() {
	file := utils.ReadFile("./day1/aoc_01.txt")
	defer file.Close()

	left := []int{}
	right := []int{}
	leftMin, rightMin := -1, -1
	left2 := []int{}
	ocurrences := make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) == 2 {
			leftVal := utils.StrToInt(parts[0])
			rightVal := utils.StrToInt(parts[1])
			// part1
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

			// part2
			left2 = append(left2, []int{leftVal}...)
			i, ok := ocurrences[rightVal]
			if !ok {
				ocurrences[rightVal] = 1
			} else {
				ocurrences[rightVal] = i + 1
			}
		} else {
			fmt.Println("Invalid format:", line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	distance := 0
	distance2 := 0
	for i := 0; i < len(left); i++ {
		distance += abs(left[i] - right[i])

		occ := 0
		v, ok := ocurrences[left[i]]
		if ok {
			occ = left[i] * v
		}
		distance2 += occ
	}
	fmt.Println("Total distance is", distance)
	fmt.Println("Total distance 2 is", distance2)
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

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
