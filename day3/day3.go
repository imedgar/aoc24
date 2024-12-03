package day3

import (
	"bufio"
	"fmt"
	"regexp"

	"github.com/imedgar/aoc24-imedgar/utils"
)

func Day03() {
	file := utils.ReadFile("./day3/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	enabled := true

	for scanner.Scan() {
		txt := scanner.Text()

		pattern := `mul\((\d{1,3}),(\d{1,3})\)|don\'t\(\)|do\(\)`
		re := regexp.MustCompile(pattern)
		matches := re.FindAllStringSubmatch(txt, -1)

		for _, match := range matches {
			if len(match) > 0 {
				fullMatch := match[0]
				if fullMatch == "do()" {
					enabled = true
				}
				if fullMatch == "don't()" {
					enabled = false
				}
				if match[1] != "" && match[2] != "" {
					first := utils.StrToInt(match[1])
					second := utils.StrToInt(match[2])
					if enabled {
						total += first * second
					}
				}
			}
		}
	}
	fmt.Println("Total", total)
}
