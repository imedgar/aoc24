package day4

import (
	"bufio"
	"fmt"

	"github.com/imedgar/aoc24-imedgar/utils"
)

var (
	patt  = []rune{'X', 'M', 'A', 'S'}
	patt2 = []rune{'M', 'A', 'S'}
)

func Day4() {
	file := utils.ReadFile("./day4/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var xmap [][]rune
	total := 0
	total2 := 0
	utils.SortRuneSlice(patt2)

	for scanner.Scan() {
		txt := scanner.Text()
		xmap = append(xmap, []rune(txt))
	}

	for i, line := range xmap {
		for ii, col := range line {
			if col == patt[0] {
				nXmas := checkXmas(xmap, i, ii)
				if nXmas > 0 {
					total += nXmas
				}
			}
			if col == patt[2] && checkXm(xmap, i, ii) {
				total2++
			}
		}
	}
	fmt.Println("total", total)
	fmt.Println("total2", total2)
}

func checkXmas(xmap [][]rune, l, c int) int {
	isXmas := []bool{
		checkRight(xmap, l, c),
		checkLeft(xmap, l, c),
		checkTop(xmap, l, c),
		checkBot(xmap, l, c),
		checkBotRight(xmap, l, c),
		checkBotLeft(xmap, l, c),
		checkTopRight(xmap, l, c),
		checkTopLeft(xmap, l, c),
	}
	n := 0
	for _, v := range isXmas {
		if v {
			n++
		}
	}
	return n
}

func checkRight(xmap [][]rune, l, c int) bool {
	line := xmap[l]
	if c+3 > len(line)-1 {
		return false
	}

	for i := 1; i < 4; i++ {
		t := xmap[l][c+i]
		if !checkMatch(t, patt[i]) {
			return false
		}
	}
	return true
}

func checkLeft(xmap [][]rune, l, c int) bool {
	if c-3 < 0 {
		return false
	}

	for i := 1; i < 4; i++ {
		t := xmap[l][c-i]
		if !checkMatch(t, patt[i]) {
			return false
		}
	}
	return true
}

func checkTop(xmap [][]rune, l, c int) bool {
	if l-3 < 0 {
		return false
	}
	for i := 1; i < 4; i++ {
		t := xmap[l-i][c]

		if !checkMatch(t, patt[i]) {
			return false
		}
	}
	return true
}

func checkBot(xmap [][]rune, l, c int) bool {
	if l+3 > len(xmap)-1 {
		return false
	}

	for i := 1; i < 4; i++ {
		t := xmap[l+i][c]
		if !checkMatch(t, patt[i]) {
			return false
		}
	}
	return true
}

func checkBotRight(xmap [][]rune, l, c int) bool {
	if l+3 > len(xmap)-1 || c+3 > len(xmap[l])-1 {
		return false
	}

	for i := 1; i < 4; i++ {
		t := xmap[l+i][c+i]
		if !checkMatch(t, patt[i]) {
			return false
		}
	}
	return true
}

func checkBotLeft(xmap [][]rune, l, c int) bool {
	if l+3 > len(xmap)-1 || c-3 < 0 {
		return false
	}

	for i := 1; i < 4; i++ {
		t := xmap[l+i][c-i]
		if !checkMatch(t, patt[i]) {
			return false
		}
	}
	return true
}

func checkTopRight(xmap [][]rune, l, c int) bool {
	if l-3 < 0 || c+3 > len(xmap[l])-1 {
		return false
	}

	for i := 1; i < 4; i++ {
		t := xmap[l-i][c+i]
		if !checkMatch(t, patt[i]) {
			return false
		}
	}
	return true
}

func checkTopLeft(xmap [][]rune, l, c int) bool {
	if l-3 < 0 || c-3 < 0 {
		return false
	}

	for i := 1; i < 4; i++ {
		t := xmap[l-i][c-i]
		if !checkMatch(t, patt[i]) {
			return false
		}
	}
	return true
}

func checkXm(xmap [][]rune, l, c int) bool {
	if l == 0 || c == len(xmap[l])-1 || l == len(xmap)-1 || c == 0 {
		return false
	}
	fDiag := []rune{xmap[l-1][c-1], xmap[l][c], xmap[l+1][c+1]}
	sDiag := []rune{xmap[l+1][c-1], xmap[l][c], xmap[l-1][c+1]}

	utils.SortRuneSlice(fDiag)
	utils.SortRuneSlice(sDiag)
	for i, v := range patt2 {
		if v != fDiag[i] || v != sDiag[i] {
			return false
		}
	}

	return true
}

func checkMatch(r, p rune) bool {
	if r != p {
		return false
	}
	return true
}
