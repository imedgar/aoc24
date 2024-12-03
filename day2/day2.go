package day2

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/imedgar/aoc24-imedgar/utils"
)

type State struct {
	safe      bool
	bhv       string
	recAmount int
	hasRecov  bool
}

var (
	state  = State{true, "", 1, false}
	output = 0
)

func Day2() {
	file := utils.ReadFile("./day2/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lvls := utils.StrSliceToInt(strings.Fields(line))

		state = State{true, "", 1, false}

		fmt.Println("handle core", lvls)
		for i := 0; i < len(lvls); i++ {
			safe := handleLvl(i, lvls)
			if !safe {
				fmt.Println("unsafe lvl", lvls[i], "core", lvls)
				state.safe = false
			}
		}
		if state.safe {
			output++
		} else {
			// fmt.Println("unsafe", parts)
		}

	}
	fmt.Println("\nsafe cores", output)
}

func handleLvl(i int, lvls []int) bool {
	if i == len(lvls)-1 { // skip last one
		return true
	}

	curr := lvls[i]
	nxt := lvls[i+1]
	prevBhv := state.bhv
	state.bhv = setBehaviour(curr, nxt)
	if prevBhv != "" && prevBhv != state.bhv {
		fmt.Println("unstable INC/DEC detected", lvls, prevBhv, state.bhv)
		return false
	}

	diff := utils.Abs(curr - nxt)

	fmt.Println("lvl", curr, nxt, "diff", diff, "bhv", state.bhv)
	if state.bhv == "EQ" { // unsafe if not inc or dec
		return false
	}
	if diff > 3 || diff < 1 { // unsafe if diff is >3 or <1
		return false
	}
	if state.bhv == "DEC" && curr < nxt { // unsafe if when Decreasing theres an Increase
		return false
	}
	if state.bhv == "INC" && curr > nxt {
		return false
	}
	return true
}

func setBehaviour(curr, next int) string {
	if curr > next {
		return "DEC"
	} else if curr < next {
		return "INC"
	}
	return "EQ"
}
