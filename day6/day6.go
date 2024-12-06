package day6

import (
	"bufio"
	"fmt"
	"strings"
	"time"

	"github.com/imedgar/aoc24-imedgar/utils"
)

var (
	lab      = [][]rune{}
	guardMov = []string{"^", ">", "v", "<"}
	guardPos = &[]int{}
)

func Day6() {
	file := utils.ReadFile("./day6/input.txt")

	fl := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		located := locateGuard(line)
		if located != "" {
			*guardPos = []int{fl, strings.Index(line, located)}
		}
		lab = append(lab, []rune(line))
		fl++
	}
	for !move(guardPos) {
		//		debug(100)
	}
	steps := 1
	for _, f := range lab {
		for _, t := range f {
			if t == 'X' {
				steps++
			}
		}
	}
	fmt.Println("finished patrolling", steps)
}

func locateGuard(floor string) string {
	for _, t := range guardMov {
		if strings.Contains(floor, t) {
			return t
		}
	}
	return ""
}

func debug(milli time.Duration) {
	clearScreen()
	printLab()
	time.Sleep(milli * time.Millisecond)
}

func printLab() {
	fmt.Println("\nLab:")
	for _, f := range lab {
		fmt.Println(string(f))
	}
}

func move(currPos *[]int) bool {
	pos := *currPos
	x, y := pos[0], pos[1]
	guard := lab[x][y]

	switch guard {
	case '^': // Move up
		if x-1 < 0 {
			return true
		}
		if lab[x-1][y] == '#' {
			lab[x][y] = 'X'
			lab[x][y+1] = '>'
			(*currPos)[1] = y + 1
		} else if lab[x-1][y] == '.' || lab[x-1][y] == 'X' {
			lab[x][y] = 'X'
			lab[x-1][y] = '^'
			(*currPos)[0] = x - 1
		}
	case '>': // Move right
		if y+1 >= len(lab[x]) {
			return true
		}
		if lab[x][y+1] == '#' {
			lab[x][y] = 'X'
			lab[x+1][y] = 'v'
			(*currPos)[0] = x + 1
		} else if lab[x][y+1] == '.' || lab[x][y+1] == 'X' {
			lab[x][y] = 'X'
			lab[x][y+1] = '>'
			(*currPos)[1] = y + 1
		}
	case 'v': // Move down
		if x+1 >= len(lab) {
			return true
		}
		if lab[x+1][y] == '#' {
			lab[x][y] = 'X'
			lab[x][y-1] = '<'
			(*currPos)[1] = y - 1
		} else if lab[x+1][y] == '.' || lab[x+1][y] == 'X' {
			lab[x][y] = 'X'
			lab[x+1][y] = 'v'
			(*currPos)[0] = x + 1
		}
	case '<': // Move left
		if y-1 < 0 {
			return true
		}
		if lab[x][y-1] == '#' {
			lab[x][y] = 'X'
			lab[x-1][y] = '^'
			(*currPos)[0] = x - 1
		} else if lab[x][y-1] == '.' || lab[x][y-1] == 'X' {
			lab[x][y] = 'X'
			lab[x][y-1] = '<'
			(*currPos)[1] = y - 1
		}
	}
	return false
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
