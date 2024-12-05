package day5

import (
	"bufio"
	"fmt"
	"slices"
	"strings"

	"github.com/imedgar/aoc24-imedgar/utils"
)

type Rule struct {
	number int
	before []int
	after  []int
}

func Day5() {
	file := utils.ReadFile("./day5/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rules := make(map[int]Rule)
	manuals := [][]int{}
	pages := 0

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {

			numbers := strings.Split(line, "|")
			if len(numbers) == 2 {
				first := utils.StrToInt(numbers[0])
				second := utils.StrToInt(numbers[1])
				firstRule, firstExist := rules[first]
				secRule, secExist := rules[second]
				if !firstExist {
					rules[first] = Rule{first, []int{second}, []int{}}
				} else {
					firstRule.before = append(firstRule.before, second)
					rules[first] = firstRule
				}
				if !secExist {
					rules[second] = Rule{second, []int{}, []int{first}}
				} else {
					secRule.after = append(secRule.after, first)
					rules[second] = secRule
				}
			}
		} else if strings.Contains(line, ",") {
			check := utils.StrSliceToInt(strings.Split(line, ","))
			checked := []int{}
			toCheck := make([]int, len(check))
			copy(toCheck, check)
			sorted := true
			for _, page := range check {
				toCheck = toCheck[1:]

				for _, b := range checked {
					if slices.Contains(rules[page].before, b) {
						sorted = false
						break
					}
				}
				for _, a := range toCheck {
					if slices.Contains(rules[page].after, a) {

						sorted = false
						break
					}
				}
				checked = append(checked, page)
			}
			if sorted {
				manuals = append(manuals, check)
				pages += check[len(check)/2]
			}
		}
	}
	fmt.Println("total pages", pages)
}

func toIdx(idx, len int) int {
	if idx == len {
		return len - 1
	}
	return idx
}
