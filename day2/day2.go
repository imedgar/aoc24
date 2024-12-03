package day2

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/imedgar/aoc24-imedgar/utils"
)

func Day2() {
	file := utils.ReadFile("./day2/input.txt")
	defer file.Close()

	safeCount := 0
	dampenedSafeCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		report := utils.StrSliceToInt(strings.Fields(scanner.Text()))

		if isSafeReport(report) {
			safeCount++
			dampenedSafeCount++
			continue
		}

		for i := 0; i < len(report); i++ {
			dampened := append([]int(nil), report[:i]...)
			dampened = append(dampened, report[i+1:]...)
			if isSafeReport(dampened) {
				dampenedSafeCount++
				break
			}
		}
	}
	fmt.Println("Safe:", safeCount)
	fmt.Println("Dampened:", dampenedSafeCount)
}

func isSafeReport(report []int) bool {
	increaseCount, decreaseCount := 0, 0

	for i := 1; i < len(report); i++ {
		curr := report[i]
		prev := report[i-1]

		if utils.Abs(curr-prev) > 3 {
			return false
		}

		if curr > prev {
			increaseCount++
		} else if curr < prev {
			decreaseCount++
		}
	}

	return increaseCount == len(report)-1 || decreaseCount == len(report)-1
}
