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

var rules = make(map[int]Rule)

func Day5() {
	file := utils.ReadFile("./day5/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	manuals := [][]int{}
	toFix := [][]int{}
	pages := 0
	fixedPages := 0

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
			toCheck := append([]int(nil), check...)
			sorted := true
			for _, page := range check {
				toCheck = toCheck[1:]
				before, _ := checkBefore(checked, page)
				if !before {
					sorted = false
				}
				after, _ := checkAfter(toCheck, page)
				if !after {
					sorted = false
				}
				checked = append(checked, page)
			}

			if sorted {
				manuals = append(manuals, check)
				pages += check[len(check)/2]
			} else {
				toFix = append(toFix, check)
			}
		}
	}
	fmt.Println("total pages", pages)
	for _, fix := range toFix { // manual to fix
		fixed := false
		for !fixed {
			toCheck := append([]int(nil), fix...)
			checked := []int{}
			isOk := true
			for _, c := range fix {
				toCheck = toCheck[1:]
				before, idxBef := checkBefore(toCheck, c)
				if !before {
					isOk = false
					fix = utils.MoveTo(fix, slices.Index(fix, toCheck[idxBef]), slices.Index(fix, c))
				}
				checked = append(checked, c)
			}
			if isOk {
				fixed = true
				fix = utils.ReverseSlice(fix)
				fixedPages += fix[len(fix)/2]
			}
		}
	}
	fmt.Println("total fixed pages", fixedPages)
}

func checkBefore(checked []int, page int) (bool, int) {
	for i, b := range checked {
		if slices.Contains(rules[page].before, b) {
			return false, i
		}
	}
	return true, -1
}

func checkAfter(toCheck []int, page int) (bool, int) {
	for i, a := range toCheck {
		if slices.Contains(rules[page].after, a) {
			return false, i
		}
	}
	return true, -1
}

func toIdx(idx, len int) int {
	if idx == len {
		return len - 1
	}
	return idx
}
