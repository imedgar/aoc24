package day7

import (
	"bufio"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/imedgar/aoc24-imedgar/utils"
)

type Equation struct {
	result    int
	operators []int
}

var toSolve = []Equation{}
var operations = []string{"+", "*", "||"}
var total = 0
var lastSolved = 0

func Day7() {
	file := utils.ReadFile("./day7/input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`^\s*(\d+):((?: \d+)+)\s*$`)
		matches := re.FindStringSubmatch(line)
		if matches == nil {
			fmt.Println("No match")
			return
		}

		result, _ := strconv.Atoi(matches[1])
		restNumbers := strings.Fields(matches[2])
		numbers := utils.StrSliceToInt(restNumbers)
		equation := Equation{result, numbers}
		toSolve = append(toSolve, equation)
	}

	for _, eq := range toSolve {
		numCombinations := int(math.Pow(float64(len(operations)), float64(len(eq.operators)-1)))
		for i := 0; i < numCombinations; i++ {
			if lastSolved == eq.result {
				continue
			}
			expression := buildExpression(eq, i)
			if eq.result == evaluateExpression(expression) {
				total += eq.result
				lastSolved = eq.result
			}
		}
	}
	fmt.Println("total", total)
}

func buildExpression(eq Equation, combination int) string {
	expression := fmt.Sprintf("%d", eq.operators[0])
	for i := 0; i < len(eq.operators)-1; i++ {
		opIndex := (combination / int(math.Pow(float64(len(operations)), float64(i)))) % len(operations)
		expression += fmt.Sprintf(" %s %d", operations[opIndex], eq.operators[i+1])
	}
	return expression
}

func evaluateExpression(expression string) int {
	tokens := strings.Fields(expression)
	result, _ := strconv.Atoi(tokens[0])
	for i := 1; i < len(tokens); i += 2 {
		op := tokens[i]
		num, _ := strconv.Atoi(tokens[i+1])
		if op == "+" {
			result += num
		} else if op == "*" {
			result *= num
		} else if op == "||" {
			strComb := fmt.Sprintf("%d%s", result, tokens[i+1])
			combination, _ := strconv.Atoi(strComb)
			result = combination
		}
	}
	return result
}
