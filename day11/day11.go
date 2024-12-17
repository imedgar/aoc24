package day11

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/imedgar/aoc24-imedgar/utils"
)

var blinks = 75

func Day11() {
	file := utils.ReadFile("./day11/input.txt")

	scanner := bufio.NewScanner(file)

	stones := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		for _, stone := range strings.Fields(line) {
			stones[stone]++
		}
	}

	for i := 0; i < blinks; i++ {
		newStones := make(map[string]int)

		for stone, count := range stones {
			stoneNum := utils.StrToInt(stone)
			stoneLen := len(stone)

			if stoneNum == 0 {
				newStones["1"] += count
			} else if stoneLen%2 == 0 {
				mid := stoneLen / 2
				left := stone[:mid]
				right := stone[mid:]
				rightNum := utils.StrToInt(right)
				if rightNum == 0 {
					newStones[left] += count
					newStones["0"] += count
				} else {
					newStones[left] += count
					newStones[strconv.Itoa(rightNum)] += count
				}
			} else {
				newStones[strconv.Itoa(stoneNum*2024)] += count
			}
		}

		stones = newStones
	}

	total := 0
	for _, count := range stones {
		total += count
	}

	fmt.Println("total stones:", total)
}
