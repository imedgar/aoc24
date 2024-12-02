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
	right := make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")

		if len(parts) == 2 {
			leftVal, rightVal := strToInt(parts[0], parts[1])
			left = append(left, []int{leftVal}...)
			i, ok := right[rightVal]
			if !ok {
				right[rightVal] = 1
			} else {
				right[rightVal] = i + 1
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
		v, ok := right[left[i]]
		if ok {
			curr = left[i] * v
		}
		distance += curr
	}
	fmt.Println("Total distance is", distance)
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
