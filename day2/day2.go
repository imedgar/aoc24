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

func Day2() {
	file := utils.ReadFile("./day2/test.txt")
	defer file.Close()

	output := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		f := utils.StrToInt(parts[0])
		s := utils.StrToInt(parts[1])
		t := utils.StrToInt(parts[2])

		state := State{true, "", 1, false}

		if f == s && s == t { // neither inc or dec
			continue
		}

		if f > t {
			state.bhv = "dec"
		} else if f < t {
			state.bhv = "inc"
		}

		for i := 0; i < len(parts); i++ {
			curr := utils.StrToInt(parts[i])
			// Handle first element
			if i == 0 {
				nxt := utils.StrToInt(parts[i+1])
				diff := utils.Abs(curr - nxt)
				// first element is corrupt
				if !isSafe(diff, curr, nxt, state.bhv) {
					nxt2 := utils.StrToInt(parts[i+2])
					diff := utils.Abs(nxt - nxt2)
					// can recover ?
					if !isSafe(diff, nxt, nxt2, state.bhv) {
						state.safe = false
						break
					}
					state.recAmount--
					i++
				}
			} else if i > 0 && i < len(parts)-1 { // handle in between
				prev := utils.StrToInt(parts[i-1])
				nxt := utils.StrToInt(parts[i+1])
				if prev == curr && curr == nxt {
					state.safe = false
					break
				}
				if prev == nxt {
					state.safe = false
					break
				}
				diffPrev := utils.Abs(prev - curr)
				if !isSafe(diffPrev, prev, curr, state.bhv) {
					if state.recAmount == 0 {
						state.safe = false
						break
					}
					diffNxt := utils.Abs(prev - nxt)
					if !isSafe(diffNxt, prev, nxt, state.bhv) {
						state.safe = false
						break
					}
					state.recAmount--
					i++
				}
			} else if i == len(parts)-1 { // Handle last element
				prev := utils.StrToInt(parts[i-1])
				diff := utils.Abs(prev - curr)
				if !isSafe(diff, prev, curr, state.bhv) {
					if state.recAmount == 0 {
						state.safe = false
						break
					}
				}
			}
		}
		if state.safe {
			output++
		} else {
			// fmt.Println("unsafe", parts)
		}

	}
	fmt.Println("safe cores", output)
}

func isSafe(diff, prev, curr int, bhv string) bool {
	return !((diff > 3 || diff < 1) ||
		(bhv == "dec" && prev < curr) ||
		(bhv == "inc" && prev > curr) ||
		(prev == curr))
}
